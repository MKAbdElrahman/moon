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
		// moon.WithErrorHandler(
		// 	func(err error, ctx moon.Context) {
		// 	ctx.SendString("error happended", http.StatusInternalServerError)
		// }),
	)

	// m.Use(moon.MiddlewareA, moon.MiddlewareB, moon.MiddlewareC)

	if err != nil {
		log.Fatal(err)
	}

	m.GET("/", func(ctx moon.Context) error {
		return ctx.SendString("Home Page", 200)
	})

	m.GET("/error", func(ctx moon.Context) error {
		return fmt.Errorf("error")
	})

	m.GET("/users/:id/:name", func(ctx moon.Context) error {
		return ctx.SendString(ctx.PathParam("id")+"  "+ctx.PathParam("name")+"   "+ctx.FirstQueryParam("q"), http.StatusOK)
	})

	http.ListenAndServe(":3000", m)

}

func MiddlewareB(next moon.HandlerFunc) moon.HandlerFunc {
	return func(c moon.Context) error {
		fmt.Println("Middleware B: Checking authorization")
		return next(c)
	}
}
