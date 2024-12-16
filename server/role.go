package server

import (
	"github.com/JobNing/user-rpc/model"
	"github.com/JobNing/user-rpc/pb/role"
)

func rolePbToModel(in *role.RoleInfo) *model.Role {
	return &model.Role{
		Model: model.Model{
			ID: in.ID,
		},
		Name: in.Name,
	}
}

func roleModelToPb(in *model.Role) *role.RoleInfo {
	return &role.RoleInfo{
		ID:   in.ID,
		Name: in.Name,
	}
}

//使用结构体复制方式
//func CreateRole(in *role.RoleInfo) (*role.RoleInfo, error) {
//	//声明model层的结构体并返回结构体指针，使用new
//	info := new(model.Role)
//	//使用copier.Copy方法将函数的参数pb中的role.RoleInfo 值复制到 model.Role里面
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
//	//使用copier.Copy方法将 model.Role 值复制到里面in里面
//	err = copier.Copy(in, res)
//	if err != nil {
//		return nil, err
//	}
//
//	return in, nil
//}

func CreateRole(in *role.RoleInfo) (*role.RoleInfo, error) {
	mod := rolePbToModel(in)
	res, err := mod.Create()
	if err != nil {
		return nil, err
	}

	return roleModelToPb(res), nil
}

func UpdateRole(in *role.RoleInfo) (*role.RoleInfo, error) {
	res, err := rolePbToModel(in).Update(in.ID)
	return roleModelToPb(res), err
}

func GetRole(id int64) (*role.RoleInfo, error) {
	info := new(model.Role)
	res, err := info.Get(id)
	if err != nil {
		return nil, err
	}

	return roleModelToPb(res), err
}

func SearchRoleByName(phone string) ([]*role.RoleInfo, error) {
	info := new(model.Role)
	res, err := info.SearchByName(phone)
	if err != nil {
		return nil, err
	}

	var roles []*role.RoleInfo
	for _, val := range res {
		roles = append(roles, roleModelToPb(val))
	}

	return roles, nil
}
