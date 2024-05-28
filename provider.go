//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	authController "learn-go/modules/auth/controllers"
	authService "learn-go/modules/auth/services"
	homeController "learn-go/modules/home/controllers"
	userModel "learn-go/modules/users/models"
	"learn-go/routes"
)

// InitializeController  initializes the Controller using Wire
func InitializeController() (*routes.Router, error) {
	wire.Build(
		userModel.NewUserModel,
		authController.NewAuthController,
		homeController.NewHomeController,
		authService.NewAuthService,
		routes.NewRouter,
	)

	return nil, nil
}
