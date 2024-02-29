package rest

import (
	"encoding/json"
	"net/http"
)

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (h *Handler) SignIn(wr http.ResponseWriter, req *http.Request) {
	var payload SignInRequest
	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: h.signInManager.SignIn(payload.Email, payload.Password)

	http.SetCookie(wr, &http.Cookie{
		Name: "token",
	})

	// TODO: реализация api/auth/sign-in
}

func (h *Handler) SignUp(wr http.ResponseWriter, req *http.Request) {
	var payload SignUpRequest
	if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: реализация api/auth/sign-up
}
