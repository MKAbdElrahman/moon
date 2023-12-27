package moon

import "net/http"

type Option func(*Moon) error

func New(options ...Option) (*Moon, error) {
	moon := &Moon{
		routes:   make([]Route, 0),
		notFound: defaultNotFound,
	}

	for _, option := range options {
		err := option(moon)
		if err != nil {
			return nil, err
		}
	}
	return moon, nil
}

func (moon *Moon) Use(m ...MiddlewareFunc) {
	moon.middleware = append(moon.middleware, m...)
}

func WithNotFound(h HandlerFunc) Option {
	return func(m *Moon) error {
		m.notFound = h
		return nil
	}
}

func defaultNotFound(c Context) error {
	http.NotFound(c.Response, c.Request)
	return nil
}
