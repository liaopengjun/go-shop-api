package system

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-shop-api/config"
	"go-shop-api/global"
	comReq "go-shop-api/model/common/request"
	"go-shop-api/model/common/response"
	"go-shop-api/model/system/request"
	Response "go-shop-api/model/system/response"
	"go-shop-api/pkg/timer"
	"go.uber.org/zap"
	"time"
)

type JobApi struct {
}

// GetJobList 计划任务
func (j *JobApi) GetJobList(c *gin.Context) {
	var p = new(request.JobListParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("请求计划列表参数有误", zap.Error(err))
		//判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ResponseError(c, config.CodeInvalidParam)
			return
		}
		//自定义错误
		response.ResponseErrorWithMsg(c, config.CodeInvalidParam, RemoveTopStructNew(errs.Translate(global.GA_TRANS)))
		return
	}
	list, total, err := jobService.GetJobList(p)
	if err != nil {
		global.GA_LOG.Error("获取计划任务列表失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	//4.处理返回
	response.ResponseSuccess(c, response.PageResult{
		List:  list,
		Total: total,
		Page:  p.Page,
		Limit: p.Limit,
	})
}

// AddJob  添加计划任务
func (j *JobApi) AddJob(c *gin.Context) {
	var p = new(request.JobParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("添加计划参数有误", zap.Error(err))
		//判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ResponseError(c, config.CodeInvalidParam)
			return
		}
		//自定义错误
		response.ResponseErrorWithMsg(c, config.CodeInvalidParam, RemoveTopStructNew(errs.Translate(global.GA_TRANS)))
		return
	}
	userId, _ := c.Get("userid")
	err := jobService.AddJob(userId.(uint), p)
	if err != nil {
		global.GA_LOG.Error("添加计划任务失败", zap.Error(err))
		if errors.Is(err, Response.ErrorJobExit) {
			response.ResponseError(c, config.CodeJobExitError)
		} else {
			response.ResponseError(c, config.CodeServerBusy)
		}
		return
	}
	response.ResponseSuccess(c, "添加任务成功")
}

//DelJob 删除计划任务
func (j *JobApi) DelJob(c *gin.Context) {
	var p = new(comReq.GetById)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("删除计划参数有误", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	err := jobService.DelJob(p.ID)
	if err != nil {
		if errors.Is(err, Response.ErrorJobNotExit) {
			response.ResponseError(c, config.CodeJobNotExitError)
		} else {
			response.ResponseError(c, config.CodeServerBusy)
		}
		return
	}
	response.ResponseSuccess(c, "删除计划任务成功")
}

// EditJob 编辑任务详情
func (j *JobApi) EditJob(c *gin.Context) {
	var p = new(request.JobParam)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("编辑计划参数有误", zap.Error(err))
		//判断err是不是validator类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			response.ResponseError(c, config.CodeInvalidParam)
			return
		}
		//自定义错误
		response.ResponseErrorWithMsg(c, config.CodeInvalidParam, RemoveTopStructNew(errs.Translate(global.GA_TRANS)))
		return
	}
	userId, _ := c.Get("userid")
	err := jobService.EditJob(userId.(uint), p)
	if err != nil {
		global.GA_LOG.Error("编辑计划任务失败", zap.Error(err))
		if errors.Is(err, Response.ErrorJobExit) {
			response.ResponseError(c, config.CodeJobExitError)
		} else {
			response.ResponseError(c, config.CodeServerBusy)
		}
		return
	}
	response.ResponseSuccess(c, "编辑任务成功")

}

// GetJobInfo 计划任务详情
func (j *JobApi) GetJobInfo(c *gin.Context) {
	var p = new(comReq.GetById)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("获取计划详情参数有误", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}
	jobInfo, err := jobService.GetJobInfo(p.ID)
	if err != nil {
		if errors.Is(err, Response.ErrorJobNotExit) {
			response.ResponseError(c, config.CodeJobNotExitError)
		} else {
			response.ResponseError(c, config.CodeServerBusy)
		}
		return
	}
	response.ResponseSuccess(c, jobInfo)
}

// StartJobService 启动计划任务
func (j *JobApi) StartJobService(c *gin.Context) {
	var p = new(comReq.GetById)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("获取计划详情参数有误", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}

	jobInfo, err := jobService.GetJobInfo(p.ID)
	if err != nil {
		if errors.Is(err, Response.ErrorJobNotExit) {
			response.ResponseError(c, config.CodeJobNotExitError)
		} else {
			response.ResponseError(c, config.CodeServerBusy)
		}
		return
	}

	//判断当前状态是否开启
	if jobInfo.Status == 2 {
		response.ResponseError(c, config.CodeJobStartError)
		return
	}

	// 添加任务
	entryId := 0
	if jobInfo.JobType == 1 { //接口
		j := &timer.HttpJob{}
		j.InvokeTarget = jobInfo.InvokeTarget
		j.CronExpression = jobInfo.CronExpression
		j.JobId = jobInfo.JobId
		j.Name = jobInfo.JobName
		entryId, err = timer.AddJob(j)
	} else { //函数
		j := &timer.FuncJob{}
		j.InvokeTarget = jobInfo.InvokeTarget
		j.CronExpression = jobInfo.CronExpression
		j.JobId = jobInfo.JobId
		j.Name = jobInfo.JobName
		j.Args = jobInfo.Args
		entryId, err = timer.AddJob(j)
	}

	if err != nil {
		global.GA_LOG.Error("启动计划任务失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}

	fmt.Printf("TaskName: %s Time: %s  StartJob success \n", jobInfo.JobName, time.Now().Format(global.TIME_FORMAT))

	// 更新任务id
	userId, _ := c.Get("userid")
	err = jobService.EditJobEntry(userId.(uint), jobInfo.JobId, entryId)
	if err != nil {
		global.GA_LOG.Error("更新计划任务job失败", zap.Error(err))
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, "启动任务成功")

}

// StopJobService 停止计划任务
func (j *JobApi) RemoveJobService(c *gin.Context) {
	var p = new(comReq.GetById)
	if err := c.ShouldBindJSON(p); err != nil {
		global.GA_LOG.Error("获取计划详情参数有误", zap.Error(err))
		response.ResponseError(c, config.CodeInvalidParam)
		return
	}

	jobInfo, err := jobService.GetJobInfo(p.ID)
	if err != nil {
		if errors.Is(err, Response.ErrorJobNotExit) {
			response.ResponseError(c, config.CodeJobNotExitError)
		} else {
			response.ResponseError(c, config.CodeServerBusy)
		}
		return
	}

	//判断当前状态是否开启
	if jobInfo.Status == 2 {
		response.ResponseError(c, config.CodeJobStopError)
		return
	}

	ch := timer.RemoveJob(jobInfo.JobName, jobInfo.EntryId)
	select {
	case res := <-ch:
		if res {
			fmt.Printf("TaskName: %s Time: %s  StopJob success \n ", jobInfo.JobName, time.Now().Format(global.TIME_FORMAT))
			//更新计划任务jobId
			userId, _ := c.Get("userid")
			err = jobService.EditJobEntry(userId.(uint), jobInfo.JobId, 0)
			if err != nil {
				global.GA_LOG.Error("更新计划任务job失败", zap.Error(err))
				response.ResponseError(c, config.CodeServerBusy)
				return
			}
		}
	case <-time.After(time.Second * 1):
		//超时处理
		response.ResponseError(c, config.CodeOperationTimeoutError)
		return
	}

	response.ResponseSuccess(c, "停止任务成功")
}
