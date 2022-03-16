package service

import (
	"context"
	"login_registration/ecode"
	"login_registration/internal/user/domain/entity"
	"login_registration/internal/user/repository"
)

type IUserService interface {
	Register(ctx context.Context, mobile string, password string, code string) (*entity.Users, error)
}
type userService struct {
	userRepo repository.IUserRepo
}

//@wire
func NewUserService(a repository.IUserRepo) IUserService {
	return &userService{
		userRepo: a,
	}
}
func (u *userService) Register(ctx context.Context, mobile string, password string, code string) (*entity.Users, error) {
	//1.判断用户是否已经存在
	user, err := u.userRepo.Get(ctx, mobile)
	if err != nil {
		return nil, err
	}
	if user.Id != 0 {
		return nil, ecode.ErrUserPhoneRepeat
	}
	return nil, nil
}
