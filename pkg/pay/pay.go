package pay

// Payer 包含支付方法的接口类型
type Payer interface {
	Pay(payType string, data map[string]interface{}) (map[string]interface{}, error)
}

// NewPay 初始化支付
func NewPay(pay_operator string) Payer {
	switch pay_operator {
	case "ali_pay":
		return &AliPay{}
	case "wx_pay":
		return &WechatPay{}
	default:
		return &AliPay{}
	}
}
