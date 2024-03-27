package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/alecthomas/log4go"
	"time"
)

var (
	Cfg *Config
)

type Config struct {
	DebugMode      bool `toml:"debug_mode"`
	Fetch8Reserves bool `toml:"fetch_8_reserves"`
	Global         Global
	SwipingCfg     Swiping `toml:"swiping_cfg"`
}

type Swiping struct {
	Start        int    `toml:"start"`
	End          int    `toml:"end"`
	RoutineCount int    `toml:"routine_count"`
	LastIndex    int    `toml:"last_index"`
	ProxyUrl     string `toml:"proxy_url"`
	SleepMills   int64  `toml:"sleep_mills"`
}

//LoadConfig
/*
Load Configuration
*/
func LoadConfig(dir string) {
	log4go.Info("work dir=%s\n", dir)
	Cfg = new(Config)
	if _, err := toml.DecodeFile(dir+"/conf.toml", Cfg); err != nil {
		log4go.Exit(err)
	}
}

/*
MySQL Configuration
*/
type MySQL struct {
	User string `toml:"user"` //user name
	Pasw string `toml:"pasw"` //password
	Prot string `toml:"prot"` //protocol
	Host string `toml:"host"` //host name or ip
	Port string `toml:"port"` //port
	Dbnm string `toml:"dbnm"` //database name
}

type TokenPair struct {
	Token TokenItem `json:"token"` //目标资产(要买或卖的币)
	Base  TokenItem `json:"base"`  //基础资产(筹码)
}

type TokenItem struct {
	Name            string `json:"name"`             //名称
	Symbol          string `json:"symbol"`           //简称
	Decimal         uint8  `json:"decimal"`          //精度
	ContractAddress string `json:"contract_address"` //合约地址
}

/*
Global Configuration
*/
type Global struct {
	Port               int32                        `toml:"port"`
	ServerOrigin       string                       `toml:"server_origin"`
	BscMainNetUrl      string                       `toml:"bsc_main_net_url"`
	BscMainNetUrlRead  string                       `toml:"bsc_main_net_url_read"`
	PrivateKey         string                       `toml:"private_key" json:"-"`
	BscChainId         string                       `toml:"bsc_chain_id"`
	CallerAddr         string                       `toml:"caller_addr"`
	RecycleAddress     string                       `toml:"recycle_address"`
	EmitorBaseUrl      string                       `toml:"emitor_base_url"`
	ContractRouter     string                       `toml:"contract_router"`
	Wallet98           string                       `toml:"wallet_98"`
	PriceRefreshMills  time.Duration                `toml:"price_refresh_mills"`
	DefiPriceSeconds   time.Duration                `toml:"defi_price_seconds"`
	MustCheck          bool                         `toml:"must_check"`
	AssetUrl           string                       `toml:"asset_url"`
	CenterUrl          string                       `toml:"center_url"`
	TgSendMsgTextTpl   string                       `toml:"tg_send_msg_text_tpl"`
	TgDefaultToken     string                       `toml:"tg_default_token"`
	TgDefaultChatId    string                       `toml:"tg_default_chat_id"`
	CalcPriceAmountIn  string                       `toml:"calc_price_amount_in"`
	DispatchAssets     []string                     `toml:"dispatch_assets"`
	BscChainRpcUrls    []string                     `toml:"bsc_chain_rpc_urls"`
	GasPriceAddon      int64                        `toml:"gas_price_addon"`
	GasPriceQuo        int64                        `toml:"gas_price_quo"`
	SwapMaxWait        int64                        `toml:"swap_max_wait"`
	EstimateMulNum     int64                        `toml:"estimate_mul_num"`
	SendTg             bool                         `toml:"send_tg"`
	PubMap             map[string]string            `toml:"pub_map"`
	TokenMap           map[string]string            `toml:"token_map"`
	ZapBscContract     string                       `toml:"zap_bsc_contract"`
	RouterMap          map[string]map[string]string `toml:"router_map"`
	IsMaster           bool                         `toml:"is_master"`
	PairPage           int32                        `toml:"pair_page"`
	PairPageSize       int32                        `toml:"pair_page_size"`
	AmountsMap         map[string]string            `toml:"amounts_map"`
	SymbolWalletMaxMap map[string]int               `toml:"symbol_wallet_max_map"`
	TargetToken        string                       `toml:"target_token"`
	BaseToken          string                       `toml:"base_token"`
	GasCoin            string                       `toml:"gas_coin"`
	CallData           string                       `toml:"call_data"`
	SendGasLimit       uint64                       `toml:"send_gas_limit"`
	BinanceApiKey      string                       `toml:"binance_api_key"`
	BinanceSecretKey   string                       `toml:"binance_secret_key"`
	Symbol             string                       `toml:"symbol"`
	USDTQuantity       float64                      `toml:"usdt_quantity"`
}

type RedisConfig struct {
	Addr               string
	DB                 int
	DialTimeout        int32
	ReadTimeout        int32
	WriteTimeout       int32
	PoolSize           int
	PoolTimeout        int32
	IdleTimeout        int32
	IdleCheckFrequency int32
	Password           string
}
