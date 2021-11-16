package application

import (
	"os"

	"github.com/chars-mc/encryptor-api/internal/encryption/domain"
)

var (
	aesSecretKey = []byte(os.Getenv("AES_SECRET_KEY"))
)

type DataService struct {
	repository domain.DataRepository
}

func NewDataService(repository domain.DataRepository) *DataService {
	return &DataService{repository}
}

func (s *DataService) Encrypt(data DataRequest, user UserDetails) (*DataResponse, error) {
	response := &DataResponse{
		ID:      1,
		Content: "encrypted_content",
	}
	return response, nil
}
