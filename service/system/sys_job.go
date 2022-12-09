package system

import (
	"go-shop-api/model/system"
	"go-shop-api/model/system/request"
	"go-shop-api/model/system/response"
	"time"
)

type JobService struct {
}

func (j *JobService) AddJob(userId uint, p *request.JobParam) (err error) {
	//校验是否已添加
	err = system.ExitJob(p.JobId, p.InvokeTarget, p.CronExpression, p.Args)
	if err != nil {
		return
	}
	jobData := system.SysJob{
		JobName:        p.JobName,
		JobType:        p.JobType,
		CronExpression: p.CronExpression,
		InvokeTarget:   p.InvokeTarget,
		Args:           p.Args,
		MisfirePolicy:  p.MisfirePolicy,
		Concurrent:     p.Concurrent,
		Status:         p.Status,
		CreateBy:       userId,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	return system.AddJob(&jobData)
}

func (j *JobService) EditJob(userId uint, p *request.JobParam) (err error) {

	//校验是否已存在
	err = system.ExitJob(p.JobId, p.InvokeTarget, p.CronExpression, p.Args)
	if err != nil {
		return err
	}

	jobData := system.SysJob{
		JobId:          p.JobId,
		JobName:        p.JobName,
		JobType:        p.JobType,
		CronExpression: p.CronExpression,
		InvokeTarget:   p.InvokeTarget,
		Args:           p.Args,
		MisfirePolicy:  p.MisfirePolicy,
		Concurrent:     p.Concurrent,
		Status:         p.Status,
		UpdateBy:       userId,
		UpdatedAt:      time.Now(),
	}
	return system.EditJob(&jobData)
}

func (j *JobService) EditJobEntry(userId uint, jobId, entryId int) error {
	jobData := system.SysJob{
		JobId:     jobId,
		EntryId:   entryId,
		UpdateBy:  userId,
		UpdatedAt: time.Now(),
	}
	return system.EditJob(&jobData)
}

func (j *JobService) DelJob(jobId int) error {
	// 校验是否存在
	jobInfo, err, total := system.GetJobInfo(jobId)
	if err != nil || total == 0 {
		return response.ErrorJobNotExit
	}
	return system.DelJob(jobInfo.JobId)
}

func (j *JobService) GetJobInfo(jobId int) (jobInfo *system.SysJob, err error) {
	jobInfo, err, total := system.GetJobInfo(jobId)
	if err != nil || total == 0 {
		return nil, response.ErrorJobNotExit
	}
	return
}

func (j *JobService) GetJobList(p *request.JobListParam) (jobList []system.SysJob, total int64, err error) {
	return system.GetJobList(p.JobName, p.Status, p.JobType, p.Page, p.Limit)
}
