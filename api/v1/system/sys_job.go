package system

import (
	"github.com/gin-gonic/gin"
)

type JobApi struct {
}

// GetJobList 计划任务
func (j *JobApi) GetJobList(c *gin.Context) {
}

// AddJob  添加计划任务
func (j *JobApi) AddJob(c *gin.Context) {

}

//DelJob 删除计划任务
func (j *JobApi) DelJob(c *gin.Context) {

}

// EditJob 编辑任务详情
func (j *JobApi) EditJob(c *gin.Context) {

}

// GetJobInfo 计划任务详情
func (j *JobApi) GetJobInfo(c *gin.Context) {

}

// StartJobService 启动计划任务
func (j *JobApi) StartJobService(c *gin.Context) {

}

// StopJobService 停止计划任务
func (j *JobApi) StopJobService(c *gin.Context) {

}
