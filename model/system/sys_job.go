package system

import (
	"go-shop-api/global"
	"go-shop-api/model/system/response"
	"gorm.io/gorm"
	"time"
)

type SysJob struct {
	JobId          int            `json:"jobId" gorm:"primaryKey;autoIncrement"`          // 编码
	JobName        string         `json:"jobName" gorm:"size:255;comment:名称"`             // 名称
	JobType        int            `json:"jobType" gorm:"size:1;comment:任务类型"`             // 任务类型
	CronExpression string         `json:"cronExpression" gorm:"size:255;comment:cron表达式"` // cron表达式
	InvokeTarget   string         `json:"invokeTarget" gorm:"size:255;comment:调用目标"`      // 调用目标
	Args           string         `json:"args" gorm:"size:255;comment:目标参数"`              // 目标参数
	MisfirePolicy  int            `json:"misfirePolicy" gorm:"size:255;comment:执行策略"`     // 执行策略
	Concurrent     int            `json:"concurrent" gorm:"size:1;comment:是否并发"`          // 是否并发
	Status         int            `json:"status" gorm:"size:1;comment:状态"`                // 状态
	EntryId        int            `json:"entry_id" gorm:"size:11;comment:job启动时返回的id"`    // job启动时返回的id
	CreateBy       uint           `json:"createBy" gorm:"index;comment:创建者"`
	UpdateBy       uint           `json:"updateBy" gorm:"index;comment:更新者"`
	CreatedAt      time.Time      `json:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt      time.Time      `json:"updatedAt" gorm:"comment:最后更新时间"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index;comment:删除时间"`
}

func ExitJob(jobId int, InvokeTarget, CronExpression string) error {
	var jobInfo = new(SysJob)
	db := global.GA_DB.Model(SysJob{})
	if jobId > 0 {
		db.Where("job_id != ?", jobId)
	}
	if db.Where("invoke_target = ? and cron_expression = ?", InvokeTarget, CronExpression).Find(&jobInfo).RowsAffected >= 1 {
		return response.ErrorJobExit
	}
	return nil
}

func AddJob(job *SysJob) (err error) {
	return global.GA_DB.Create(&job).Error
}

func EditJob(job *SysJob) (err error) {
	return global.GA_DB.Updates(job).Error
}

func DelJob(jobId int) (err error) {
	var job = new(SysJob)
	return global.GA_DB.Where("job_id = ?", jobId).Delete(&job).Error
}

func GetJobInfo(jobId int) (job *SysJob, err error, count int64) {
	result := global.GA_DB.Where("job_id = ? ", jobId).First(&job)
	err = result.Error
	count = result.RowsAffected
	return
}

func GetJobList(jobName string, status, jobType int, page, limit int) (jobList []SysJob, count int64, err error) {
	db := global.GA_DB.Model(SysJob{})
	if jobName != "" {
		db.Where("job_name = ?", jobName)
	}
	if status > 0 {
		db.Where("status = ?", status)
	}
	if jobType > 0 {
		db.Where("job_type = ?", jobType)
	}
	err = db.Where("deleted_at is NULL").Count(&count).Error
	offset := (page - 1) * limit
	err = db.Where("deleted_at is NULL").Limit(limit).Offset(offset).Order("job_id desc").Find(&jobList).Error
	return
}
