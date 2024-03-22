package middleware

import (
	"bs-server/src/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理权限中间接

func AuthMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取token
		//token := context.Request.Header.Get("Authorization")

		// 调用校验token的函数
		authorizated := true

		if authorizated {
			context.Next()
			return
		} else {
			context.JSON(http.StatusUnauthorized, response.AuthFail)
			context.Abort()
			return
		}
	}
}
