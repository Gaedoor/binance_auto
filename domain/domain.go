package domain

import (
	"github.com/alecthomas/log4go"
	binance_connector "github.com/binance/binance-connector-go"
	"github.com/xiaoyueya/binance2024/conf"
)

var BinanceCli *binance_connector.Client

func InitBinance() {
	apiKey := conf.Cfg.Global.BinanceApiKey
	secretKey := conf.Cfg.Global.BinanceSecretKey
	baseURL := "https://api.binance.com"

	BinanceCli = binance_connector.NewClient(apiKey, secretKey, baseURL)
	log4go.Info("init binance client success")
}
