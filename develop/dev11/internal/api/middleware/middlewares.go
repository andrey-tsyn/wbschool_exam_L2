package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

func WrapWithLoggerMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Debugf("Request recieved: %s", request.URL.Path)
		handler.ServeHTTP(writer, request)
		log.Debugf("Request handled: %s", request.URL.Path)
	})
}
