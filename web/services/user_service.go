package services

import (
	"basic_framework/web/response"
)

type UserService struct {
}

// 上传设备信息
func (this UserService) GetUser(username, password string, age int) (resp response.UserResp) {

	resp.Username = username
	resp.Password = password
	resp.Age = age

	return
}
