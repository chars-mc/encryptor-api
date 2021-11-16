package infrastructure

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/chars-mc/encryptor-api/internal/api/server"
	"github.com/chars-mc/encryptor-api/internal/encryption/application"
)

type DataHandler struct {
	service application.DataUseCases
}

func NewDataHandler(service application.DataUseCases) *DataHandler {
	return &DataHandler{service}
}

func (h *DataHandler) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	dataRequest := application.DataRequest{}
	err := json.NewDecoder(r.Body).Decode(&dataRequest)
	if err != nil {
		server.WriteErrorJSON(w, http.StatusInternalServerError, errors.New("Internal server error"))
		return
	}

	response, err := h.service.Encrypt(dataRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			server.WriteErrorJSON(w, http.StatusNotFound, errors.New("Cannot encrypt file"))
			return
		}
		server.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}
	server.WriteJSON(w, http.StatusOK, response)
}
