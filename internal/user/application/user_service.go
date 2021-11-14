package application

import (
	"errors"

	"github.com/chars-mc/encryptor-api/internal/user/domain"
)

type UserService struct {
	repository domain.UserRepository
}

func NewUserService(repository domain.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) Login(user UserLoginRequest) (*UserResponse, error) {
	u, err := s.repository.FetchByUsername(user.Username)
	if err != nil {
		return nil, err
	}

	// TODO: use an encrypted password
	if user.Password != u.Password {
		return nil, errors.New("Wrong password")
	}

	// TODO: generate a valid token
	response := &UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Role:     u.Role.String(),
		Token:    "token",
	}

	return response, nil
}
