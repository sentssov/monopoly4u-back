package router

import (
	"net/http"
)

type SignUpRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func SignUp(wr http.ResponseWriter, req *http.Request) {
	_, err := wr.Write([]byte("Hello from sign-up endpoint!"))
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
	}
}
