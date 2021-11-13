package domain

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
