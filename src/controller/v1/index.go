package v1

import "github.com/gin-gonic/gin"

func RegisterRouter(router *gin.Engine) {
	// 接口版本管理
	group := router.Group("/api/v1")
	{
		group.GET("/model/list")
	}
}
