package router

import "net/http"

type SignInRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignUpRequest struct {
	Email    string `json:"email" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
	//RepeatPassword string `json:"repeat_password" binding:"required"`
}

func (h *Handler) SignIn(wr http.ResponseWriter, req *http.Request) {
	_, err := wr.Write([]byte("SignIn endpoint!"))
	if err != nil {
		return
	}
}

func (h *Handler) SignUp(wr http.ResponseWriter, req *http.Request) {
	_, err := wr.Write([]byte("SignUp endpoint!"))
	if err != nil {
		return
	}
}
