package ecode

import "github.com/go-kirito/pkg/errors"

var (
	ErrUserPhoneRepeat  = errors.BadRequest("U000001", "用户已存在")
	ErrUserPhoneFailure = errors.BadRequest("U000002", "手机号格式不正确")
	ErrCodeFailure      = errors.BadRequest("U000003", "验证码错误")
	ErrUserLoginFailure = errors.BadRequest("U000004", "账号名或密码错误")
	ErrUserForbidden    = errors.BadRequest("U000005", "账号已禁用")
	ErrUserFailure      = errors.BadRequest("U000006", "用户不存在")
	//ErrCompanyAccountUpdateFailure = errors.BadRequest("C100002", "修改客户账号失败")
	//ErrCompanyAccountGetFailure    = errors.BadRequest("C100003", "获取客户账号失败")
	//ErrCompanyAccountDeleteFailure = errors.BadRequest("C100004", "删除客户账号失败")
	//ErrCompanyAccountPhoneRepeat   = errors.BadRequest("C100005", "账号的手机号已存在")
	//ErrCompanyAccountRepeat        = errors.BadRequest("C100006", "账号名已存在")
	//ErrCompanyAccountPhoneFailure  = errors.BadRequest("C100009", "手机号格式错误")

)
