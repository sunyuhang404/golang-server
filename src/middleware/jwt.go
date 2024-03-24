package middleware

import (
	"bs-server/src/response"
	"bs-server/src/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {

		token := context.Query("token")

		if token == "" {
			context.JSON(http.StatusOK, response.AuthFail)
		} else {
			claims, err := util.ParserToken(token)

			expirateTime, _ := claims.GetExpirationTime()

			if err != nil {
				context.JSON(http.StatusOK, response.Error.WithData(err))
			} else if time.Now().Unix() > expirateTime.Unix() {
				context.JSON(http.StatusOK, response.Error.WithData("token 过期"))
			}
		}
		context.Next()
	}
}
