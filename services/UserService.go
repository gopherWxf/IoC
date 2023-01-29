package services

import "fmt"

type UserService struct {
	order *OrderService
}

func NewUserService(order *OrderService) *UserService {
	return &UserService{order: order}
}

func (us *UserService) GetUserInfo(uid int) {
	fmt.Println("获取用户ID=", uid, "的详细信息")
}
func (us *UserService) GetOrderInfo(uid int) {
	us.order.GetOrderInfo(uid)
}
