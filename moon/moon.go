package moon

import (
	"net/http"
	"regexp"
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

	ctx := Context{
		Response:   w,
		Request:    r,
		PathParams: pathParams, // Attach the path parameters to the context

	}

	// attach the pathParams to the context
	m.applyMiddleware(h)(ctx)
}
func (m *Moon) findRouteHandler(path string, method string) HandlerFunc {
	for _, route := range m.routes {
		if route.method == method && IsRouteMatch(route.path, path) {
			return route.handlerFunc
		}
	}
	return m.notFound
}
func (m *Moon) parsePathParams(path string, method string) map[string]string {

	for _, route := range m.routes {
		if route.method == method && IsRouteMatch(route.path, path) {
			pathParams := ExtractParams(route.path, path)
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

func ExtractParamsFromRoute(routePath string) []string {
	// Define a regular expression pattern to match parameters in a route path
	pattern := `\{(\w+)\}`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Use FindAllStringSubmatch to find all matches of the pattern in the route path
	matches := re.FindAllStringSubmatch(routePath, -1)

	// Extract the captured groups (parameter names) from the matches
	var params []string
	for _, match := range matches {
		params = append(params, match[1])
	}

	return params
}
