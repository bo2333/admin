package main

import (
	"boadmin/server"
	"boadmin/server/middleware"
	util2 "boadmin/server/util"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"
)

func loadViews(router *gin.Engine, htmlDir string) {
	router.SetFuncMap(middleware.GlobalTemplateFun) // 自定义模板函数
	var files []string
	filepath.Walk(htmlDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			files = append(files, path)
		}
		return nil
	})

	router.LoadHTMLFiles(files...)

}

func main() {
	// go build -gcflags "all=-N -l" main.go
	// dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./main
	router := server.Route()
	if router == nil {
		os.Exit(0)
	}

	loadViews(router, "./resources")
	router.StaticFS("/resources/statics", http.Dir("./resources/statics"))
	router.StaticFile("/favicon.ico", "./resources/statics/img/favicon.ico")

	url := util2.Ini.SBase.Host + ":" + util2.Ini.SBase.Port
	log.Println(url)
	srv := &http.Server{
		Addr: url,
		Handler: router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("exit ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("exit error:", err)
	}
	log.Println("exiting")
}
