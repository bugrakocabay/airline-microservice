package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bugrakocabay/airline/auth-service/domain/ports"
	"github.com/bugrakocabay/airline/auth-service/domain/token"
)

type AuthHandler struct {
	DataRepository ports.IAuthRepository
	tokenMaker     token.Maker
}

func NewAuthHandler(authRepo ports.IAuthRepository, tokenMaker token.Maker) *AuthHandler {
	return &AuthHandler{
		DataRepository: authRepo,
		tokenMaker:     tokenMaker,
	}
}

type newUserResponse struct {
	UserId    int    `json:"user_id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (app *AuthHandler) createUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email     string `json:"email" validate:"required"`
		Password  string `json:"password" validate:"required"`
		Firstname string `json:"firstname" validate:"required"`
		Lastname  string `json:"lastname" validate:"required"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.validate(requestPayload)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	userId, err := app.DataRepository.Insert(requestPayload.Email, requestPayload.Password, requestPayload.Firstname, requestPayload.Lastname)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created user!: %s", requestPayload.Email),
		Data: newUserResponse{
			UserId:    userId,
			Email:     requestPayload.Email,
			Firstname: requestPayload.Firstname,
			Lastname:  requestPayload.Lastname,
		},
	}

	_ = app.writeJSON(w, http.StatusCreated, payload)
}

type authenticateResponse struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Token     string `json:"token"`
}

func (app *AuthHandler) authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.validate(requestPayload)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.DataRepository.GetByEmail(requestPayload.Email)
	if err != nil {
		_ = app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := app.DataRepository.PasswordMatches(requestPayload.Password, user.Password)
	if err != nil || !valid {
		_ = app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	userId := strconv.Itoa(user.ID)
	authToken, err := app.tokenMaker.CreateToken(userId, user.Email, time.Hour)
	if err != nil || !valid {
		_ = app.errorJSON(w, errors.New("token error"), http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in as user: %s", user.Email),
		Data: authenticateResponse{
			Email:     user.Email,
			Firstname: user.FirstName,
			Lastname:  user.LastName,
			Token:     authToken,
		},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}
