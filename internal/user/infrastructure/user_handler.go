package infrastructure

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/chars-mc/encryptor-api/internal/api/server"
	"github.com/chars-mc/encryptor-api/internal/user/application"
)

type UserHandler struct {
	service application.UserUsecases
}

func NewUserHandler(service application.UserUsecases) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := application.UserLoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()

	if err != nil {
		server.WriteErrorJSON(w, http.StatusBadRequest, errors.New("Cannot process the data"))
		return
	}

	if err := user.Verify(); err != nil {
		server.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.service.Login(user)
	if err != nil {
		if err == sql.ErrNoRows {
			server.WriteErrorJSON(w, http.StatusNotFound, errors.New("Incorrect username or password"))
			return
		}
		server.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}
	server.WriteJSON(w, http.StatusOK, response)
	return
}

func (h *UserHandler) SignUpHandler(w http.ResponseWriter, r *http.Request) {
	userSignUpRequest := application.UserSignUpRequest{}
	err := json.NewDecoder(r.Body).Decode(&userSignUpRequest)
	defer r.Body.Close()

	if err != nil {
		server.WriteErrorJSON(w, http.StatusBadRequest, errors.New("Cannot process the data"))
		return
	}

	err = userSignUpRequest.Verify()
	if err != nil {
		server.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.service.SignUp(userSignUpRequest)
	if err != nil {
		if err == sql.ErrNoRows {
			server.WriteErrorJSON(w, http.StatusInternalServerError, errors.New("Internal server error, cannot create the user"))
			return
		}
		server.WriteErrorJSON(w, http.StatusBadRequest, err)
		return
	}
	server.WriteJSON(w, http.StatusOK, response)
}
