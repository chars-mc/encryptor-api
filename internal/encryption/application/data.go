package application

import "errors"

var (
	errContentLength = errors.New("The content length must be between 10 and 22 characters")
)

type DataRequest struct {
	Content     string `json:"content"`
	IdDataType  int    `json:"id_data_type"`
	IdAlgorithm int    `json:"id_algorithm"`
}

func (d *DataRequest) Verify() error {
	if len(d.Content) < 10 || len(d.Content) > 22 {
		return errContentLength
	}
	return nil
}

type DataResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}

// UserDetails is used to register the user actions
type UserDetails struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}
