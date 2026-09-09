package v1

import (
	"context"

	pb "{{ .App.ModuleName }}/pkg/api/proto/v1"
)

type UserService struct{}

func (u *UserService) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{Response: "test-username"}, nil
}
