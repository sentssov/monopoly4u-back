package router

import (
	"encoding/json"
	"log"
	"monopoly-auth/internal"
	"net/http"
)

type SignInRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SignIn(wr http.ResponseWriter, req *http.Request) {
	var sReq SignInRequest
	if err := json.NewDecoder(req.Body).Decode(&sReq); err != nil {
		wr.WriteHeader(http.StatusInternalServerError)
		log.Println("The request body cannot be readed: ", err.Error())
	}

	players = append(players, internal.Player{
		Nickname: sReq.Username,
		Password: sReq.Password,
	})
	pJson, _ := json.Marshal(players)

	wr.Header().Add("Content-Type", "application/json")
	_, err := wr.Write(pJson)
	if err != nil {
		log.Fatal(err)
	}
}
