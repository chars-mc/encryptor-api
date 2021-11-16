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
	query := `
		INSERT INTO data(content, id_data_type, id_algorithm)
		VALUES(?, ?, ?);
	`
	result, err := r.db.Exec(query, data.Content, data.DataType, data.Algorithm)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
