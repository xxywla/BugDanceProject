package main

import (
	"douyinapp/repository"
	"douyinapp/service"

	"github.com/gin-gonic/gin"
)

func main() {
	go service.RunMessageServer()

	r := gin.Default()

	initRouter(r)

	repository.Init()

	service.NewConsumer()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
