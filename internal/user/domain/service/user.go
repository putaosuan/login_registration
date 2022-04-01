package service

import (
	"context"
	"github.com/go-kirito/pkg/zconfig"
	"github.com/golang-jwt/jwt/v4"
	"my_sso/ecode"
	"my_sso/internal/user/domain/entity"
	"my_sso/internal/user/domain/valobj"
	"my_sso/internal/user/repository"
	"my_sso/pkg/util"
	"time"
)

type IUserService interface {
	Register(ctx context.Context, mobile string, password string, code string) (*entity.Users, error)
	SendCode(ctx context.Context, mobile string) error
	Login(ctx context.Context, mobile string, password string) (string, *entity.Users, error)
	Get(ctx context.Context, id int64) (*entity.Users, error)
}
type userService struct {
	userRepo      repository.IUserRepo
	userCacheRepo repository.IUserCacheRepo
}

//@wire
func NewUserService(a repository.IUserRepo, b repository.IUserCacheRepo) IUserService {
	return &userService{
		userRepo:      a,
		userCacheRepo: b,
	}
}
func (u *userService) Login(ctx context.Context, mobile string, password string) (string, *entity.Users, error) {
	//1.校验参数
	if !entity.VerifyMobileFormat(mobile) {
		return "", nil, ecode.ErrUserPhoneFailure
	}
	//2.判断用户是否存在
	user, err := u.userRepo.GetByMobile(ctx, mobile)
	if err != nil {
		return "", nil, err
	}
	if user.Id == 0 {
		return "", nil, ecode.ErrUserLoginFailure
	}
	//3.验证是否被封号
	if user.Status == 2 {
		return "", nil, ecode.ErrUserForbidden
	}
	//4.验证密码是否正确
	if !user.VerifyEncryptPassword(password, user.Passwd) {
		return "", nil, ecode.ErrUserLoginFailure
	}
	//5.生成jwt返回
	mp := jwt.MapClaims{}
	mp["uid"] = user.Id
	mp["name"] = user.Name
	mp["mobile"] = user.Mobile
	mp["email"] = user.Email
	mp["exp"] = time.Now().Add(util.RemianSecondWithToDay()).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, mp)
	key := zconfig.GetString("application.appKey")
	token, err := claims.SignedString([]byte(key))
	if err != nil {
		return "", nil, err
	}
	return token, user, nil
}
func (u *userService) Register(ctx context.Context, mobile string, password string, code string) (*entity.Users, error) {
	//1.校验参数
	if !entity.VerifyMobileFormat(mobile) {
		return nil, ecode.ErrUserPhoneFailure
	}
	//2.判断用户是否已经存在
	user, err := u.userRepo.GetByMobile(ctx, mobile)
	if err != nil {
		return nil, err
	}
	if user.Id != 0 {
		return nil, ecode.ErrUserPhoneRepeat
	}
	//3.验证code
	//cacheCode, err := u.userCacheRepo.GetCodeCache(ctx, mobile)
	//if err != nil {
	//	return nil, err
	//}
	//if cacheCode != "" {
	//	return nil, ecode.ErrUserCacheRepeat
	//}
	//if cacheCode != code {
	//	return nil, ecode.ErrCodeFailure
	//}
	//4.新增账户
	user = &entity.Users{
		Mobile: mobile,
		Passwd: password,
		Ctime:  uint32(int(time.Now().Unix())),
		Status: 1, //账户正常
		Mtime:  time.Now(),
	}
	user.Passwd, user.Salt = user.Encrypt()
	user2, err := u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}
	trace := &entity.Trace{
		Uid:   user2.Id,
		Type:  valobj.TraceTypeReg,
		Ip:    uint32(util.IpStringToInt(util.GetPeerAddr(ctx))),
		Ctime: user2.Ctime,
	}
	trace2, err := u.userRepo.CreateTrace(ctx, trace)
	if err != nil {
		return nil, err
	}
	device := &entity.Device{
		Uid:    user2.Id,
		Client: util.GetUserAgent(ctx),
		//Model:  "",
		Ip:    trace2.Ip,
		Ctime: user2.Ctime,
	}
	_, err = u.userRepo.CreateDevice(ctx, device)
	if err != nil {
		return nil, err
	}
	return user2, nil
}
func (u *userService) SendCode(ctx context.Context, mobile string) error {
	//1.校验参数
	if !entity.VerifyMobileFormat(mobile) {
		return ecode.ErrUserPhoneFailure
	}
	//2.判断手机号在规定3分钟内时间是否发送过
	code, err := u.userCacheRepo.GetCodeCache(ctx, mobile)
	if err != nil {
		return err
	}
	if code != "" {
		return ecode.ErrUserCacheRepeat
	}
	//3.生成验证信息，发送
	newCode := util.GetRandomNum(6)
	//发送
	_ = newCode
	//存入缓存
	err = u.userCacheRepo.CreateCodeCache(ctx, mobile, newCode)
	if err != nil {
		return err
	}
	return nil
}
func (u *userService) Get(ctx context.Context, id int64) (*entity.Users, error) {
	user, err := u.userRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if user.Id == 0 {
		return nil, ecode.ErrUserFailure
	}
	return user, nil
}
