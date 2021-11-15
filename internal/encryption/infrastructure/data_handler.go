package infrastructure

import (
	"net/http"

	"github.com/chars-mc/encryptor-api/internal/encryption/application"
)

type DataHandler struct {
	service application.DataUseCases
}

func NewDataHandler(service application.DataUseCases) *DataHandler {
	return &DataHandler{service}
}

func (h *DataHandler) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("encrypt"))
}
