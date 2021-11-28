package configs

// 微信支付
type wxpay struct {

	// 商家账号
	MchId string `yaml:"MchId"`
	// 商户证书序列号
	CertificateSerialNo string `yaml:"CertificateSerialNo"`
	// 商户APIv3密钥
	MchAPIv3Key string `yaml:"MchAPIv3Key"`

	// 内网支付ip限制，支付回调除外
	AuthIpPay []string `yaml:"AuthIpPay"`

	// 认养
	FosterPay fosterPay `yaml:"FosterPay"`

	// app
	BeeAppPay beeAppPay `yaml:"BeeAppPay"`
}

// 认养小程序支付
type fosterPay struct {
	//
	AppId string `yaml:"AppId"`

	// 交易类型
	Currency string `yaml:"Currency"`

	// 完整回调地址
	NotifyUrl string `yaml:"NotifyUrl"`
}

// app支付
type beeAppPay struct {
	AppId string `yaml:"AppId"`
}
