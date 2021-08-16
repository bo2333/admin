package server

import (
	"boadmin/server/admin/account"
	"boadmin/server/admin/home"
	"boadmin/server/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)



func Route() *gin.Engine {
	if !middleware.Logs() {
		return nil
	}

	router := gin.Default()
	router.Delims("{[", "]}")
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: int(30 * 60), // 30min sessions 保存30分钟
		Path:   "/",
	})
	router.Use(sessions.Sessions("mysession", store))

	/*******登录路由**********/
	{
		router.GET("/login", (&home.LoginRes{}).Index())
		router.POST("/login", (&home.LoginRes{}).Login())
	}


	/******* index **********/
	gIndex := router.Group("")
	gIndex.Use( middleware.Auth())
	{
		gIndex.GET("", (&home.LayoutRes{}).View())
		gIndex.GET("/data", (&home.LayoutRes{}).Data())
		gIndex.GET("/home", (&home.HomeRes{}).HomeView())
		gIndex.GET("/home-data", (&home.HomeRes{}).HomeData())
	}


	gAccount := router.Group("/account")
	gAccount.Use( middleware.Auth())
	{
		gAccount.GET("", middleware.Access(account.AccountManage|account.AccountRead), (&account.AccountRes{}).Index())
		gAccount.GET("/page", middleware.Access(account.AccountManage|account.AccountRead), (&account.AccountRes{}).Page())
		gAccount.GET("/group", middleware.Access(account.AccountManage|account.AccountRead), (&account.AccountRes{}).Group())
		gAccount.POST("/change-password",  (&account.AccountRes{}).ChangePassword())
		gAccount.POST("/change-skin",  (&account.AccountRes{}).ChangeSkin())
		gAccount.POST("/modify", middleware.Access(account.AccountManage), (&account.AccountRes{}).Modify())
		gAccount.POST("/reset-password",  (&account.AccountRes{}).ResetPassword())


	}

	gActionLog := router.Group("/action")
	gActionLog.Use( middleware.Auth())
	{
		gActionLog.GET("", middleware.Access(account.AccountManage|account.AccountRead), (&account.ActionLogRes{}).Index())
		gActionLog.GET("/data", middleware.Access(account.AccountManage|account.AccountRead), (&account.ActionLogRes{}).Data())
		gActionLog.GET("/page", middleware.Access(account.AccountManage|account.AccountRead), (&account.ActionLogRes{}).Page())
	}

	/******* home **********/
	/*
	gHome := router.Group("/home")
	gHome.Use(middleware.Auth())
	{
		gHome.GET("", (&home.HomeRes{}).View())
		gHome.GET("/index", (&home.HomeRes{}).Home())
	}
	*/

	return router
}