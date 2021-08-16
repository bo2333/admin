package middleware

import (
	"boadmin/server/admin/account"
	"boadmin/server/middleware/session"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Access 权限验证
func Access(access uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		sacce, err := session.Get(c, "access")
		if err != nil {
			return
		}

		if !account.IsAccess(sacce.(uint), access) {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"Status": "无权限禁止访问",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}