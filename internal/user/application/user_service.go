package application

import (
	"errors"
	"strconv"
	"time"

	"github.com/chars-mc/encryptor-api/internal/user/domain"
	"github.com/golang-jwt/jwt/v4"
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

	claims := &Claims{
		Role: u.Role.String(),
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.Itoa(u.ID),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		// TODO: handle the return error
		return nil, err
	}

	response := &UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Role:     u.Role.String(),
		Token:    tokenString,
	}

	return response, nil
}
