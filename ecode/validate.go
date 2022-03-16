package ecode

import "github.com/go-kirito/pkg/errors"

var (
	ErrValidateFail = errors.BadRequest("v0000001", "参数校验失败")
)
