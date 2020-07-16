package Middleware

import (
	"log"
	"net/http"
)

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		log.Printf("[INFO] Incoming Request from : %s, resource : %s, method : %s", request.RemoteAddr, request.URL.Path, request.Method)
		next.ServeHTTP(response, request)
	})
}


