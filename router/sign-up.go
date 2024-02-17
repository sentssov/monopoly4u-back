package router

import "net/http"

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignUp(wr http.ResponseWriter, req *http.Request) {
	_, err := wr.Write([]byte("This is sign up endpoint!"))
	if err != nil {
		return
	}
}
