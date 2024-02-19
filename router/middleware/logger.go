package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

var Logger *logrus.Logger

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {
		Logger.Debugf("Incoming request to %s %s", req.Method, req.URL)
		next.ServeHTTP(wr, req)
	})
}
