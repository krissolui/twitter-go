package user

import (
	"auth-service/internal/config"
	customerror "auth-service/internal/customError"
	"auth-service/internal/entity"
	"auth-service/internal/storage"
	"fmt"
	"net/http"

	httputils "github.com/krissolui/go-utils/http-utils"
)

type UserService struct {
	storage           *storage.Storage
	sessionServiceURL string
}

type Service interface {
	CreateUser(entity.CreateUserEntry) (*entity.User, *customerror.CustomError)
	UpdateUser(string, entity.User) (*entity.User, *customerror.CustomError)
	DeleteUser(string, entity.UserPassword) *customerror.CustomError
	Login(entity.LoginEntry) (*entity.User, *customerror.CustomError)
	UpdatePassword(string, entity.UpdatePasswordEntry) *customerror.CustomError
	GenerateSession(string) (*entity.Session, *customerror.CustomError)
}

func NewUserService(cfg *config.Config, storage *storage.Storage) Service {
	return &UserService{
		storage:           storage,
		sessionServiceURL: cfg.SessionServiceURL,
	}
}

func (s *UserService) VerifyEmailUsernameNotUsed(userEntry entity.User) (bool, *customerror.CustomError) {
	user, err := s.storage.GetUserByEmail(userEntry.Email)
	if err != nil {
		return false, customerror.InternalError("failed to get user by email")
	}

	if user != nil && user.UserID != userEntry.UserID {
		return false, nil
	}

	user, err = s.storage.GetUserByUsername(userEntry.Username)
	if err != nil {
		return false, customerror.InternalError("failed to get user by username")
	}

	if user != nil && user.UserID != userEntry.UserID {
		return false, nil
	}

	return true, nil
}

func (s *UserService) CreateUser(userEntry entity.CreateUserEntry) (*entity.User, *customerror.CustomError) {
	emailUsernameValid, customError := s.VerifyEmailUsernameNotUsed(entity.User{
		Username: userEntry.Username,
		Email:    userEntry.Email,
	})
	if customError != nil {
		return nil, customError
	}
	if !emailUsernameValid {
		return nil, customerror.InvalidParams("email and username must be unique")
	}

	userID, err := s.storage.CreateUser(userEntry)
	if err != nil {
		return nil, customerror.InternalError("failed to create user")
	}

	user, err := s.storage.GetUserById(userID)
	if err != nil {
		return nil, customerror.InternalError("failed to get new user")
	}

	return user, nil
}

func (s *UserService) UpdateUser(userID string, userEntry entity.User) (*entity.User, *customerror.CustomError) {
	userRecord, customError := s.GetUser(userID)
	if customError != nil {
		return nil, customError
	}

	if userEntry.Email == "" {
		userEntry.Email = userRecord.Email
	}

	if userEntry.Username == "" {
		userEntry.Username = userRecord.Username
	}

	_, err := s.storage.UpdateUser(userEntry)
	if err != nil {
		return nil, customerror.InternalError("failed to update user")
	}

	user, customError := s.GetUser(userID)
	if customError != nil {
		return nil, customError
	}

	return user, nil
}

func (s *UserService) DeleteUser(userID string, userPassword entity.UserPassword) *customerror.CustomError {
	if verified := s.storage.VerifyUserPassword(userID, userPassword.Password); !verified {
		return customerror.InvalidCredentials("invalid credentials")
	}

	err := s.storage.DeleteUser(userID)
	if err != nil {
		return customerror.InternalError("failed to delete user")
	}

	return nil
}

func (s *UserService) Login(loginEntry entity.LoginEntry) (*entity.User, *customerror.CustomError) {
	user, err := s.storage.GetUserByEmail(loginEntry.Email)
	if err != nil {
		return nil, customerror.InternalError("failed to get user by email")
	}

	if user == nil {
		return nil, customerror.InvalidCredentials("invalid credentials")
	}

	verified := s.storage.VerifyUserPassword(user.UserID, loginEntry.Password)
	if !verified {
		return nil, customerror.InvalidCredentials("invalid credentials")
	}

	return user, nil
}

func (s *UserService) UpdatePassword(userID string, userPassword entity.UpdatePasswordEntry) *customerror.CustomError {
	_, customError := s.GetUser(userID)
	if customError != nil {
		return customError
	}

	if verified := s.storage.VerifyUserPassword(userID, userPassword.OldPassword); !verified {
		return customerror.InvalidCredentials("invalid credentials")
	}

	_, err := s.storage.UpdateUserPassword(userPassword)
	if err != nil {
		return customerror.InternalError("failed to update user password")
	}

	return nil
}

func (s *UserService) GetUser(userID string) (*entity.User, *customerror.CustomError) {
	user, err := s.storage.GetUserById(userID)
	if err != nil {
		return nil, customerror.InternalError("failed to get user")
	}

	if user == nil {
		return nil, customerror.InvalidParams("user not found")
	}

	return user, nil
}

func (s *UserService) GenerateSession(userID string) (*entity.Session, *customerror.CustomError) {
	path := fmt.Sprintf("/%s", userID)
	query := make(map[string]string, 0)

	res, err := httputils.SendPlainRequest(s.sessionServiceURL, path, http.MethodPost, query)
	if err != nil {
		return nil, customerror.InternalError(err.Error())
	}

	session, err := httputils.ReadJSONRes[entity.Session](res)
	if err != nil {
		return nil, customerror.InternalError(err.Error())
	}

	return session, nil
}
