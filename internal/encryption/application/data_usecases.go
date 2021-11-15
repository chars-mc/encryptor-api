package application

type DataUseCases interface {
	Encrypt(data DataRequest) (*DataResponse, error)
}
