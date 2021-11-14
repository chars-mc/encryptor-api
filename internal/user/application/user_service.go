package application

import (
	"errors"
	"strconv"
	"time"

	"github.com/chars-mc/encryptor-api/internal/api/security"
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

	if !security.CheckPasswordHash(user.Password, u.Password) {
		return nil, errors.New("Wrong password")
	}

	token, err := GenerateToken(u)
	if err != nil {
		// TODO: handle the return error
		return nil, err
	}

	response := &UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Role:     u.Role.String(),
		Token:    token,
	}

	return response, nil
}

func (s *UserService) SignUp(newUser UserSignUpRequest) (*UserResponse, error) {
	// check if the user already exists
	userExists, _ := s.repository.FetchByUsername(newUser.Username)
	if userExists != nil {
		return nil, errUserAlreadyExists
	}

	passwordHash, err := security.HashPassword(newUser.Password)
	if err != nil {
		return nil, errCannotGenereateHashPassword
	}
	user := domain.User{
		Username: newUser.Username,
		Password: passwordHash,
		Role:     domain.Role(newUser.Role),
	}

	_, err = s.repository.Save(user)
	if err != nil {
		return nil, err
	}

	u, err := s.repository.FetchByUsername(newUser.Username)
	if err != nil {
		return nil, err
	}

	token, err := GenerateToken(u)
	if err != nil {
		// TODO: handle the return error
		return nil, err
	}
	response := &UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Role:     u.Role.String(),
		Token:    token,
	}
	return response, nil
}

func GenerateToken(u *domain.User) (string, error) {
	claims := &Claims{
		Role: u.Role.String(),
		StandardClaims: jwt.StandardClaims{
			Id:        strconv.Itoa(u.ID),
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}
