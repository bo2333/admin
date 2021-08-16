package home

import (
	"boadmin/server/middleware/session"
	util2 "boadmin/server/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LayoutRes struct {
	util2.Res
}

func (ret *LayoutRes) View() gin.HandlerFunc {
	return func(c *gin.Context) {
		skin, err := session.Get(c, "skin")
		if err == nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"skinId": skin,
			})
		} else {
			c.HTML(http.StatusBadRequest, "index.html", nil)
		}
	}
}

func (ret *LayoutRes) Data() gin.HandlerFunc {
	return func(c *gin.Context) {
		info, err := session.GetBase(c)
		if err != nil {
			return
		}
		reqIp := c.ClientIP()
		if reqIp == "::1" {
			reqIp = "127.0.0.1"
		}
		navs := GetNavs(info.Access)
		ret.Json(c, util2.Success,  gin.H{
			"title":        util2.Ini.SLogo.Title,
			"navLogoIcon":  util2.Ini.SLogo.Icon,
			"navLogoTitle": util2.Ini.SLogo.Title,
			"navs":         navs,
			"userInfo":     info,
			"website":      util2.Ini.SBase.Website,
			"email":        util2.Ini.SBase.Email,
			"version":      util2.Ini.SBase.Version,
			"ip":           reqIp,
		})
	}
}