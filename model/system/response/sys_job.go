package response

import "errors"

var (
	ErrorJobExit    = errors.New("任务已存在")
	ErrorJobNotExit = errors.New("任务不存在")
)
