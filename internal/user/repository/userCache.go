package repository

import (
	"context"
	"my_sso/pkg/zcache"
	"time"
)

type IUserCacheRepo interface {
	CreateCodeCache(ctx context.Context, mobile string, code string) error
	GetCodeCache(ctx context.Context, mobile string) (string, error)
}
type userCacheRepo struct {
	cache zcache.ICache
}

//@wire
func NewUserCacheRepo() IUserCacheRepo {
	c := zcache.NewCache()
	return &userCacheRepo{
		cache: c,
	}
}
func (u *userCacheRepo) CreateCodeCache(ctx context.Context, mobile string, code string) error {
	if err := u.cache.Set(ctx, mobile, code, 5*time.Minute); err != nil {
		return err
	}
	return nil
}
func (u *userCacheRepo) GetCodeCache(ctx context.Context, mobile string) (string, error) {
	var err error
	code, err := u.cache.Get(ctx, mobile)
	if err != nil {
		return "", err
	}
	return code, nil
}
