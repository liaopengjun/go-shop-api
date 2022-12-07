package system

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-shop-api/config"
	"go-shop-api/global"
	comReq "go-shop-api/model/common/request"
	"go-shop-api/model/common/response"
	"go-shop-api/model/system/request"
	Response "go-shop-api/model/system/response"
	"go.uber.org/zap"
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
		response.ResponseError(c, config.CodeServerBusy)
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
		response.ResponseError(c, config.CodeServerBusy)
		return
	}
	response.ResponseSuccess(c, jobInfo)
}

// StartJobService 启动计划任务
func (j *JobApi) StartJobService(c *gin.Context) {

}

// StopJobService 停止计划任务
func (j *JobApi) StopJobService(c *gin.Context) {

}
