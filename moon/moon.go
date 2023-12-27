package moon

import (
	"net/http"
)

type HandlerFunc func(Context) error
type MiddlewareFunc func(next HandlerFunc) HandlerFunc

type Moon struct {
	routes     []Route
	notFound   HandlerFunc
	middleware []MiddlewareFunc
}

func (m *Moon) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	method := r.Method

	h := m.findRouteHandler(path, method)
	pathParams := m.parsePathParams(path, method)

	queryParams := extractQueryParamsFromRawQuery(r.URL.RawQuery)

	ctx := Context{
		Response:    w,
		Request:     r,
		PathParams:  pathParams, // Attach the path parameters to the context
		QueryParams: queryParams,
	}

	// attach the pathParams to the context
	m.applyMiddleware(h)(ctx)
}
func (m *Moon) findRouteHandler(path string, method string) HandlerFunc {
	for _, route := range m.routes {
		if route.method == method && isRouteMatch(route.path, path) {
			return route.handlerFunc
		}
	}
	return m.notFound
}
func (m *Moon) parsePathParams(path string, method string) map[string]string {

	for _, route := range m.routes {
		if route.method == method && isRouteMatch(route.path, path) {
			pathParams := extractPathParams(route.path, path)
			return pathParams
		}
	}
	return map[string]string{}
}

func (moon *Moon) applyMiddleware(h HandlerFunc) HandlerFunc {
	// Apply middleware in reverse order
	for i := len(moon.middleware) - 1; i >= 0; i-- {
		h = moon.middleware[i](h)
	}
	return h
}
