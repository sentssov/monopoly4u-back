package router

import (
	"encoding/json"
	"io"
	"monopoly-auth/internal"
	"net/http"
)

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

	mng := internal.NewPlayerManager()

	player, err := mng.CreatePlayer(sReq.Email, sReq.Nickname, sReq.Password)
	if err != nil {
		Logger.Errorf("sign-up player creation error: %s", err.Error())
		http.Error(wr, err.Error(), http.StatusBadRequest)
		return
	}

	Players = append(Players, player)

	wr.WriteHeader(http.StatusCreated)
}
