package main

import (
	"IoC/Config"
	. "IoC/Injector"
	"IoC/services"
	"fmt"
)

func main() {
	serviceConfig := Config.NewServiceConfig()
	BeanFactory.ExprMap = map[string]interface{}{
		"ServiceConfig": serviceConfig,
	}
	BeanFactory.Set(serviceConfig)

	userService := services.NewUserService()
	BeanFactory.Apply(userService)
	fmt.Println(userService.Order)
}
