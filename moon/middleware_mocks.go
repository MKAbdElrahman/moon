package moon

import (
	"fmt"
	"time"
)

// Middleware A: Simple logging middleware
func MiddlewareA(next HandlerFunc) HandlerFunc {
	return func(c Context) error {
		// Perform some action before the request is handled
		fmt.Println("Middleware A: Logging request")

		// Call the next handler in the chain
		err := next(c)

		// Perform some action after the request is handled

		return err
	}
}

// Middleware B: Simple authorization middleware
func MiddlewareB(next HandlerFunc) HandlerFunc {
	return func(c Context) error {
		// Perform some authorization logic
		fmt.Println("Middleware B: Checking authorization")

		// Call the next handler in the chain
		return next(c)
	}
}

// Middleware C: Simple timing middleware
func MiddlewareC(next HandlerFunc) HandlerFunc {
	return func(c Context) error {
		// Record the start time of the request
		startTime := time.Now()

		// Call the next handler in the chain
		err := next(c)

		// Calculate and log the request processing time
		elapsed := time.Since(startTime)
		fmt.Printf("Middleware C: Request processed in %s\n", elapsed)

		return err
	}
}
