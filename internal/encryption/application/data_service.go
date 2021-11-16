package application

import "github.com/chars-mc/encryptor-api/internal/encryption/domain"

type DataService struct {
	repository domain.DataRepository
}

func NewDataService(repository domain.DataRepository) *DataService {
	return &DataService{repository}
}

func (s *DataService) Encrypt(data DataRequest) (*DataResponse, error) {
	response := &DataResponse{
		ID:      1,
		Content: "encrypted_content",
	}
	return response, nil
}
