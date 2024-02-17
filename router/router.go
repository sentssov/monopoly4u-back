package router

import "github.com/go-chi/chi/v5"

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/api/sign-in", SignIn)
	r.Post("/api/sign-up", SignUp)
	return r
}
