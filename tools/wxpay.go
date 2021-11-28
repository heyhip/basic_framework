package tools

import (
	"basic_framework/configs"
	"basic_framework/core/log"
	"context"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/auth/verifiers"
	"github.com/wechatpay-apiv3/wechatpay-go/core/downloader"
	"github.com/wechatpay-apiv3/wechatpay-go/core/notify"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

var WxClient *core.Client
var WxHandler *notify.Handler

// 微信支付
func init() {

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(configs.Yaml.Wxpay.PrivateKey)
	if err != nil {
		log.Panic("load merchant private key error")
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(configs.Yaml.Wxpay.MchId, configs.Yaml.Wxpay.CertificateSerialNo, mchPrivateKey, configs.Yaml.Wxpay.MchAPIv3Key),
	}

	WxClient, err = core.NewClient(ctx, opts...)
	if err != nil {
		log.Panic("new wechat pay client err:%s", err)
	}

	// 获取平台证书访问器
	certVisitor := downloader.MgrInstance().GetCertificateVisitor(configs.Yaml.Wxpay.MchId)
	WxHandler = notify.NewNotifyHandler(configs.Yaml.Wxpay.MchAPIv3Key, verifiers.NewSHA256WithRSAVerifier(certVisitor))
}
