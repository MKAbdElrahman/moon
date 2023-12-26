package router

import "net/http"

// defaultNotFoundHandler is the default not-found handler for the router.
var defaultNotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	http.NotFound(w, req)
})

// MiddlewareFunc represents a middleware function.
type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

// Route represents a specific route in the router.
type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

// Router is the interface that routers should implement.
type Router interface {
	AddRoute(route Route)
	FindRoute(path string) http.HandlerFunc
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

type GroupRouter struct {
	routes          []Route
	notFoundHandler http.HandlerFunc
	middleware      []MiddlewareFunc
}

// Option is a functional option type for configuring the router.
type Option func(*GroupRouter) error

// WithNotFoundHandler sets a custom not-found handler for the router.
func WithNotFoundHandler(handler http.HandlerFunc) Option {
	return func(r *GroupRouter) error {
		r.notFoundHandler = handler
		return nil
	}
}

// WithMiddleware adds middleware functions to the router.
func WithMiddleware(middleware ...MiddlewareFunc) Option {
	return func(r *GroupRouter) error {
		r.middleware = append(r.middleware, middleware...)
		return nil
	}
}

// Use adds middleware functions to the router.
func (r *GroupRouter) Use(middleware ...MiddlewareFunc) {
	r.middleware = append(r.middleware, middleware...)
}

// NewGroupRouter creates a new router with options.
func NewGroupRouter(options ...Option) (*GroupRouter, error) {
	router := &GroupRouter{
		routes:          make([]Route, 0),
		notFoundHandler: defaultNotFoundHandler,
	}

	for _, option := range options {
		if err := option(router); err != nil {
			return nil, err
		}
	}

	return router, nil
}

// AddRoute adds a route to the group router.
func (r *GroupRouter) AddRoute(route Route) {
	r.routes = append(r.routes, route)
}

// FindRoute finds a route in the group router for the given path and method.
func (r *GroupRouter) FindRoute(method, path string) http.HandlerFunc {
	for _, route := range r.routes {
		if route.Path == path && route.Method == method {
			return route.Handler
		}
	}
	return nil
}

// ServeHTTP implements the http.Handler interface for the group router.
func (r *GroupRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	// Find the route based on path and method
	handler := r.FindRoute(method, path)

	if handler != nil {
		// Apply router-level middleware
		handler = applyMiddleware(handler, r.middleware)

		// Execute the final handler
		handler(w, req)
	} else {
		r.notFoundHandler(w, req)
	}
}

// Define the applyMiddleware function
func applyMiddleware(handler http.HandlerFunc, middleware []MiddlewareFunc) http.HandlerFunc {
	// Apply router-level middleware in reverse order
	for i := len(middleware) - 1; i >= 0; i-- {
		handler = middleware[i](handler)
	}
	return handler
}
