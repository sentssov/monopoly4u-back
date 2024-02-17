package router

import (
	"net/http"
)

type SignInRequest struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func SignIn(wr http.ResponseWriter, req *http.Request) {
	_, err := wr.Write([]byte("Hello from sign-in endpoint!"))
	if err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
	}
}
