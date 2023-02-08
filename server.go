package main

import (
	"douyinapp/controller"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()
	r.GET("/douyin/user/:id", func(c *gin.Context) {
		userId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {

		}
		data := controller.UserInfo(controller.UserRequest{UserId: userId, Token: "111"})
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
