package router

import "net/http"

type HttpHandler struct{}

func (h *HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, err := res.Write([]byte("This is a monopoly game!"))
	if err != nil {
		return
	}
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{}
}
