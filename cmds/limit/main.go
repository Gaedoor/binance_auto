package main

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/xiaoyueya/binance2024/domain"
	"math"
	"strings"
)

func getPow(logsize string) int {
	if strings.Index(logsize, "1") == 0 {
		return 0
	}
	return 1 - strings.Index(logsize, "1")
}

func main() {
	lotSize := "0.01"

	domain.InitBinance()
	symbol := "BOMEUSDT"

	resp, err := domain.BinanceCli.NewTickerPriceService().Symbol(symbol).Do(context.Background())
	if err != nil {
		panic(err)
	}
	price := resp.Price
	fmt.Printf("price=%s\n", price)

	er, err := domain.BinanceCli.NewExchangeInfoService().Do(context.Background())
	if err != nil {
		panic(err)
	}

	for _, item := range er.Symbols {
		if item.Symbol == symbol {
			for n, filter := range item.Filters {
				if filter.FilterType == "LOT_SIZE" {
					lotSize = filter.StepSize
					break
				}
				fmt.Printf("%d,type=%s,filter.StepSize=%s\n", n, filter.FilterType, filter.StepSize)
			}
			fmt.Printf("symbol=%s,ba=%s,qa=%s,qp=%d,bp=%d\n", item.Symbol, item.BaseAsset, item.QuoteAsset, item.QuotePrecision, item.BaseAssetPrecision)
			break
		}
	}
	fmt.Printf("lotsize=%s\n", lotSize)
	priceDecimal, err := decimal.NewFromString(resp.Price)
	if err != nil {
		panic(err)
	}
	ceilPrice, _ := priceDecimal.Ceil().Float64()

	quantity := math.Ceil(1/ceilPrice*math.Pow10(getPow(lotSize))) / math.Pow10(getPow(lotSize))
	fmt.Printf("quantity=%f\n", quantity)
	//return

	order, err := domain.BinanceCli.NewCreateOrderService().Symbol(symbol).Type("LIMIT").Side("BUY").Price(ceilPrice).Quantity(quantity).TimeInForce("GTC").Do(context.Background())
	if err != nil {
		fmt.Printf("order err=%v\n", err)
		return
	}

	fmt.Printf("%v\n", order)

}
