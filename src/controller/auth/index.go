package auth

import "github.com/gin-gonic/gin"

func RegisterAuthRouter(router *gin.Engine) {
	group := router.Group("/api/auth")
	{
		// 登录
		group.POST("/login", Login)

		// 注册
		group.POST("/register", Register)

		// 获取用户信息
		group.GET("/user", GetUserInfo)
	}
}
