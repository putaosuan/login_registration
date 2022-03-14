// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/go-kirito/pkg/application"
	"login_registration/api/login"
	"login_registration/internal/login_registration/usecase"
)

// Injectors from wire.go:

func MakeUseCase() (*UseCases, error) {
	iLoginUseCase := usecase.NewLoginUseCase()
	useCases := &UseCases{
		iloginusecase0: iLoginUseCase,
	}
	return useCases, nil
}

// wire.go:

type UseCases struct {
	iloginusecase0 login.ILoginUseCase
}

func RegisterService(app application.Application) error {
	uc, err := MakeUseCase()
	if err != nil {
		return err
	}
	login.RegisterLoginServer(app, uc.iloginusecase0)
	return nil
}
