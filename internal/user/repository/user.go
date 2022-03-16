package repository

import (
	"context"
	"login_registration/internal/user/domain/entity"
)

type IUserRepo interface {
	Get(ctx context.Context, mobile string) (*entity.Users, error)
}
type userRepo struct {
}

//@wire
func NewUserRepo() IUserRepo {
	return &userRepo{}
}
func (u *userRepo) Get(ctx context.Context, mobile string) (*entity.Users, error) {
	//err := zdb.NewOrm(ctx).Table("users").Find().Error()
	return nil, nil
}
