package system

import (
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
	"go-shop-api/model/system"
	"go-shop-api/model/system/request"
	"go-shop-api/utils"
	"time"
)

type LoginLogService struct {
}

func (l *LoginLogService) CreateLoginLog(c *gin.Context, name string, status string, msg string) (err error) {
	ua := user_agent.New(c.Request.UserAgent())
	browserName, browserVersion := ua.Browser()
	browser := browserName + " " + browserVersion
	loginLogData := system.SysLoginLog{
		Username:      name,
		Status:        status,
		Ipaddr:        c.ClientIP(), // 请求ip
		LoginLocation: "",
		Browser:       browser,
		Os:            ua.OS(),
		Platform:      ua.Platform(),
		LoginTime:     time.Now(),
		Remark:        c.Request.UserAgent(),
		Msg:           msg,
	}
	return system.CreateLoginLog(&loginLogData)
}

func (l *LoginLogService) GetLoginLogList(p *request.ParamLoginLogList) (list interface{}, total int64, err error) {
	list, total, err = system.GetLoginLogList(p.UserName, p.Status, p.Ip, int(p.Page), int(p.Limit))
	return
}

func (l *LoginLogService) DelLoginLog(id string) error {
	ids := utils.StringToSlice(id)
	return system.DelLoginLog(ids)
}
