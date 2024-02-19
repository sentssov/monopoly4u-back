package router

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"io"
	"monopoly-auth/internal"
	"net/http"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

func SignUp(wr http.ResponseWriter, req *http.Request) {
	sReq := SignUpRequest{}
	if err := json.NewDecoder(req.Body).Decode(&sReq); err != nil {
		Logger.Errorf("sign-up json decoding error: %s", err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Logger.Errorf("sign-up json closed error: %s", err.Error())
			return
		}
	}(req.Body)
	hash, err := bcrypt.GenerateFromPassword([]byte(sReq.Password), bcrypt.MinCost)
	if err != nil {
		Logger.Errorf("sign-up hash generation error: %s", err.Error())
		return
	}
	Players = append(Players, internal.Player{
		Email:        sReq.Email,
		Nickname:     sReq.Nickname,
		PasswordHash: hash,
	})

	wr.WriteHeader(http.StatusCreated)
}
