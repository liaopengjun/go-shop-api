package timer

import (
	"fmt"
	"go-shop-api/global"
	"go.uber.org/zap"
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
func (h HttpJob) Run() {

}

// FuncJob Run 函数运行
func (f FuncJob) Run() {

}

func AddJob(job Job) (int, error) {
	if job == nil {
		return 0, nil
	}
	return job.addJob()
}

//  addJob 实现添加接口任务
func (h *HttpJob) addJob() (int, error) {
	job := HttpJob{}
	EntryId, err := tm.AddTaskByJob(h.Name, h.CronExpression, job)
	if err != nil {
		fmt.Printf("TaskName: %s Time: %s  Error:%s ", h.Name, time.Now().Format(global.TIME_FORMAT), zap.Error(err))
		return 0, err
	}
	return int(EntryId), err
}

// addJob 实现添加系统函数任务
func (f *FuncJob) addJob() (int, error) {
	job := FuncJob{}
	EntryId, err := tm.AddTaskByJob(f.Name, f.CronExpression, job)
	if err != nil {
		fmt.Printf("TaskName: %s Time: %s  Error:%s ", f.Name, time.Now().Format(global.TIME_FORMAT), zap.Error(err))
		return 0, err
	}
	return int(EntryId), err
}

func RemoveJob(name string, entryID int) chan bool {
	return tm.Remove(name, entryID)
}
