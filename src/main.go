package main

import (
	"bs-server/src/controller"
	"bs-server/src/response"
	"bs-server/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Login struct {
	User    string `form:"username" json:"user" uri:"user" binding:"required"`
	Passord string `form:"password" json:"password" uri:"password" binding:"required"`
}

func main() {
	fmt.Printf("main")

	//util.Init()

	router := controller.InitRouter()

	// 自定义服务
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8080),
		Handler:        router,
		ReadTimeout:    util.ReadTimeout,
		WriteTimeout:   util.WriteTimeout,
		IdleTimeout:    60 * time.Second, // 等待最大时间
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

	router.GET("/login", func(context *gin.Context) {
		var json Login
		// 将request中的Body中的数据按照JSON格式解析到json变量中
		if err := context.ShouldBindJSON(&json); err != nil {
			context.JSON(http.StatusBadRequest, response.Error)
			return
		}

		if json.User != "aaa" || json.Passord != "123" {
			context.JSON(http.StatusUnauthorized, response.AuthFail)
			return
		}

		context.JSON(http.StatusOK, response.OK.WithData(json))
	})
}
