package main

import (
	. "IoC/Injector"
	"IoC/services"
	"fmt"
)

func main() {
	BeanFactory.Set(services.NewOrderService("1.0"))
	userService := services.NewUserService()
	BeanFactory.Apply(userService)
	fmt.Println(userService.Order)
}
