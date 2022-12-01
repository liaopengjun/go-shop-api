package pay

type WechatPay struct {
}

func (ali *WechatPay) Pay(payType string, PayData map[string]interface{}) (payResult map[string]interface{}, err error) {
	switch payType {
	case "WapPay":
		return nil, nil
	}
	return
}
