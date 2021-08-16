package middleware

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func Logs() bool {
	f, err := os.Create("logs/server.log")
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(f)
		gin.DefaultErrorWriter = io.MultiWriter(f)
		return true
	}
	return false
	// gin.Logger(),
	// gin.Recovery()
}