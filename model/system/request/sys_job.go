package request

type JobParam struct {
	JobId          int    `json:"job_id"`
	JobName        string `json:"job_name"`
	JobType        int    `json:"job_type"`
	CronExpression string `json:"cron_expression"`
	InvokeTarget   string `json:"invoke_target"`
	Args           string `json:"args"`
	MisfirePolicy  int    `json:"misfire_policy"`
	Concurrent     int    `json:"concurrent"`
	Status         int    `json:"status"`
}

type JobListParam struct {
	JobName string `json:"job_name"`
	JobType int    `json:"job_type"`
	Status  int    `json:"status"`
	Page    int    `json:"page" binding:"numeric"`
	Limit   int    `json:"limit" binding:"required,numeric"`
}
