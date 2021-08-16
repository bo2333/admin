package session

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)


// SBase sessions 内容 修改本处，要修改 frame.html 的 getPublicData
type SBase struct {
	Id    uint `json:"id"`
	Name   string `json:"name"`
	Access    uint `json:"access"`
	Utm		int64 `json:"utm"`
	Skin	uint8 `json:"skin"`
}

type SInfo struct {
	SBase
	Acct 	string `json:"acct"`
	Pass 	string `json:"pass"`
}

func Set(c *gin.Context, acct string, pass string, id uint, name string, access uint, skin uint8) bool {
	//session 存储数据
	gob.Register(SInfo{})
	sess := sessions.Default(c)
	sess.Set("id", id)
	sess.Set("name", name)
	sess.Set("skin", skin)
	sess.Set("access", access)
	sess.Set("utm", time.Now().Unix())
	sess.Set("acct", acct)
	sess.Set("pass", pass)
	err := sess.Save()
	if err != nil {
		return false
	}
	return true
}

func Mod(c *gin.Context, key string, val interface{})  bool {
	sess := sessions.Default(c)
	sess.Set(key, val)
	err:= sess.Save()
	if err != nil {
		return false
	}
	return true
}

// Get 前端读取
func Get(c *gin.Context, key string) (interface{}, interface{}) {
	href := "/login"
	sess := sessions.Default(c)
	val := sess.Get(key)
	if val == nil {
		// 取不到就是没有登录
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<script type="text/javascript">top.location.href="%s"</script>`, href)
		return 0, 1
	}

	return val, nil
}

func GetBase(c *gin.Context) (SBase, interface{}) {
	href := "/login"
	sess := sessions.Default(c)
	id := sess.Get("id")
	var info SBase
	if id == nil {
		// 取不到就是没有登录
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<script type="text/javascript">top.location.href="%s"</script>`, href)
		return info, 1
	}

	info = SBase{
		Id: id.(uint),
		Name: sess.Get("name").(string),
		Access: sess.Get("access").(uint),
		Utm: sess.Get("utm").(int64),
		Skin: sess.Get("skin").(uint8),
	}

	return info, nil
}


// GetInfo 后端读取
func GetInfo(c *gin.Context) (SInfo, interface{}) {
	href := "/login"
	sess := sessions.Default(c)

	id := sess.Get("id")
	var info SInfo
	if id == nil {
		// 取不到就是没有登录
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<script type="text/javascript">top.location.href="%s"</script>`, href)
		return info, 1
	}

	info.Id =  id.(uint)
	info.Name = sess.Get("name").(string)
	info.Access = sess.Get("access").(uint)
	info.Utm = sess.Get("utm").(int64)
	info.Skin = sess.Get("skin").(uint8)
	info.Acct = sess.Get("acct").(string)
	info.Pass = sess.Get("pass").(string)

	return info, nil
}