package main

import (
	"IoC/Config"
	. "IoC/Injector"
	"IoC/services"
)

func main() {
	serviceConfig := Config.NewServiceConfig()
	BeanFactory.Config(serviceConfig)
	//BeanFactory.Set(serviceConfig)

	userService := services.NewUserService()
	BeanFactory.Apply(userService)
	userService.GetUserInfo(98489)
}
