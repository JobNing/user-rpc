package _type

import (
	"context"
	"github.com/JobNing/user-rpc/internal/server"
	"github.com/JobNing/user-rpc/pb/type"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TypeService struct {
	_type.UnimplementedTypeServer
}

func (s *TypeService) GetTypeListByLV1ID(ctx context.Context, in *_type.GetTypeListByLV1IDRequest) (*_type.GetTypeListByLV1IDResponse, error) {
	if in.GetID() == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "id is required")
	}

	info, err := server.GetByLV1ID(in.GetID())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &_type.GetTypeListByLV1IDResponse{
		Info: info,
	}, nil
}
