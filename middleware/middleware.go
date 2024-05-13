package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func CreateStack(xs ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for _, middleware := range xs {
			next = middleware(next)
		}
		return next
	}
}
