package home

import (
	"boadmin/server/admin/account"
	"boadmin/server/middleware/session"
	util2 "boadmin/server/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRes struct {
	util2.Res
}


func (ret *LoginRes) Index() gin.HandlerFunc {
	return func(c *gin.Context) {
		ret.Html(c, "login.html", gin.H{
			"logo_ico":   util2.Ini.SLogo.Icon,
			"logo_title": util2.Ini.SLogo.Title,
		})
	}

}

func (ret *LoginRes) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		acct := c.PostForm("acct")
		password := c.PostForm("password")
		adminUser, err := account.InfoByAcct(acct)
		if err != nil {
			ret.Json(c, util2.Acct, "账号不存在")
			return
		}
		//判断密码是否正确
		pass := util2.Encryption(acct, password, util2.Ini.SBase.PasswordSalt)
		if  pass == adminUser.Password {
			// 保存 Sessions
			if !session.Set(c, acct, pass, adminUser.Id, adminUser.Name, adminUser.Access, adminUser.Skin) {
				return
			}
			_ = account.UpdateLoginHost(adminUser.Id, c.ClientIP())
			// c.Redirect(http.StatusMovedPermanently,"/")
			ret.Json(c, util2.Success, "/")
		} else {
			ret.Json(c, util2.Pass, "账号密码错误")
		}
	}
}

func (ret *LoginRes) LoginOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		ss := sessions.Default(c)
		ss.Delete("userInfo")
		err := ss.Save()
		if err != nil {
			return
		}
		c.Redirect(http.StatusFound, "/login")
	}
}
