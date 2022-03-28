package service

import (
	"context"
	"fmt"
	"login_registration/ecode"
	"login_registration/internal/user/domain/entity"
	"login_registration/internal/user/domain/valobj"
	"login_registration/internal/user/repository"
	"login_registration/pkg/util"
	"time"
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
	//1.校验参数
	if !entity.VerifyMobileFormat(mobile) {
		return nil, ecode.ErrUserPhoneFailure
	}
	//2.判断用户是否已经存在
	user, err := u.userRepo.Get(ctx, mobile)
	if err != nil {
		return nil, err
	}
	if user.Id != 0 {
		return nil, ecode.ErrUserPhoneRepeat
	}
	//3.验证code

	//4.新增账户
	user = &entity.Users{
		Mobile: mobile,
		Passwd: password,
		Ctime:  uint32(int(time.Now().Unix())),
		Status: 1, //账户正常
		Mtime:  time.Now(),
	}
	user.Passwd, user.Salt = user.Encrypt()
	//user2, err := u.userRepo.CreateUser(ctx, user)
	//if err != nil {
	//	return nil, err
	//}
	fmt.Println(user.Salt)
	fmt.Println(user.Passwd)
	trace := &entity.Trace{
		Id:   0,
		Uid:  0,
		Type: valobj.TraceTypeReg,
		Ip:   uint32(util.IpStringToInt(util.GetRealAddr(ctx))),
		Ext:  "",
		//Ctime: user2.Ctime,
	}
	fmt.Println(util.GetPeerAddr(ctx), "----")
	fmt.Println(trace.Ip)
	return nil, nil
}
