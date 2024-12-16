package api

import (
	role1 "github.com/JobNing/user-rpc/api/role"
	user1 "github.com/JobNing/user-rpc/api/user"
	"github.com/JobNing/user-rpc/pb/role"
	"github.com/JobNing/user-rpc/pb/user"
	"google.golang.org/grpc"
)

func Register(s *grpc.Server) {
	user.RegisterUserServer(s, &user1.UserService{})
	role.RegisterRoleServer(s, &role1.RoleService{})
}
