package main

import (
	"fmt"
	"log"
	"moon/moon"
	"net/http"
)

func main() {

	m, err := moon.New(
		moon.WithNotFound(
			func(ctx moon.Context) error {
				return ctx.SendString("route not registered", http.StatusNotFound)
			},
		),
	)

	// m.Use(moon.MiddlewareA, moon.MiddlewareB, moon.MiddlewareC)

	if err != nil {
		log.Fatal(err)
	}

	m.GET("/", func(ctx moon.Context) error {
		return ctx.SendString("Home Page", 200)
	})

	m.GET("/users/:id/:name", func(ctx moon.Context) error {
		return ctx.SendString(ctx.PathParam("id")+ "  "+ ctx.PathParam("name"), http.StatusOK)
	})

	http.ListenAndServe(":3000", m)

}

func MiddlewareB(next moon.HandlerFunc) moon.HandlerFunc {
	return func(c moon.Context) error {
		fmt.Println("Middleware B: Checking authorization")
		return next(c)
	}
}
