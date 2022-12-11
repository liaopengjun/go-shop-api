package timer

import (
	"go-shop-api/pkg/redis"
)

//注册计划任务方法
var FuncExecList = map[string]JobsExec{
	"ClearTokenBackList": ClearTokenBackList{},
}

type JobsExec interface {
	Exec(arg interface{}) error
}

func CallExec(e JobsExec, arg interface{}) error {
	return e.Exec(arg)
}

// TODO： 按照下面定义函数执行计划
type ClearTokenBackList struct {
}

func (e ClearTokenBackList) Exec(arg interface{}) error {
	return redis.ClearTokenBlackList()
}
