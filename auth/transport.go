package auth

import (
	"church-adoration/models"
	"encoding/json"
	"net/http"
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
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	// returns true if user can login, otherwise error with information
	user, token, err := h.AuthUsecase.Login(loginData.Email, loginData.Password)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "Logged in successfully",
		Data: map[string]interface{}{
			"token": token,
			"user":  user,
		},
	})
}

func (h *HttpUserHandler) Signup(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)
	u.CreatedAt = time.Now()
	u.Role = "NORMAL_USER"

	err := u.Validate()
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	userResp, _, err := h.AuthUsecase.Signup(&u)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: err,
		})
		return
	}

	json.NewEncoder(w).Encode(models.Response{
		Error:   false,
		Message: "User created successfully",
		Data:    &userResp,
	})
}

func (h *HttpUserHandler) ActivateUser(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "id")

	if userID == "" {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: "This is not a valid link",
			Data: nil,
		})
		return
	}

	err := h.AuthUsecase.ActivateUser(userID)
	if err != nil {
		json.NewEncoder(w).Encode(models.Response{
			Error:   true,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	w.Write([]byte("User activated successfully"))
}

func NewAuthHttpHandler(r *chi.Mux, us AuthUsecase) {
	handler := HttpUserHandler{
		AuthUsecase: us,
	}

	r.Post("/p/auth/login", handler.Login)
	r.Post("/p/auth/signup", handler.Signup)
	r.Get("/p/auth/activate/{id}", handler.ActivateUser)

}
