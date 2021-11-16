package application

type DataRequest struct {
	Content     string `json:"content"`
	IdDataType  int    `json:"id_data_type"`
	IdAlgorithm int    `json:"id_algorithm"`
}

type DataResponse struct {
	ID      int    `json:"id"`
	Content []byte `json:"content"`
}

// UserDetails is used to register the user actions
type UserDetails struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}
