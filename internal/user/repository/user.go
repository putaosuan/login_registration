package repository

import (
	"context"
	"github.com/go-kirito/pkg/zdb"
	"my_sso/internal/user/domain/entity"
	"my_sso/internal/user/repository/model"
	"time"
)

type IUserRepo interface {
	Get(ctx context.Context, mobile string) (*entity.Users, error)
	CreateUser(ctx context.Context, users *entity.Users) (*entity.Users, error)
	CreateTrace(ctx context.Context, trace *entity.Trace) (*entity.Trace, error)
	CreateDevice(ctx context.Context, device *entity.Device) (*entity.Device, error)
}
type userRepo struct {
}

//@wire
func NewUserRepo() IUserRepo {
	return &userRepo{}
}
func (u *userRepo) Get(ctx context.Context, mobile string) (*entity.Users, error) {
	user := &model.Users{}
	err := zdb.NewOrm(ctx).Where("mobile = ?", mobile).Find(user).Error()
	if err != nil {
		return nil, err
	}
	return u.toResp(user), nil
}
func (u *userRepo) CreateUser(ctx context.Context, users *entity.Users) (*entity.Users, error) {
	user := u.toModelUsers(users)
	user.Mtime = time.Now()
	err := zdb.NewOrm(ctx).Create(user).Error()
	if err != nil {
		return nil, err
	}
	return u.toResp(user), nil
}

func (u *userRepo) CreateTrace(ctx context.Context, trace *entity.Trace) (*entity.Trace, error) {
	t := u.toModelTrace(trace)
	err := zdb.NewOrm(ctx).Create(t).Error()
	if err != nil {
		return nil, err
	}
	return trace, nil
}

func (u *userRepo) CreateDevice(ctx context.Context, device *entity.Device) (*entity.Device, error) {
	d := u.toModelDevice(device)
	err := zdb.NewOrm(ctx).Create(d).Error()
	if err != nil {
		return nil, err
	}
	return device, nil
}

func (u *userRepo) toResp(users *model.Users) *entity.Users {
	return &entity.Users{
		Id:     users.Id,
		Name:   users.Name,
		Email:  users.Email,
		Mobile: users.Mobile,
		Passwd: users.Passwd,
		Salt:   users.Salt,
		Ext:    users.Ext,
		Status: users.Status,
		Ctime:  users.Ctime,
		Mtime:  users.Mtime,
	}
}
func (u *userRepo) toModelUsers(users *entity.Users) *model.Users {
	return &model.Users{
		Name:   users.Name,
		Email:  users.Email,
		Mobile: users.Mobile,
		Passwd: users.Passwd,
		Salt:   users.Salt,
		Ext:    users.Ext,
		Status: users.Status,
		Ctime:  users.Ctime,
	}
}
func (u *userRepo) toModelTrace(trace *entity.Trace) *model.Trace {
	return &model.Trace{
		Id:    trace.Id,
		Uid:   trace.Uid,
		Type:  int8(trace.Type),
		Ip:    trace.Ip,
		Ext:   trace.Ext,
		Ctime: trace.Ctime,
	}
}
func (u *userRepo) toModelDevice(device *entity.Device) *model.Device {
	return &model.Device{
		Id:     device.Id,
		Uid:    device.Uid,
		Client: device.Client,
		Model:  device.Model,
		Ip:     device.Ip,
		Ext:    device.Ext,
		Ctime:  device.Ctime,
	}
}
