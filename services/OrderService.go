package services

import "fmt"

type OrderService struct {
	Version string
}

func NewOrderService(version string) *OrderService {
	return &OrderService{Version: version}
}

func (us *OrderService) GetOrderInfo(uid int) {
	fmt.Println("获取订单ID=", uid, "的详细信息")
}
