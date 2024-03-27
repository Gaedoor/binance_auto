package main

import (
	"context"
	"fmt"
	"github.com/alecthomas/log4go"
	binanceconnector "github.com/binance/binance-connector-go"
	"github.com/shopspring/decimal"
	"github.com/xiaoyueya/binance2024/conf"
	"github.com/xiaoyueya/binance2024/domain"
	"github.com/xiaoyueya/binance2024/util"
	"math"
	"strings"
	"time"
)

func getPow(lotsize string) int {
	if strings.Index(lotsize, "1") == 0 {
		return 0
	}
	return 1 - strings.Index(lotsize, "1")
}

func main() {

	dir := util.Getwd()
	conf.LoadConfig(dir)
	log4go.Info("binance2024 service version time=%s,program init ing......", time.Now().Format(time.RFC3339))
	domain.InitBinance()

	lotSize := "0.01"
	symbol := conf.Cfg.Global.Symbol
	usd := conf.Cfg.Global.USDTQuantity

	er, err := domain.BinanceCli.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		panic(err)
	}

	for _, item := range er.Symbols {
		if item.Symbol == symbol {
			for _, filter := range item.Filters {
				if filter.FilterType == "LOT_SIZE" {
					lotSize = filter.StepSize
				}
				if filter.FilterType == "NOTIONAL" {
					log4go.Info("filter.MinNotional=%s", filter.MinNotional)
				}
			}
			log4go.Info("symbol=%s,ba=%s,qa=%s,qp=%d,bp=%d", item.Symbol, item.BaseAsset, item.QuoteAsset, item.QuotePrecision, item.BaseAssetPrecision)
			break
		}
	}
	log4go.Info("lotsize=%s", lotSize)

	for {

		resp, err := domain.BinanceCli.NewTickerPriceService().Symbol(symbol).Do(context.Background())
		if err != nil {
			log4go.Info("get price from binance error=%v", err)
			continue
		}
		price := resp.Price
		log4go.Info("price=%s", price)
		priceDecimal, err := decimal.NewFromString(resp.Price)
		if err != nil {
			log4go.Info("decimal from string error=%v", err)
			continue
		}
		ceilPrice, _ := priceDecimal.Float64()

		log4go.Info("price=%f", ceilPrice)

		quantity := math.Ceil(usd/ceilPrice*math.Pow10(getPow(lotSize))) / math.Pow10(getPow(lotSize))
		log4go.Info("quantity=%f", quantity)

		order, err := domain.BinanceCli.NewCreateOrderService().Symbol(symbol).Type("MARKET").Side("BUY").Quantity(quantity).Do(context.Background())
		if err != nil {
			log4go.Info("order err=%v", err)
			continue
		}

		fmt.Printf("resp type=%T\n", order)
		full, ok := order.(*binanceconnector.CreateOrderResponseFULL)
		if ok {
			log4go.Info("%+v", full)
			if full.OrderId > 0 {
				log4go.Info("抢币成功：order id=%d,client order id=%s,order list id=%d", full.OrderId, full.ClientOrderId, full.OrderListId)
				break
			}
		}
		time.Sleep(time.Millisecond * 100)
	}
}
