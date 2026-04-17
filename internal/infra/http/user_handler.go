package http

import (
	"encoding/json"
	"net/http"
	"go-crud/internal/usecase"
)

type UserHandler struct {
	usecase *usecase.UserUsecase
}

func NewUserHandler(u *usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: u}
}

type createUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	json.NewDecoder(r.Body).Decode(&req)

	err := h.usecase.Create(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	users, _ := h.usecase.List()
	json.NewEncoder(w).Encode(users)
}