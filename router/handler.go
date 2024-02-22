package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	logger *logrus.Logger
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/api/auth/sign-in", h.SignIn)
	r.Post("/api/auth/sign-up", h.SignUp)

	return r
}

func NewHandler(logger *logrus.Logger) *Handler {
	return &Handler{logger}
}
