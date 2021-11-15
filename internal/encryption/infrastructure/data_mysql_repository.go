package infrastructure

import (
	"github.com/chars-mc/encryptor-api/internal/database"
	"github.com/chars-mc/encryptor-api/internal/encryption/domain"
)

type DataMySQLRepository struct {
	db *database.MySQLClient
}

func NewDataMySQLRepository(db *database.MySQLClient) *DataMySQLRepository {
	return &DataMySQLRepository{db}
}

func (r *DataMySQLRepository) Save(data domain.Data) (int64, error) {
	return 0, nil
}
