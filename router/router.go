package router

import (
	"github.com/go-chi/chi/v5"
	"monopoly-auth/router/middleware"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.LogRequest)
	r.Post("/api/sign-in", SignIn)
	r.Post("/api/sign-up", SignUp)
	return r
}
