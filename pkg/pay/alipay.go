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

func (ali *AliPay) Pay(payType string, PayData map[string]interface{}) (payResult map[string]interface{}, err error) {
	switch payType {
	case "WapPay":
		payResult, err = TradeWapPay(PayData)
	}
	return
}

// TradeWapPay 手机Wap支付
func TradeWapPay(PayData map[string]interface{}) (payResult map[string]interface{}, err error) {
	// 建立连接
	client, err := alipay.NewClient(global.GA_CONFIG.AppId, global.GA_CONFIG.PrivateKey, false)
	if err != nil {
		return nil, err
	}

	//配置公共参数
	client.SetCharset("utf-8").SetSignType(alipay.RSA2).SetNotifyUrl(PayData["notify_url"].(string))

	//请求参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", PayData["subject"])
	bm.Set("out_trade_no", PayData["order_code"])
	//bm.Set("quit_url", "https://www.fmm.ink")
	bm.Set("total_amount", PayData["total_amount"])
	bm.Set("product_code", "QUICK_WAP_WAY")

	// 发送支付请求
	payUrl, err := client.TradeWapPay(ctx, bm)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	result["payUrl"] = payUrl
	return result, nil
}
