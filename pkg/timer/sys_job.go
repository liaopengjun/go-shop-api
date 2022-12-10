package timer

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"go-shop-api/global"
	"time"
)

var tm = NewTimer()

// JobCore 核心参数函数
type JobCore struct {
	InvokeTarget   string
	Name           string
	JobId          int
	EntryId        int
	CronExpression string
	Args           string
}

// HttpJob 接口类型
type HttpJob struct {
	JobCore
}

// FuncJob 函数类型
type FuncJob struct {
	JobCore
}

// Job 添加任务
type Job interface {
	Run()
	addJob() (int, error)
}

// HttpJob Run 接口请求
func (h *HttpJob) Run() {
	fmt.Printf("%s [RunJob] HttpJob %s exec success \n", time.Now().Format(global.TIME_FORMAT), h.Name)
}

// FuncJob Run 函数运行
func (f *FuncJob) Run() {
	startTime := time.Now()
	jobF := FuncExecList[f.InvokeTarget]
	if jobF == nil {
		return
	}
	err := CallExec(jobF.(JobsExec), f.Args)
	if err != nil {
		// 如果失败暂停一段时间重试
		fmt.Println(time.Now().Format(global.TIME_FORMAT), " [ERROR] mission failed! ", err)
	}
	// 结束时间
	endTime := time.Now()
	// 执行时间
	latencyTime := endTime.Sub(startTime)
	fmt.Printf("%s [RunJob] FuncJob %s exec success Time: %s \n", time.Now().Format(global.TIME_FORMAT), f.Name, latencyTime)
}

func AddJob(job Job) (int, error) {
	if job == nil {
		return 0, nil
	}
	return job.addJob()
}

//  addJob 实现添加接口任务
func (h *HttpJob) addJob() (int, error) {
	EntryId, err := tm.AddTaskByJob(h.Name, h.CronExpression, h, cron.WithSeconds())
	if err != nil {
		fmt.Printf("TaskName: %s Time: %s  Error:%s ", h.Name, time.Now().Format(global.TIME_FORMAT), err)
		return 0, err
	}
	return int(EntryId), err
}

// addJob 实现添加系统函数任务
func (f *FuncJob) addJob() (int, error) {
	EntryId, err := tm.AddTaskByJob(f.Name, f.CronExpression, f, cron.WithSeconds())
	if err != nil {
		fmt.Printf("TaskName: %s Time: %s  Error:%s ", f.Name, time.Now().Format(global.TIME_FORMAT), err)
		return 0, err
	}
	return int(EntryId), err
}

func RemoveJob(name string, entryID int) chan bool {
	return tm.Remove(name, entryID)
}
