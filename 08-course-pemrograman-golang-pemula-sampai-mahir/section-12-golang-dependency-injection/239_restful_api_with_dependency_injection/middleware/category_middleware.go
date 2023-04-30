package middleware

import "net/http"

type CategoryMiddleware struct {
	Handler http.Handler
}

func (c *CategoryMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c.Handler.ServeHTTP(writer, request)
}

func NewCategoryMiddleware(handler http.Handler) *CategoryMiddleware {
	return &CategoryMiddleware{
		Handler: handler,
	}
}
