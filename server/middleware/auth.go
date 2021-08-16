package middleware

import (
	"boadmin/server/middleware/session"
	"github.com/gin-gonic/gin"
)


// Auth 登录验证
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		skin, err := session.Get(c, "skin")
		if err != nil {
			return
		}
		c.Set("skinId", skin.(uint8))
		c.Next()
	}
}