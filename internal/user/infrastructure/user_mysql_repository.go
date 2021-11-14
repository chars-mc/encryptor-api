package infrastructure

import (
	"github.com/chars-mc/encryptor-api/internal/database"
	"github.com/chars-mc/encryptor-api/internal/user/domain"
)

type UserMySQLRepository struct {
	db *database.MySQLClient
}

func NewUserMySQLRepository(db *database.MySQLClient) *UserMySQLRepository {
	return &UserMySQLRepository{db}
}

func (r *UserMySQLRepository) FetchByUsername(username string) (*domain.User, error) {
	user := &domain.User{}
	query := `
		SELECT * FROM user WHERE username = ?;
	`
	err := r.db.QueryRow(query, username).
		Scan(&user.ID, &user.Username, &user.Password, &user.Created_at, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserMySQLRepository) Save(newUser domain.User) (int64, error) {
	query := `
		INSERT INTO user(username, password, id_user_role)
		VALUES(?, ?, ?);
	`
	result, err := r.db.Exec(query, newUser.Username, newUser.Password, newUser.Role)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
