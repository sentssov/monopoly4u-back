package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sirupsen/logrus"
	"monopoly-auth/internal/services"
)

type Handler struct {
	playerManager *services.PlayerManager
	signInManager services.Authentication
	logger        *logrus.Logger
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

func NewHandler(
	playerManager *services.PlayerManager,
	signInManager services.Authentication,
	logger *logrus.Logger) *Handler {
	return &Handler{
		playerManager: playerManager,
		signInManager: signInManager,
		logger:        logger,
	}
}
