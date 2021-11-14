package application

type UserUsecases interface {
	Login(user UserLoginRequest) (*UserResponse, error)
}
