package user

import (
	"intracs_anpr_api/model"
	"intracs_anpr_api/repository"
)

type repositories interface {
	GetExistingUser(username string) (model.User, error)
	CreateUser(user model.UserRaw, password string, recoveryKey string) (model.UserRaw, error)
	GetUserByEmailAndPass(email string, password string) (model.User, error)
}

type Controller struct {
	service repositories
}

func InitController(userRepo *repository.UserRepo) *Controller {
	return &Controller{
		service: userRepo,
	}
}
