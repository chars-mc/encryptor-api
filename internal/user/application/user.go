package application

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var (
	secretKey                      = []byte(os.Getenv("SECRET_KEY"))
	errUsernameOrPasswordEmpty     = errors.New("The username and password cannot be empty")
	errUsernameLenght              = errors.New("The username length must be between 10 and 30 characters")
	errPasswordLenght              = errors.New("The password length must be between 10 and 30 characters")
	errUserRoleDoesNotExists       = errors.New("The user role doesn't exists")
	errUserAlreadyExists           = errors.New("The user already exists")
	errCannotGenereateHashPassword = errors.New("Cannot generate hash password")
)

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserLoginRequest) Verify() error {
	if u.Username == "" || u.Password == "" {
		return errUsernameOrPasswordEmpty
	}
	return nil
}

type UserSignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
}

func (u *UserSignUpRequest) Verify() error {
	if u.Username == "" || u.Password == "" {
		return errUsernameOrPasswordEmpty
	}
	if len(u.Username) < 10 || len(u.Username) > 30 {
		return errUsernameLenght
	}
	if len(u.Password) < 10 || len(u.Password) > 30 {
		return errPasswordLenght
	}
	if u.Role <= 0 || u.Role > 2 {
		return errUserRoleDoesNotExists
	}
	return nil
}

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}
