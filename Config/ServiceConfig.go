package Config

import (
	"IoC/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
func (sc *ServiceConfig) GormDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:55001)/asyncflow?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
