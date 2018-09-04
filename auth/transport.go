package auth

import (
	"encoding/json"
	"net/http"
	"schoolsystem/auth-microservice/models"
	"time"

	"github.com/go-chi/chi"
)

type HttpUserHandler struct {
	AuthUsecase AuthUsecase
}

func (h *HttpUserHandler) Login(w http.ResponseWriter, r *http.Request) {
	loginData := models.Login{}
	json.NewDecoder(r.Body).Decode(&loginData)

	// validate the data
	err := loginData.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	// returns true if user can login, otherwise error with information
	canLogin, err := h.AuthUsecase.Login(loginData.Email, loginData.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(canLogin)
}

func (h *HttpUserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	json.NewDecoder(r.Body).Decode(&u)

	err := u.Validate()
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	userResp, err := h.AuthUsecase.Signup(&u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "User created successfully",
		Data:    &userResp,
	})
}

func NewAuthHttpHandler(r *chi.Mux, us AuthUsecase) {
	handler := HttpUserHandler{
		AuthUsecase: us,
	}

	r.Post("/p/auth/login", handler.Login)
	r.Post("/p/auth/signup", handler.Signup)
}
