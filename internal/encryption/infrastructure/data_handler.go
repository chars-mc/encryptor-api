package infrastructure

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/chars-mc/encryptor-api/internal/api/security"
	"github.com/chars-mc/encryptor-api/internal/api/server"
	"github.com/chars-mc/encryptor-api/internal/encryption/application"
	"github.com/golang-jwt/jwt/v4"
)

type DataHandler struct {
	service application.DataUseCases
}

func NewDataHandler(service application.DataUseCases) *DataHandler {
	return &DataHandler{service}
}

func (h *DataHandler) EncryptHandler(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	jwtToken, err := security.ParseToken(token)
	if err != nil {
		log.Printf("Cannot parse token: %v", err)
		server.WriteErrorJSON(w, http.StatusUnauthorized, errors.New("You must to be logged in"))
		return
	}

	userDetails := application.UserDetails{}
	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		userDetails.ID = claims["id"].(string)
		userDetails.Role = claims["id"].(string)
	} else {
		server.WriteErrorJSON(w, http.StatusUnauthorized, errors.New("You must to be logged in"))
		return
	}

	dataRequest := application.DataRequest{}
	err = json.NewDecoder(r.Body).Decode(&dataRequest)
	defer r.Body.Close()

	if err != nil {
		server.WriteErrorJSON(w, http.StatusBadRequest, errors.New("Cannot process the data"))
		return
	}

	if err := dataRequest.Verify(); err != nil {
		server.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.service.Encrypt(dataRequest, userDetails)
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
