package moon

import (
	"net/http"
)

type Route struct {
	path        string
	method      string
	handlerFunc HandlerFunc
}

func (m *Moon) AddRoute(path string, method string, handlerFunc HandlerFunc) {

	// Validate and sanitize path
	path = sanitizePath(path)

	m.routes = append(m.routes, Route{
		method:      method,
		path:        path,
		handlerFunc: handlerFunc,
	})
}

func sanitizePath(path string) string {
	if path == "" {
		return "/"
	}
	if path[0] != '/' {
		return "/" + path
	}
	return path
}

func (m *Moon) GET(path string, h HandlerFunc) {
	m.AddRoute(path, http.MethodGet, h)
}

func (m *Moon) POST(path string, h HandlerFunc) {
	m.AddRoute(path, http.MethodPost, h)
}

func (m *Moon) PUT(path string, h HandlerFunc) {
	m.AddRoute(path, http.MethodPut, h)
}

func (m *Moon) DELETE(path string, h HandlerFunc) {
	m.AddRoute(path, http.MethodDelete, h)
}

func (m *Moon) PATCH(path string, h HandlerFunc) {
	m.AddRoute(path, http.MethodPatch, h)
}

func (m *Moon) OPTIONS(path string, h HandlerFunc) {
	m.AddRoute(path, http.MethodOptions, h)
}

func (m *Moon) HEAD(path string, h HandlerFunc) {
	m.AddRoute(path, http.MethodHead, h)
}
