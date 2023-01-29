package Config

import (
	"IoC/services"
	"log"
)

type ServiceConfig struct {
}

func NewServiceConfig() *ServiceConfig {
	return &ServiceConfig{}
}

func (sc *ServiceConfig) OrderService() *services.OrderService {
	log.Printf("初始化OrderService\n")
	return services.NewOrderService("2.0")
}
