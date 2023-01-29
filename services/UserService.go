package services

import "fmt"

type UserService struct {
	Order IOrder `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetUserInfo(uid int) {
	fmt.Println("获取用户ID=", uid, "的详细信息")
}
