package repository

import (
	"context"
	"login_registration/internal/user/domain/entity"
	"login_registration/internal/user/domain/valobj"
)

type IUserRepo interface {
	Get(ctx context.Context, mobile string) (*entity.Users, error)
	CreateUser(ctx context.Context, mobile string, password string) (*entity.Users, error)
	CreateTrace(ctx context.Context, ctime uint32, ip uint32, typ valobj.TraceType) (*entity.Trace, error)
	CreateDevice(ctx context.Context, ctime uint32, ip uint32, client string) (*entity.Device, error)
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
