package controller

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/messaging-go-service/internal/model"
	"github.com/messaging-go-service/internal/repository"
	httputil "github.com/messaging-go-service/pkg/http"
)

type UserController interface {
	SearchUsers(w http.ResponseWriter, r *http.Request)
	GetUserDetail(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

type UserControllerImpl struct {
	UserRepository repository.UserRepository
}

func NewUserController(userRepository repository.UserRepository) UserController {
	return &UserControllerImpl{
		UserRepository: userRepository,
	}
}

func (c *UserControllerImpl) SearchUsers(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	users, err := c.UserRepository.GetAllUsers(context.Background())
	if err != nil {
		httputil.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	var filteredUsers []model.User
	for _, user := range users {
		if name == "" || contains(user.Username, name) {
			filteredUsers = append(filteredUsers, user)
		}
	}

	if len(filteredUsers) == 0 {
		httputil.WriteResponse(w, http.StatusNotFound, map[string]string{"error": "No user found"})
		return
	}

	response := struct {
		Message string       `json:"message"`
		Data    []model.User `json:"data"`
	}{
		Message: "Search results have been retrieved successfully",
		Data:    filteredUsers,
	}
	httputil.WriteResponse(w, http.StatusOK, response)
}

func (c *UserControllerImpl) GetUserDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDstr := vars["id"]

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		httputil.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid user id"})
		return
	}

	user, err := c.UserRepository.GetUserByID(context.Background(), userID)
	if err != nil {
		httputil.WriteResponse(w, http.StatusNotFound, map[string]string{"error": "User not found"})
		return
	}

	response := struct {
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}{
		Message: "User detail has been retrieved",
		Data:    *user,
	}

	httputil.WriteResponse(w, http.StatusOK, response)
}

func (c *UserControllerImpl) CreateUser(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := httputil.ReadRequest(r, &requestBody); err != nil {
		httputil.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	newUser := model.User{
		Username: requestBody.Username,
		Email:    requestBody.Email,
		Password: requestBody.Password,
	}

	if err := c.UserRepository.CreateUser(context.Background(), &newUser); err != nil {
		httputil.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error creating user"})
		return
	}

	response := struct {
		Message string     `json:"message"`
		Data    model.User `json:"data"`
	}{
		Message: "User has been created",
		Data:    newUser,
	}

	httputil.WriteResponse(w, http.StatusOK, response)
}

func (c *UserControllerImpl) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDstr := vars["id"]

	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		httputil.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid user id"})
		return
	}

	var user model.User

	var requestBody struct {
		Username       string `json:"username"`
		Desc           string `json:"desc"`
		Displayname    string `json:"displayname"`
		ProfilePicture string `json:"profile_picture"`
	}

	if err := httputil.ReadRequest(r, &requestBody); err != nil {
		httputil.WriteResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	updatedUser := model.User{
		Username:       requestBody.Username,
		Desc:           requestBody.Desc,
		ProfilePicture: requestBody.ProfilePicture,
	}

	user.ID = userID
	if err := c.UserRepository.UpdateUser(context.Background(), userID, &updatedUser); err != nil {
		httputil.WriteResponse(w, http.StatusInternalServerError, map[string]string{"error": "Error updating user"})
		return
	}

	response := struct {
		Message string `json:"message"`
	}{
		Message: "User has been updated",
	}

	httputil.WriteResponse(w, http.StatusOK, response)
}

func contains(s, substr string) bool {
	return len(s) > 0 && len(substr) > 0 && strings.Contains(s, substr)
}
