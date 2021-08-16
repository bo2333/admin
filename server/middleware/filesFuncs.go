package middleware

import (
	"boadmin/server/admin/account"
	"html/template"
	"time"
)

var GlobalTemplateFun template.FuncMap

func init() {
	GlobalTemplateFun = template.FuncMap{
		"formatAsDate": func(t time.Time, format string) string {
			return t.Format(format)
		},
		// 判断权限
		"isAccess": func(uAccess uint, hAccess uint) bool {
			return account.IsAccess(uAccess, hAccess)
		},
		"attr":func(s string) template.HTMLAttr{
			return template.HTMLAttr(s)
		},
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
	}
}
