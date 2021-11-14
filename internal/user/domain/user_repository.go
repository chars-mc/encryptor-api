package domain

type UserRepository interface {
	FetchByUsername(username string) (*User, error)
	Save(newUser User) (int64, error)
}
