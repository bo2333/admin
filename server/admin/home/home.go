package home

import (
	util2 "boadmin/server/util"
	"github.com/gin-gonic/gin"
	"runtime"
	"strings"
	"time"
)

type HomeRes struct {
	util2.Res
}


func (ret *HomeRes) HomeView() gin.HandlerFunc {
	return func(c *gin.Context) {
		ret.Html(c, "home.html", gin.H{})
	}
}

func (ret *HomeRes) HomeData() gin.HandlerFunc {
	return func(c *gin.Context) {
		userAgent := c.GetHeader("User-Agent")
		browser := "Other"
		if strings.Contains(userAgent, "MSIE") {
			browser = "MSIE"
		} else if strings.Contains(userAgent, "Firefox") {
			browser = "Firefox"
		} else if strings.Contains(userAgent, "Chrome") {
			browser = "Chrome"
		} else if strings.Contains(userAgent, "Safari") {
			browser = "Safari"
		} else if strings.Contains(userAgent, "Opera") {
			browser = "Opera"
		}

		// 取出版本
		l := strings.Count(browser,"") - 1
		idx := strings.Index(userAgent, browser) + l
		browserVersion := userAgent[idx:]
		idx = strings.Index(userAgent, " ")
		if idx <= -1 {
			browserVersion = userAgent[1:]
		} else {
			browserVersion = userAgent[1:idx]
		}
		now := time.Now()
		zone, _ := now.Zone()
		ret.Json(c, util2.Success, gin.H{
			"ginVersion":     gin.Version,
			"browser":        browser,
			"browserVersion": browserVersion,
			"goVersion":      runtime.Version(),
			"dbVersion":      util2.DbVersion(),
			"timezone":       zone,
			"now":            time.Now().Format("2006-01-02 15:04:05"),
		})
	}
}