package goaccount

type ConfigEthereum struct {
	PrivateKey string `json:"PrivateKey"`
}

type Config struct {
	Ethereum *ConfigEthereum `json:"Ethereum"`
}
