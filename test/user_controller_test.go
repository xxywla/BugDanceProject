package test

import (
	"douyinapp/controller"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestUserInfo(t *testing.T) {
	res := controller.UserInfo(controller.UserRequest{UserId: 1, Token: "999"})
	output := res.StatusCode
	var expect int32 = 0
	assert.Equal(t, expect, output)
}
