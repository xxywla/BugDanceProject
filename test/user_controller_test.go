package test

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func TestUserInfo(t *testing.T) {

}

// Yimin code

func TestRegisterLoginInfo(t *testing.T) {
	e := newExpect(t)

	rand.Seed(time.Now().UnixNano())
	registerValue := fmt.Sprintf("douyin%d", rand.Intn(65535))

	registerResp := e.POST("/douyin/user/register/").
		WithQuery("username", registerValue).WithQuery("password", registerValue).
		WithFormField("username", registerValue).WithFormField("password", registerValue).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	registerResp.Value("status_code").Number().IsEqual(0)
	registerResp.Value("user_id").Number().Gt(0)
	registerResp.Value("token").String().Length().Gt(0)

	loginResp := e.POST("/douyin/user/login/").
		WithQuery("username", registerValue).WithQuery("password", registerValue).
		WithFormField("username", registerValue).WithFormField("password", registerValue).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	loginResp.Value("status_code").Number().IsEqual(0)
	loginResp.Value("user_id").Number().Gt(0)
	loginResp.Value("token").String().Length().Gt(0)

	infoResp := e.GET("/douyin/user/").
		WithQuery("user_id", int64(loginResp.Value("user_id").Number().Raw())).WithQuery("token", loginResp.Value("token").String()).
		Expect().
		Status(http.StatusOK).
		JSON().Object()
	infoResp.Value("status_code").Number().IsEqual(0)
	infoResp.Value("status_msg").String().IsEqual("")
	infoResp.Value("user").Object().Value("name").String().IsEqual(registerValue)
}
