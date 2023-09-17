package routes

import (
	customerror "auth-service/internal/customError"
	"auth-service/internal/entity"
	"auth-service/internal/util"
	"net/http"

	"github.com/go-chi/chi/v5"
	httputils "github.com/krissolui/go-utils/http-utils"
)

func (router *Router) signup(w http.ResponseWriter, r *http.Request) {
	userEntry, err := httputils.ReadJSONReq[entity.CreateUserEntry](r)
	if err != nil {
		util.ErrorResponse(w, customerror.InvalidParams(err.Error()))
		return
	}

	user, customError := router.service.CreateUser(*userEntry)
	if customError != nil {
		util.ErrorResponse(w, customError)
		return
	}

	session, customError := router.service.GenerateSession(user.UserID)
	if customError != nil {
		util.ErrorResponse(w, customError)
		return
	}

	util.WriteResponse(w, "user created", session)
}

func (router *Router) login(w http.ResponseWriter, r *http.Request) {
	loginEntry, err := httputils.ReadJSONReq[entity.LoginEntry](r)
	if err != nil {
		util.ErrorResponse(w, customerror.InvalidParams(err.Error()))
		return
	}

	user, customError := router.service.Login(*loginEntry)
	if customError != nil {
		util.ErrorResponse(w, customError)
		return
	}

	session, customError := router.service.GenerateSession(user.UserID)
	if customError != nil {
		util.ErrorResponse(w, customError)
		return
	}

	util.WriteResponse(w, "logged in", session)
}

func (router *Router) updateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	userEntry, err := httputils.ReadJSONReq[entity.User](r)
	if err != nil {
		util.ErrorResponse(w, customerror.InvalidParams(err.Error()))
		return
	}
	userEntry.UserID = userID

	user, customError := router.service.UpdateUser(userID, *userEntry)
	if customError != nil {
		util.ErrorResponse(w, customError)
		return
	}

	util.WriteResponse(w, "user updated", user)
}

func (router *Router) updatePassword(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	userPassword, err := httputils.ReadJSONReq[entity.UpdatePasswordEntry](r)
	if err != nil {
		util.ErrorResponse(w, customerror.InvalidParams(err.Error()))
		return
	}
	userPassword.UserID = userID

	customError := router.service.UpdatePassword(userID, *userPassword)
	if customError != nil {
		util.ErrorResponse(w, customError)
		return
	}

	util.WriteResponse(w, "user password updated", userID)
}

func (router *Router) deleteUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")

	userPassword, err := httputils.ReadJSONReq[entity.UserPassword](r)
	if err != nil {
		util.ErrorResponse(w, customerror.InvalidParams(err.Error()))
		return
	}

	customError := router.service.DeleteUser(userID, *userPassword)
	if customError != nil {
		util.ErrorResponse(w, customError)
		return
	}

	util.WriteResponse(w, "user deleted")
}
