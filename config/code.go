package config

type ResCode int64

//定义返回状态
const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidToken
	CodeMenuExist
	CodeMenuChildExist
	CodeAuthExist
	CodeAuthChildExit
	CodeAuthApiExit
	CodeNoPermission
	CodeFileError
	CodeLowerShelfError
	CodeCreateOrderError
	CodePayOrderError
	CodeJobExitError
	CodeJobNotExitError
	CodeJobStartError
	CodeJobStopError
	CodeRequestTimeoutError
	CodeOperationTimeoutError
	CodeJobInServiceError
)

//定义返回信息
var codeMsgMap = map[ResCode]string{
	CodeSuccess:               "success",
	CodeInvalidParam:          "请求参数有误",
	CodeUserExist:             "用户存在",
	CodeUserNotExist:          "用户不存在",
	CodeInvalidPassword:       "用户名或密码错误",
	CodeServerBusy:            "服务繁忙",
	CodeInvalidToken:          "登陆信息已过期",
	CodeNeedLogin:             "请先登录用户",
	CodeMenuExist:             "菜单已存在",
	CodeMenuChildExist:        "菜单存在下级",
	CodeAuthExist:             "角色已存在",
	CodeAuthChildExit:         "角色存在下级",
	CodeAuthApiExit:           "api已存在",
	CodeNoPermission:          "权限不足",
	CodeFileError:             "文件接受失败",
	CodeLowerShelfError:       "商品已下架",
	CodeCreateOrderError:      "创建订单失败",
	CodePayOrderError:         "支付订单失败",
	CodeJobExitError:          "计划任务已存在",
	CodeJobNotExitError:       "计划任务不存在",
	CodeJobStartError:         "当前任务关闭状态不能被启用",
	CodeJobStopError:          "当前任务关闭状态不能关闭",
	CodeRequestTimeoutError:   "请求超时",
	CodeOperationTimeoutError: "操作超时",
	CodeJobInServiceError:     "任务运行中无法删除",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		return codeMsgMap[CodeServerBusy]
	}
	return msg
}
