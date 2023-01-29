package services

import (
	"fmt"
	"gorm.io/gorm"
)

type UserService struct {
	Order *OrderService `inject:"-"`
	DB    *gorm.DB      `inject:"-"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) GetUserInfo(uid int) {
	user := &UserNodel{}
	us.DB.Raw("select id,user_id from t_lark_task_1 where id=?", uid).First(user)
	fmt.Println(user)
}

type UserNodel struct {
	ID     int    `gorm:"column:id"`
	UserID string `gorm:"column:user_id"`
}
