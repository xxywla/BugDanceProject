package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func initSession(r *gin.Engine) {
	// redis.NewStore(redis最大空闲连接数, tcp/udp, address, password, key)
	// store, err := redis.NewStore(100, "tcp", "localhost:6379", "", []byte("session"))
	store := cookie.NewStore([]byte("secret"))

	// if err != nil {
	// 	fmt.Print("Unable to create session container")
	// 	return
	// }

	r.Use(sessions.Sessions("mysession", store))
}
