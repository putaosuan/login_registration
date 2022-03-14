package usecase

import (
	"context"
	"fmt"

	pb "login_registration/api/login"
)

type LoginUseCase struct {
}

// @wire
func NewLoginUseCase() pb.ILoginUseCase {
	return &LoginUseCase{}
}

func (s *LoginUseCase) CreateLogin(ctx context.Context, req *pb.CreateLoginRequest) (*pb.CreateLoginReply, error) {
	return &pb.CreateLoginReply{}, nil
}
func (s *LoginUseCase) UpdateLogin(ctx context.Context, req *pb.UpdateLoginRequest) (*pb.UpdateLoginReply, error) {
	return &pb.UpdateLoginReply{}, nil
}
func (s *LoginUseCase) DeleteLogin(ctx context.Context, req *pb.DeleteLoginRequest) (*pb.DeleteLoginReply, error) {
	return &pb.DeleteLoginReply{}, nil
}
func (s *LoginUseCase) GetLogin(ctx context.Context, req *pb.GetLoginRequest) (*pb.GetLoginReply, error) {
	return &pb.GetLoginReply{}, nil
}
func (s *LoginUseCase) ListLogin(ctx context.Context, req *pb.ListLoginRequest) (*pb.ListLoginReply, error) {
	fmt.Println("nihao")
	return &pb.ListLoginReply{}, nil
}
