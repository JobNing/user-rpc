package server

import (
	model2 "github.com/JobNing/user-rpc/internal/model"
	"github.com/JobNing/user-rpc/pb/user"
)

func userPbToModel(in *user.UserInfo) *model2.User {
	return &model2.User{
		Model: model2.Model{
			ID: in.ID,
		},
		Username: in.Username,
		Phone:    in.Phone,
		Age:      in.Age,
		Sex:      int64(in.Sex),
	}
}

func userModelToPb(in *model2.User) *user.UserInfo {
	return &user.UserInfo{
		ID:       in.ID,
		Username: in.Username,
		Phone:    in.Phone,
		Age:      in.Age,
		Sex:      user.SexType(in.Sex),
	}
}

//使用结构体复制方式
//func CreateUser(in *user.UserInfo) (*user.UserInfo, error) {
//	//声明model层的结构体并返回结构体指针，使用new
//	info := new(model.User)
//	//使用copier.Copy方法将函数的参数pb中的user.UserInfo 值复制到 model.User里面
//	err := copier.Copy(info, in)
//	if err != nil {
//		return nil, err
//	}
//
//	res, err := info.Create()
//	if err != nil {
//		return nil, err
//	}
//
//	//使用copier.Copy方法将 model.User 值复制到里面in里面
//	err = copier.Copy(in, res)
//	if err != nil {
//		return nil, err
//	}
//
//	return in, nil
//}

func CreateUser(in *user.UserInfo) (*user.UserInfo, error) {
	mod := userPbToModel(in)
	res, err := mod.Create()
	if err != nil {
		return nil, err
	}

	return userModelToPb(res), nil
}

func UpdateUser(in *user.UserInfo) (*user.UserInfo, error) {
	res, err := userPbToModel(in).Update(in.ID)
	return userModelToPb(res), err
}

func GetUser(id int64) (*user.UserInfo, error) {
	info := new(model2.User)
	res, err := info.Get(id)
	if err != nil {
		return nil, err
	}

	return userModelToPb(res), err
}

func GetUserByPhone(phone string) (*user.UserInfo, error) {
	info := new(model2.User)
	res, err := info.GetByPhone(phone)
	if err != nil {
		return nil, err
	}

	return userModelToPb(res), err
}
