package router

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
	_, err := wr.Write([]byte("SignIn endpoint!"))
	if err != nil {
		return
	}
}

func (h *Handler) SignUp(wr http.ResponseWriter, req *http.Request) {
	var dto SignUpRequest
	if err := json.NewDecoder(req.Body).Decode(&dto); err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		h.logger.Errorf("Incoming request (Bad Request): %s", err.Error())
		return
	}

	player, err := h.manager.CreatePlayer(dto.Email, dto.Nickname, dto.Password)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusBadRequest)
		h.logger.Errorf("Error of creating new player: %s", err.Error())
		return
	}

	res, err := json.Marshal(&SignUpResponse{
		ID:    player.ID.String(),
		Email: player.Email.String(),
	})

	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		h.logger.Errorf("Error of marshaling the sign-up response to json: %s", err.Error())
		return
	}

	_, err = wr.Write(res)
	if err != nil {
		h.logger.Errorf("Error of writing response: %s", err.Error())
		return
	}
}
