package request

type JobParam struct {
	JobId          int    `json:"job_id"`
	JobName        string `json:"job_name" binding:"required"`
	JobType        int    `json:"job_type" binding:"required"`
	CronExpression string `json:"cron_expression" binding:"required"`
	InvokeTarget   string `json:"invoke_target" binding:"required"`
	Args           string `json:"args"`
	MisfirePolicy  int    `json:"misfire_policy" binding:"required,numeric"`
	Concurrent     int    `json:"concurrent" binding:"required,numeric"`
	Status         int    `json:"status" binding:"required,numeric"`
}

type JobListParam struct {
	JobName string `json:"job_name"`
	JobType int    `json:"job_type"`
	Status  int    `json:"status"`
	Page    int    `json:"page" binding:"numeric"`
	Limit   int    `json:"limit" binding:"required,numeric"`
}
