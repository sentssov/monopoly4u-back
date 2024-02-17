package router

import "net/http"

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignIn(wr http.ResponseWriter, req *http.Request) {
	_, err := wr.Write([]byte("This is sign-in endpoint!"))
	if err != nil {
		return
	}
}
