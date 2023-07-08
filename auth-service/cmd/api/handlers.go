package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type userResponse struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Token     string `json:"token"`
}

func (app *Config) authenticate(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Models.User.GetByEmail(requestPayload.Email)
	if err != nil {
		_ = app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	valid, err := user.PasswordMatches(requestPayload.Password)
	if err != nil || !valid {
		_ = app.errorJSON(w, errors.New("invalid credentials"), http.StatusBadRequest)
		return
	}

	userId := strconv.Itoa(user.ID)
	token, err := app.tokenMaker.CreateToken(userId, user.Email, time.Hour)
	if err != nil || !valid {
		_ = app.errorJSON(w, errors.New("token error"), http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Logged in as user: %s", user.Email),
		Data: userResponse{
			Email:     user.Email,
			Firstname: user.FirstName,
			Lastname:  user.LastName,
			Token:     token,
		},
	}

	_ = app.writeJSON(w, http.StatusOK, payload)
}

type CreateUserRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
}

type NewUserResponse struct {
	UserId    int    `json:"user_id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (app *Config) createUser(w http.ResponseWriter, r *http.Request) {
	var requestPayload CreateUserRequest

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	userId, err := app.Models.User.Insert(requestPayload.Email, requestPayload.Password, requestPayload.Firstname, requestPayload.Lastname)
	if err != nil {
		_ = app.errorJSON(w, err, http.StatusInternalServerError)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: fmt.Sprintf("Created user!: %s", requestPayload.Email),
		Data: NewUserResponse{
			UserId:    userId,
			Email:     requestPayload.Email,
			Firstname: requestPayload.Firstname,
			Lastname:  requestPayload.Lastname,
		},
	}

	log.Println(payload)

	_ = app.writeJSON(w, http.StatusCreated, payload)
}
