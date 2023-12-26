package router

import "net/http"

// Get adds a route for the GET method.
func (r *GroupRouter) Get(path string, handler http.HandlerFunc) {
	r.AddRoute(Route{
		Path:    path,
		Method:  http.MethodGet,
		Handler: handler,
	})
}

// Post adds a route for the POST method.
func (r *GroupRouter) Post(path string, handler http.HandlerFunc) {
	r.AddRoute(Route{
		Path:    path,
		Method:  http.MethodPost,
		Handler: handler,
	})
}

// Put adds a route for the PUT method.
func (r *GroupRouter) Put(path string, handler http.HandlerFunc) {
	r.AddRoute(Route{
		Path:    path,
		Method:  http.MethodPut,
		Handler: handler,
	})
}

// Delete adds a route for the DELETE method.
func (r *GroupRouter) Delete(path string, handler http.HandlerFunc) {
	r.AddRoute(Route{
		Path:    path,
		Method:  http.MethodDelete,
		Handler: handler,
	})
}

// Patch adds a route for the PATCH method.
func (r *GroupRouter) Patch(path string, handler http.HandlerFunc) {
	r.AddRoute(Route{
		Path:    path,
		Method:  http.MethodPatch,
		Handler: handler,
	})
}

// Options adds a route for the OPTIONS method.
func (r *GroupRouter) Options(path string, handler http.HandlerFunc) {
	r.AddRoute(Route{
		Path:    path,
		Method:  http.MethodOptions,
		Handler: handler,
	})
}

// Head adds a route for the HEAD method.
func (r *GroupRouter) Head(path string, handler http.HandlerFunc) {
	r.AddRoute(Route{
		Path:    path,
		Method:  http.MethodHead,
		Handler: handler,
	})
}
