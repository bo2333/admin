package util

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Res struct {
}

const (
	Success         = 0 // 成功
	Fail			= 1 // 失败
	Param			= 2 // 参数错误
	CONF			= 3 // 配置错误

	DbSave			= 50 // 数据保存出错

	Acct			= 100 // 帐号不存在
	PassLen			= 110 // 密码长度
	Pass			= 111 // 密码不匹配

	Login			= 200 // 没有登录


)

// Html 验证auth后的使用
func (Base *Res) Html(c *gin.Context, url string, obj map[string]interface{}) {
	skin, err := c.Get("skinId")
	if err {
		obj["skinId"] = skin
	}
	c.HTML(http.StatusOK, url, obj)
}


func (Base *Res) Json(c *gin.Context, code int, msg interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": msg,
	})
}

func (Base *Res) Table(c *gin.Context, code int, count int64, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"count": count,
		"msg": msg,
		"data": data,
	})
}