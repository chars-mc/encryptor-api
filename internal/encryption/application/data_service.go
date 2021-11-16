package application

import (
	"encoding/hex"
	"errors"

	"github.com/chars-mc/encryptor-api/internal/api/security"
	"github.com/chars-mc/encryptor-api/internal/encryption/domain"
)

type DataService struct {
	repository domain.DataRepository
}

func NewDataService(repository domain.DataRepository) *DataService {
	return &DataService{repository}
}

func (s *DataService) Encrypt(data DataRequest, user UserDetails) (*DataResponse, error) {
	dataEncrypted := domain.Data{}

	switch data.IdAlgorithm {
	case domain.AES:
		c, err := security.AesEncrypt([]byte(data.Content))
		if err != nil {
			return nil, err
		}
		dataEncrypted.Content = hex.EncodeToString(c)
		dataEncrypted.Algorithm = domain.AES
	case domain.Blowsifh:
		byteText := []byte(data.Content)
		paddedplaintext := security.BlowfishChecksizeAndPad(byteText)
		c, err := security.BlowfishEncrypt(paddedplaintext)
		if err != nil {
			return nil, errors.New("Cannot encrypt data with the algorithm selected")
		}
		dataEncrypted.Algorithm = domain.Blowsifh
		dataEncrypted.Content = hex.EncodeToString(c)
	default:
		return nil, errors.New("Cannot encrypt data with the algorithm selected")
	}

	dataEncrypted.DataType = domain.DataType(data.IdDataType)
	dataId, err := s.repository.Save(dataEncrypted)
	if err != nil {
		return nil, err
	}

	response := &DataResponse{
		ID:      int(dataId),
		Content: dataEncrypted.Content,
	}
	return response, nil
}
