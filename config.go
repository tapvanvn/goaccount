package goaccount

type ConfigEthereum struct {
	PrivateKey        string `json:"PrivateKey"`
	AutoCreateAccount bool   `json:"AutoCreateAccount"`
}
type ConfigMomoMiniApp struct {
	AppID             string `json:"AppID"`
	AutoCreateAccount bool   `json:"AutoCreateAccount"`
}

type Config struct {
	Ethereum    *ConfigEthereum    `json:"Ethereum"`
	MomoMiniApp *ConfigMomoMiniApp `json:"MomoMiniApp"`
}
