package main

import (
	. "IoC/Injector"
	"IoC/services"
	"fmt"
)

func main() {
	BeanFactory.Set(services.NewOrderService("1.0"))
	order := BeanFactory.Get((*services.OrderService)(nil))
	fmt.Println(order)
}
