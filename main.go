package main

import (
	"chi/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux, err := router.NewGroupRouter(router.WithNotFoundHandler(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Sorry! Not Found")
	}))

	mux.Use(router.LoggingMiddleware)

	mux.Use(MiddlewareA, MiddlewareB, MiddlewareC)

	if err != nil {
		panic(err)
	}

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home")
	})

	http.ListenAndServe(":8080", mux)
}

// MiddlewareA logs a message.
func MiddlewareA(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("MiddlewareA: Executed")

		// Call the next handler in the chain
		next(w, req)
	}
}

// MiddlewareB logs a message.
func MiddlewareB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("MiddlewareB: Executed")

		// Call the next handler in the chain
		next(w, req)
	}
}

// MiddlewareC logs a message.
func MiddlewareC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("MiddlewareC: Executed")

		// Call the next handler in the chain
		next(w, req)
	}
}
