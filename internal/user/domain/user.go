package domain

import "time"

type Role int

const (
	Encryptor Role = iota + 1
	Decryptor
)

var roleString = []string{
	"Encryptor",
	"Decryptor",
}

func (r Role) String() string {
	return roleString[r-1]
}

type User struct {
	ID         int
	Username   string
	Password   string
	Created_at time.Time
	Role       Role
}
