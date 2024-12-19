package role

import (
	"context"
	"github.com/JobNing/user-rpc/internal/server"
	"github.com/JobNing/user-rpc/pb/role"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoleService struct {
	role.UnimplementedRoleServer
}

func (s *RoleService) CreateRole(ctx context.Context, in *role.CreateRoleRequest) (*role.CreateRoleResponse, error) {
	if in.GetIn().Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "rolename is required")
	}

	info, err := server.CreateRole(in.GetIn())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &role.CreateRoleResponse{
		Info: info,
	}, nil
}

func (s *RoleService) UpdateRole(ctx context.Context, in *role.UpdateRoleRequest) (*role.UpdateRoleResponse, error) {
	if in.GetIn().Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "rolename is required")
	}

	info, err := server.UpdateRole(in.GetIn())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &role.UpdateRoleResponse{
		Info: info,
	}, nil
}

func (s *RoleService) GetRole(ctx context.Context, in *role.GetRoleRequest) (*role.GetRoleResponse, error) {
	if in.ID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id is required")
	}

	info, err := server.GetRole(in.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &role.GetRoleResponse{
		Info: info,
	}, nil
}

func (s *RoleService) SearchRoleByName(ctx context.Context, in *role.SearchRoleByNameRequest) (*role.SearchRoleByNameResponse, error) {
	info, err := server.SearchRoleByName(in.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &role.SearchRoleByNameResponse{
		Infos: info,
	}, nil
}
