package main

import (
	"bs-server/src/controller"
	"bs-server/src/model"
	"bs-server/src/util"
	"fmt"
	"net/http"
	"time"
)

type Login struct {
	User    string `form:"username" json:"user" uri:"user" binding:"required"`
	Passord string `form:"password" json:"password" uri:"password" binding:"required"`
}

func main() {
	fmt.Printf("main")

	util.Setup()
	model.InitModel()

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
}
