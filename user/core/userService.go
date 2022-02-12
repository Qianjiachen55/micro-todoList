package core

import (
	"context"
	"errors"
	"github.com/Qianjiachen55/micro-todoList/model"
	"github.com/Qianjiachen55/micro-todoList/services"
	"gorm.io/gorm"
)

func BuildUser(user model.User) *services.UserModel {
	userModel := services.UserModel{
		ID:        uint32(user.ID),
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.Unix(),
		UpdatedAt: user.UpdatedAt.Unix(),
	}

	return &userModel
}


func (*UserService) UserLogin(ctx context.Context,req *services.UserRequest,resp *services.UserDetailResponse )error {
	var user model.User

	resp.Code=200


	if err := model.DB.First(&user,"user_name=?",req.UserName).Error;err!=nil{
		if err==gorm.ErrRecordNotFound{
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	if user.CheckPassword(req.Password) ==false{
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}

func (*UserService) UserRegister (ctx context.Context,req *services.UserRequest,resp *services.UserDetailResponse) error{
	if req.PasswordConfirm!=req.Password{
		return errors.New("两次密码不一致")
	}

	var count int64 =0

	if err := model.DB.Model(&model.User{}).Where("user_name=?",req.UserName).Count(&count).Error;err != nil{
		return err
	}

	if count >0{
		err := errors.New("用户已经存在")
		return err
	}

	user := model.User{
		UserName:       req.UserName,
	}

	if err := user.SetPassword(req.Password);err != nil{
		return err
	}

	if err := model.DB.Create(&user).Error;err != nil{
		return  err
	}
	resp.UserDetail = BuildUser(user)
	return nil
}