package goaccount

type ConfigEthereum struct {
	PrivateKey        string `json:"PrivateKey"`
	AutoCreateAccount bool   `json:"AutoCreateAccount"`
}
type ConfigMomoMiniApp struct {
	AppID             string `json:"AppID"`
	AutoCreateAccount bool   `json:"AutoCreateAccount"`
	OpenSecret        string `json:"OpenSecret"`
	OpenPrivateKey    string `json:"OpenPrivateKey"`
	OpenPublicKey     string `json:"OpenPublicKey"`
	IsDev             bool   `json:"IsDev"`
}

type Config struct {
	Ethereum    *ConfigEthereum    `json:"Ethereum"`
	MomoMiniApp *ConfigMomoMiniApp `json:"MomoMiniApp"`
}
