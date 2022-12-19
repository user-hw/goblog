package main

import (
	"fmt"
	"goblog/common"
	"goblog/router"
	"log"
	"net/http"
)

func init() {
	// 模板加载
	common.LoadTemplate()
}

func main() {
	fmt.Println("hello")
	serve := http.Server{
		Addr: "localhost:8080",
	}

	// 路由功能
	router.Router()

	if err := serve.ListenAndServe(); err != nil {
		log.Print(err)
	}
}
