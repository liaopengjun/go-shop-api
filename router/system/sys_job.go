package system

import (
	"github.com/gin-gonic/gin"
	v1 "go-shop-api/api/v1"
)

type JobRouter struct {
}

// InitJobRouterRouter 计划任务
func (s *JobRouter) InitJobRouterRouter(Router *gin.RouterGroup) {
	jobRouter := Router.Group("job")
	var JobApi = v1.ApiGroupApp.SystemApiGroup.JobApi
	{
		jobRouter.POST("getJobList", JobApi.GetJobList)    // 任务列表
		jobRouter.POST("getJobInfo", JobApi.GetJobInfo)    // 任务详情
		jobRouter.POST("addJob", JobApi.AddJob)            // 添加任务
		jobRouter.POST("editJob", JobApi.EditJob)          // 编辑任务
		jobRouter.POST("delJob", JobApi.DelJob)            // 删除任务
		jobRouter.POST("startJob", JobApi.StartJobService) // 启动任务
		jobRouter.POST("stopJob", JobApi.StopJobService)   // 停止任务
	}
}
