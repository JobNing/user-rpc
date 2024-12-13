package api

import (
	"google.golang.org/grpc"
	role1 "user-rpc/api/role"
	user1 "user-rpc/api/user"
	"user-rpc/pb/role"
	"user-rpc/pb/user"
)

func Register(s *grpc.Server) {
	user.RegisterUserServer(s, &user1.UserService{})
	role.RegisterRoleServer(s, &role1.RoleService{})
}
