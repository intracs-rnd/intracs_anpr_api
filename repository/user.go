package repository

import (
	"fmt"
	"intracs_anpr_api/model"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) GetExistingUser(email string) (model.User, error) {
	var user model.User
	fields := []string{"uuid", "fullname", "email", "created_at", "updated_at"}

	result := repo.db.Select(fields).Where("email = ?", email).First(&user)

	if result.Error != nil {
		fmt.Println("Failed get existing user", result.Error)
		return user, result.Error
	}

	return user, nil
}

func (repo *UserRepo) GetUserByEmailAndPass(email string, password string) (model.User, error) {
	var user model.User
	fields := []string{"uuid", "fullname", "email", "created_at", "updated_at"}

	result := repo.db.Select(fields).Where("email = ? AND password = ?", email, password).First(&user)

	if result.Error != nil {
		fmt.Println("Failed get existing user", result.Error)
		return user, result.Error
	}

	return user, nil
}

func (repo *UserRepo) CreateUser(user model.UserRaw, password string, recoveryKey string) (model.UserRaw, error) {
	result := repo.db.Create(&user)

	if result.Error != nil {
		fmt.Println("Failed create new user", result.Error)
		return user, result.Error
	}

	return user, nil
}
