package timer

import (
	"fmt"
	"go-shop-api/global"
	"time"
)

//注册计划任务方法
var FuncExecList = map[string]JobsExec{
	"ExamplesFun": ExamplesFun{},
}

type JobsExec interface {
	Exec(arg interface{}) error
}

func CallExec(e JobsExec, arg interface{}) error {
	return e.Exec(arg)
}

// TODO： 按照下面定义函数执行计划
type ExamplesFun struct {
}

func (e ExamplesFun) Exec(arg interface{}) error {
	str := time.Now().Format(global.TIME_FORMAT) + " [INFO] JobCore ExamplesOne exec success"
	switch arg.(type) {
	case string:
		if arg.(string) != "" {
			fmt.Println("string", arg.(string))
			fmt.Println(str, arg.(string))
		} else {
			fmt.Println("arg is nil")
			fmt.Println(str, "arg is nil")
		}
		break
	}
	return nil
}
