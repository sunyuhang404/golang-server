package controller

import (
	"bs-server/src/controller/auth"
	v1 "bs-server/src/controller/v1"
	"bs-server/src/middleware"
	"bs-server/src/response"
	"bs-server/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	// 创建具有默认中间件的路由
	router := gin.Default()

	// 登录校验中间件
	router.Use(middleware.AuthMiddleWare())

	gin.SetMode(util.RunMode)

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, response.OK.WithData("pong"))
	})

	// 权限接口
	auth.RegisterAuthRouter(router)

	// 接口版本管理，可能需要多个版本的接口并存
	v1Group := v1.RegisterRouter(router)
	v1Group.Use(middleware.JWT())

	return router
}
