package application

type DataRequest struct {
	Content     string `json:"content"`
	IdDataType  int    `json:"id_data_type"`
	IdAlgorithm int    `json:"id_algorithm"`
}

type DataResponse struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
}
