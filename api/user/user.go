package user

import (
	"context"
	"github.com/JobNing/user-rpc/pb/user"
	"github.com/JobNing/user-rpc/server"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	user.UnimplementedUserServer
}

func (s *UserService) CreateUser(ctx context.Context, in *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	if in.GetIn().Username == "" {
		return nil, status.Errorf(codes.InvalidArgument, "username is required")
	}
	if len(in.GetIn().Phone) != 11 {
		return nil, status.Errorf(codes.InvalidArgument, "phone is invalid")
	}

	info, err := server.CreateUser(in.GetIn())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &user.CreateUserResponse{
		Info: info,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, in *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	if len(in.GetIn().Phone) != 11 && in.GetIn().Phone != "" {
		return nil, status.Errorf(codes.InvalidArgument, "phone is invalid")
	}

	info, err := server.UpdateUser(in.GetIn())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &user.UpdateUserResponse{
		Info: info,
	}, nil
}
func (s *UserService) GetUser(ctx context.Context, in *user.GetUserRequest) (*user.GetUserResponse, error) {
	if in.ID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id is required")
	}

	info, err := server.GetUser(in.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &user.GetUserResponse{
		Info: info,
	}, nil
}
func (s *UserService) GetUserByPhone(ctx context.Context, in *user.GetUserByPhoneRequest) (*user.GetUserByPhoneResponse, error) {
	if len(in.Phone) != 11 {
		return nil, status.Errorf(codes.InvalidArgument, "phone is invalid")
	}

	info, err := server.GetUserByPhone(in.Phone)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &user.GetUserByPhoneResponse{
		Info: info,
	}, nil
}
