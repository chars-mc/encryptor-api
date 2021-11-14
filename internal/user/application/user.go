package application

import "errors"

var (
	errUsernameOrPasswordEmpty = errors.New("The username and password cannot be empty")
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

type UserResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Token    string `json:"token"`
}
