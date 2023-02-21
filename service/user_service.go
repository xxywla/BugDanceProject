package service

import (
	"douyinapp/entity"
	"douyinapp/repository"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func UserInfo(userId int64) (entity.User, error) {
	return repository.NewUserDaoInstance().QueryUserInfoById(userId)
}

// Yimin code
// create user info and set authentication
func CreateUser(username string, password string) (*entity.UserVo, error) {
	node, err := NewWorker(1)
	if err != nil {
		fmt.Printf("Failed to generate id: %s", err)
		return nil, err
	}
	id := node.GetId()
	encryptedPassword, err := EncryptPassword(password)
	if err != nil || len(encryptedPassword) == 0 {
		return nil, err
	}
	user := &entity.User{Id: id, Name: username, Password: encryptedPassword}
	if err := repository.NewUserDaoInstance().CreateUser(user); err != nil {
		return nil, err
	} else {
		userVo := user2Vo(user)
		return userVo, nil
	}
}

// verify username and password
func VerifyAccount(username string, password string) error {
	user, err := repository.NewUserDaoInstance().GetUserInfoByName(username)
	if err != nil {
		fmt.Printf("No User Found: %s", err)
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Printf("Wrong password: %s", err)
		return err
	}
	return nil
}

func GetUserInfoById(id int64) (*entity.UserVo, error) {
	user, err := repository.NewUserDaoInstance().GetUserInfoById(id)
	userVo := user2Vo(user)
	return userVo, err
}

func GetUserInfoByName(username string) (*entity.UserVo, error) {
	user, err := repository.NewUserDaoInstance().GetUserInfoByName(username)
	userVo := user2Vo(user)
	return userVo, err
}

func EncryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Failed to encrypt token: %s", err)
		return "", err
	}
	return string(hash), nil
}

func user2Vo(user *entity.User) *entity.UserVo {
	userVo := &entity.UserVo{Id: user.Id, Name: user.Name, FollowCount: 0, FollowerCount: 0, IsFollow: false}
	return userVo
}
