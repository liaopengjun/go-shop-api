package request

type JobParam struct {
	JobId          int    `json:"jobId"`
	JobName        string `json:"jobName" binding:"required"`
	JobType        int    `json:"jobType" binding:"required"`
	CronExpression string `json:"cronExpression" binding:"required"`
	InvokeTarget   string `json:"invokeTarget" binding:"required"`
	Args           string `json:"args"`
	Status         int    `json:"status" binding:"required,numeric"`
}

type JobListParam struct {
	JobName string `json:"job_name"`
	JobType int    `json:"job_type"`
	Status  int    `json:"status"`
	Page    int    `json:"page" binding:"numeric"`
	Limit   int    `json:"limit" binding:"required,numeric"`
}
