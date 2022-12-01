package enum

func GetNewBeeMallOrderStatusEnumByStatus(status int) (int, string) {
	switch status {
	case 0:
		return 0, "待支付"
	case 1:
		return 1, "已支付"
	case 2:
		return 2, "配货完成"
	case 3:
		return 3, "出库成功"
	case 4:
		return 4, "交易成功"
	case -1:
		return -1, "手动关闭"
	case -2:
		return -2, "超时关闭"
	case -3:
		return -3, "商家关闭"
	default:
		return -9, "error"
	}
}

func GetNewBeeMallOrderPayTypeEnumByStatus(status int) (int, string) {
	switch status {
	case 0:
		return 0, "无"
	case 1:
		return 1, "支付宝"
	case 2:
		return 1, "微信"
	default:
		return -1, "error"
	}
}
