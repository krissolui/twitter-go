package main

import (
	"auth-service/storage"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	httputils "github.com/krissolui/go-utils/http-utils"
)

func (app *Config) invalidPath(w http.ResponseWriter, req *http.Request) {
	app.errorResponse(w, errors.New("invalid path"), EInvalidPath, http.StatusNotFound)
}

func (app *Config) verifyToken(w http.ResponseWriter, req *http.Request) {
	// TODO::read and verify token
	app.writeResponse(w, "verify token")
}

func (app *Config) verifyEmailUsernameNotUsed(userEntry storage.CreateUserEntry) (bool, error) {
	user, err := app.Storage.GetUserByEmail(userEntry.Email)
	if err != nil {
		return false, err
	}

	if user != nil && user.UserID != userEntry.UserID {
		return false, nil
	}

	user, err = app.Storage.GetUserByUsername(userEntry.Username)
	if err != nil {
		return false, err
	}

	if user != nil && user.UserID != userEntry.UserID {
		return false, nil
	}

	return true, nil
}

func (app *Config) signup(w http.ResponseWriter, req *http.Request) {
	userEntry, err := httputils.ReadJSON[storage.CreateUserEntry](req)
	if err != nil {
		app.errorResponse(w, err, EInvalidParams)
		return
	}

	emailUsernameValid, err := app.verifyEmailUsernameNotUsed(*userEntry)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}
	if !emailUsernameValid {
		app.errorResponse(w, errors.New("email and username must be unique"), EInvalidParams)
		return
	}

	userID, err := app.Storage.CreateUser(*userEntry)
	if err != nil {
		app.errorResponse(w, err, EInvalidParams)
		return
	}

	user, err := app.Storage.GetUserById(userID)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}

	app.writeResponse(w, "user created", *user)
}

func (app *Config) login(w http.ResponseWriter, req *http.Request) {
	loginEntry, err := httputils.ReadJSON[storage.LoginEntry](req)
	if err != nil {
		app.errorResponse(w, err, EInvalidParams)
		return
	}

	user, err := app.Storage.GetUserByEmail(loginEntry.Email)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}

	if user == nil {
		app.errorResponse(w, errors.New("invalid credentials"), EInvalidCredentials)
		return
	}

	verified := app.Storage.VerifyUserPassword(user.UserID, loginEntry.Password)
	if !verified {
		app.errorResponse(w, errors.New("invalid credentials"), EInvalidCredentials)
		return
	}

	// TODO::create session and send to session-service and user

	app.writeResponse(w, "logged in", user)
}

func (app *Config) logout(w http.ResponseWriter, req *http.Request) {
	// TODO::move to session-service

	userID := chi.URLParam(req, "userID")
	app.writeResponse(w, "logged out", userID)
}

func (app *Config) updateUser(w http.ResponseWriter, req *http.Request) {
	userID := chi.URLParam(req, "userID")
	user, err := httputils.ReadJSON[storage.User](req)
	if err != nil {
		app.errorResponse(w, err, EInvalidParams)
		return
	}
	user.UserID = userID

	userRecord, err := app.Storage.GetUserById(userID)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}

	if userRecord == nil {
		app.errorResponse(w, errors.New("user not found"), EInvalidParams)
		return
	}

	if user.Email == "" {
		user.Email = userRecord.Email
	}

	if user.Username == "" {
		user.Username = userRecord.Username
	}

	_, err = app.Storage.UpdateUser(*user)
	if err != nil {
		app.errorResponse(w, err, EInvalidParams)
		return
	}

	user, err = app.Storage.GetUserById(userID)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}

	app.writeResponse(w, "user updated", user)
}

func (app *Config) updatePassword(w http.ResponseWriter, req *http.Request) {
	userID := chi.URLParam(req, "userID")
	userPassword, err := httputils.ReadJSON[storage.UpdatePasswordEntry](req)
	if err != nil {
		app.errorResponse(w, err, EInvalidParams)
		return
	}
	userPassword.UserID = userID

	userRecord, err := app.Storage.GetUserById(userID)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}

	if userRecord == nil {
		app.errorResponse(w, errors.New("user not found"), EInvalidParams)
		return
	}

	if verified := app.Storage.VerifyUserPassword(userID, userPassword.OldPassword); !verified {
		app.errorResponse(w, errors.New("invalid credentials"), EInvalidCredentials)
		return
	}

	_, err = app.Storage.UpdateUserPassword(*userPassword)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}

	app.writeResponse(w, "user password updated", userID)
}

func (app *Config) deleteUser(w http.ResponseWriter, req *http.Request) {
	userID := chi.URLParam(req, "userID")

	userPassword, err := httputils.ReadJSON[storage.UserPassword](req)
	if err != nil {
		app.errorResponse(w, err, EInvalidParams)
		return
	}

	if verified := app.Storage.VerifyUserPassword(userID, userPassword.Password); !verified {
		app.errorResponse(w, errors.New("invalid credentials"), EInvalidCredentials)
		return
	}

	err = app.Storage.DeleteUser(userID)
	if err != nil {
		app.errorResponse(w, err, EBadRequest)
		return
	}

	app.writeResponse(w, "user deleted")
}
