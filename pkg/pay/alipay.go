package pay

import (
	"context"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"go-shop-api/global"
)

var ctx = context.Background()

type AliPay struct {
}

func (ali *AliPay) Pay(payType string, data interface{}) (payResult interface{}, err error) {

	return
}

// TradeWapPay 手机Wap支付
func TradeWapPay() (payUrl string, err error) {
	// 建立连接
	client, err := alipay.NewClient(global.GA_CONFIG.AppId, global.GA_CONFIG.PrivateKey, false)
	if err != nil {
		return "", err
	}
	//配置公共参数
	client.SetCharset("utf-8").SetSignType(alipay.RSA2).SetReturnUrl("https://www.fmm.ink").SetNotifyUrl("https://www.fmm.ink")

	//请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", "沙箱环境")
	bm.Set("out_trade_no", "GZ201901301040355703")
	bm.Set("quit_url", "https://www.fmm.ink")
	bm.Set("total_amount", "100.00")
	bm.Set("product_code", "QUICK_WAP_WAY")

	// 发送支付请求
	payUrl, err = client.TradeWapPay(ctx, bm)
	if err != nil {
		return "", err
	}

	return payUrl, nil
}
