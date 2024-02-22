package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"monopoly-auth/internal"
	"monopoly-auth/router/middleware"
)

var Logger logrus.Logger
var Players []*internal.Player

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.LogRequest)
	r.Post("/api/sign-in", SignIn)
	r.Post("/api/sign-up", SignUp)
	return r
}
