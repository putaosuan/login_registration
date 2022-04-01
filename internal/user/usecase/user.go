package usecase

import (
	"context"
	"fmt"
	"my_sso/ecode"
	"my_sso/internal/user/domain/service"

	pb "my_sso/api/user"
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
	if err := req.Validate(); err != nil {
		if err := req.Validate(); err != nil {
			ecode.ErrValidateFail.Message = err.Error()
			return nil, ecode.ErrValidateFail
		}
	}
	token, user, err := s.userService.Login(ctx, req.Mobile, req.Password)
	if err != nil {
		return &pb.LoginReply{}, err
	}
	return &pb.LoginReply{
		Id:     int64(user.Id),
		Name:   user.Name,
		Email:  user.Email,
		Mobile: user.Mobile,
		Token:  fmt.Sprintf("Bearer %s", token),
	}, nil
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
func (s *UserUseCase) UserCode(ctx context.Context, req *pb.UserCodeRequest) (*pb.UserCodeReply, error) {
	if err := req.Validate(); err != nil {
		if err := req.Validate(); err != nil {
			ecode.ErrValidateFail.Message = err.Error()
			return nil, ecode.ErrValidateFail
		}
	}
	err := s.userService.SendCode(ctx, req.Mobile)
	if err != nil {
		return &pb.UserCodeReply{}, err
	}
	return &pb.UserCodeReply{}, nil
}
func (s *UserUseCase) UserGet(ctx context.Context, req *pb.UserGetRequest) (*pb.UserGetReply, error) {
	user, err := s.userService.Get(ctx, req.Id)
	if err != nil {
		return &pb.UserGetReply{}, err
	}
	return &pb.UserGetReply{
		Id:     int64(user.Id),
		Name:   user.Name,
		Email:  user.Email,
		Mobile: user.Mobile,
	}, nil
}
