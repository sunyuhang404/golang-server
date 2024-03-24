package v1

import "github.com/gin-gonic/gin"

func RegisterRouter(router *gin.Engine) *gin.RouterGroup {
	// 接口版本管理
	group := router.Group("/api/v1")
	{
		// 查询 npc 列表
		group.GET("/npc/list", GetNpcList)

		group.GET("/npc/info", GetNpcInfo)
	}

	return group
}
