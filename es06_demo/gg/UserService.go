package gg

import (
	"Go_ElasticSearch7/es06_demo/AppInit"
)

type UserService struct {
}

//构造函数
func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) GetUserById(uid int) (*UserModel, error) {
	user := NewUserModel()
	db := AppInit.GetDB().Table("users").Where("user_id=?", uid).First(&user)
	return user, db.Error
}
