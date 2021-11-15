package domain

type DataRepository interface {
	Save(data Data) (int64, error)
}
