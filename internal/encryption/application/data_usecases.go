package application

type DataUseCases interface {
	Encrypt(data DataRequest, user UserDetails) (*DataResponse, error)
}
