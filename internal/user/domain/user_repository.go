package domain

type UserRepository interface {
	FetchByUsername(username string) (*User, error)
}
