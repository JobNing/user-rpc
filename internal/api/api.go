package api

import (
	role1 "github.com/JobNing/user-rpc/internal/api/role"
	type1 "github.com/JobNing/user-rpc/internal/api/type"
	user1 "github.com/JobNing/user-rpc/internal/api/user"
	"github.com/JobNing/user-rpc/pb/role"
	_type "github.com/JobNing/user-rpc/pb/type"
	"github.com/JobNing/user-rpc/pb/user"
	"google.golang.org/grpc"
)

func Register(s *grpc.Server) {
	user.RegisterUserServer(s, &user1.UserService{})
	role.RegisterRoleServer(s, &role1.RoleService{})
	_type.RegisterTypeServer(s, &type1.TypeService{})
}
