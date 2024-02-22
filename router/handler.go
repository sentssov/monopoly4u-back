package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"monopoly-auth/router/middleware"
)

type Handler struct {
	logger *logrus.Logger
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.LogRequest)
	r.Post("/api/auth/sign-in", h.SignIn)
	r.Post("/api/auth/sign-up", h.SignUp)

	return r
}

func NewHandler(logger *logrus.Logger) *Handler {
	return &Handler{logger}
}
