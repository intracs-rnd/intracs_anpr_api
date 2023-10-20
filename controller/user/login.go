package user

import (
	"errors"
	"fmt"
	"intracs_anpr_api/internal/auth"
	"intracs_anpr_api/model"
)

func (c *Controller) Login(email string, password string) (model.AuthToken, error) {
	authToken := model.AuthToken{}
	user := model.User{}

	user, err := c.service.GetExistingUser(email)
	if err != nil {
		fmt.Println("Could not get existing user.", err)
		return authToken, err
	}
	if user.Email == "" {
		return authToken, errors.New("user not found")
	}

	passwordHash, err := auth.HashPassword(password)
	if err != nil {
		fmt.Println("Failed hash password", err)
		return authToken, errors.New("failed hash password")
	}

	user, err = c.service.GetUserByEmailAndPass(email, passwordHash)
	if err != nil {
		fmt.Println("Could not get login user.", err)
		return authToken, err
	}
	if user.Email == "" {
		return authToken, errors.New("email or password invalid")
	}

	token := auth.GenerateJWT(user)
	return token, nil
}
