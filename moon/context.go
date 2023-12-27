package moon

import (
	"fmt"
	"net/http"
)

const charsetUTF8 = "charset=UTF-8"

const HeaderContentType = "Content-Type"

const MIMETextPlain = "text/plain"

const MIMETextPlainCharsetUTF8 = MIMETextPlain + "; " + charsetUTF8

type Context struct {
	Response    http.ResponseWriter
	Request     *http.Request
	PathParams  map[string]string
	QueryParams map[string][]string
}

func (c *Context) SendBlob(code int, contentType string, b []byte) (err error) {
	c.writeContentType(contentType)
	c.Response.WriteHeader(code)
	_, err = c.Response.Write(b)
	return
}
func (c Context) SendString(s string, code int) error {
	return c.SendBlob(code, MIMETextPlainCharsetUTF8, []byte(s))
}

func (c *Context) writeContentType(value string) {
	header := c.Response.Header()
	if header.Get(HeaderContentType) == "" {
		header.Set(HeaderContentType, value)
	}
}

func (c *Context) PathParam(name string) string {
	return c.PathParams[name]
}

// Param returns query parameter by name.
func (c *Context) QueryParam(name string) []string {
	return c.QueryParams[name]
}

// Param returns query parameter by name.
func (c *Context) FirstQueryParam(name string) string {

	if len(c.QueryParam(name)) > 0 {
		return c.QueryParams[name][0]
	}

	fmt.Println(c.QueryParams[name], name)
	return ""
}
