package system

import (
	"gorm.io/gorm"
	"time"
)

type SysJob struct {
	JobId          int            `json:"jobId" gorm:"primaryKey;autoIncrement"`          // 编码
	JobName        string         `json:"jobName" gorm:"size:255;comment:名称"`             // 名称
	JobGroup       string         `json:"jobGroup" gorm:"size:255;comment:任务分组"`          // 任务分组
	JobType        int            `json:"jobType" gorm:"size:1;comment:任务类型"`             // 任务类型
	CronExpression string         `json:"cronExpression" gorm:"size:255;comment:cron表达式"` // cron表达式
	InvokeTarget   string         `json:"invokeTarget" gorm:"size:255;comment:调用目标"`      // 调用目标
	Args           string         `json:"args" gorm:"size:255;comment:目标参数"`              // 目标参数
	MisfirePolicy  int            `json:"misfirePolicy" gorm:"size:255;comment:执行策略"`     // 执行策略
	Concurrent     int            `json:"concurrent" gorm:"size:1;comment:是否并发"`          // 是否并发
	Status         int            `json:"status" gorm:"size:1;comment:状态"`                // 状态
	EntryId        int            `json:"entry_id" gorm:"size:11;comment:job启动时返回的id"`    // job启动时返回的id
	DataScope      string         `json:"dataScope" gorm:"-"`
	CreateBy       int            `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy       int            `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt      time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt      time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}
