package router

import (
	"encoding/json"
	"io"
	"monopoly-auth/tools"
	"net/http"
)

func SignIn(wr http.ResponseWriter, req *http.Request) {
	sReq := SignInRequest{}
	if err := json.NewDecoder(req.Body).Decode(&sReq); err != nil {
		Logger.Errorf("sign-in: json decode error: %s", err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			Logger.Errorf("sign-in: close body error: %s", err.Error())
			return
		}
	}(req.Body)

	for _, player := range Players {
		if player.Email.String() == sReq.Email {
			hashed, err := tools.HashPassword(sReq.Password)
			if err != nil {
				Logger.Errorf("sign-in: hash password error: %s", err.Error())
			}
			if tools.CompareHash(player.PasswordHash, hashed) {
				wr.WriteHeader(http.StatusUnauthorized)
				return
			} else {
				wr.WriteHeader(http.StatusOK)
			}
		}
	}
}
