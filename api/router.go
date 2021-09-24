package api

import (
	"server/api/controller"
	"server/api/middleware"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {

	r.Use(middleware.ResponseHanlder())
	//註冊：已被註冊 成功（記錄至Slice中（acc pwd id））PWD有md5加密 done, sql added
	r.POST("/registry", controller.Registry)
	//登入：成功（會記錄acc 登入時間unix） 失敗 未註冊 有md5比對, 有JWT產生 done, sql added
	r.POST("/login", controller.Login)
	//登出：JWT比對 JWT刪除, sql added
	r.POST("/logout", controller.Logout)
	// sql added
	r.POST("/chpwd", controller.ChangePwd)
	//使用者login log刪除：刪除成功 輸入不存在使用者會回傳找不到, sql added
	r.DELETE("/logdel", controller.Logdel)
	//使用者查詢：回傳id,帳號,JWT 查詢不存在會回傳不存在, sql added, redis added
	r.GET("/query/:account", controller.Query)
	//使用者login log查詢：回傳帳號,曾登入時間 查詢不存在使用者會回傳找不到, sql added
	r.GET("/memberlog", controller.Memberlog) //查詢會員登入記錄
	r.NoRoute(controller.NotFound)
}
