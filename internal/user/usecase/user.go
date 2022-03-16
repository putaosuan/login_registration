package usecase

import (
	"context"
	"login_registration/ecode"
	"login_registration/internal/user/domain/service"

	pb "login_registration/api/user"
)

type UserUseCase struct {
	userService service.IUserService
}

// @wire
func NewUserUseCase(a service.IUserService) pb.IUserUseCase {
	return &UserUseCase{
		userService: a,
	}
}

func (s *UserUseCase) UserLogin(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {

	return &pb.LoginReply{}, nil
}
func (s *UserUseCase) UserRegister(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	if err := req.Validate(); err != nil {
		if err := req.Validate(); err != nil {
			ecode.ErrValidateFail.Message = err.Error()
			return nil, ecode.ErrValidateFail
		}
	}
	_, err := s.userService.Register(ctx, req.Mobile, req.Password, req.Code)
	if err != nil {
		return &pb.RegisterReply{}, err
	}
	return &pb.RegisterReply{}, nil
}
