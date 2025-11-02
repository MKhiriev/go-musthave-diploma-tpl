package handler

import (
	"context"
	"net/http"
	"time"
)

func withRequestTimeout(timeout time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			cw := contextResponseWriter{
				ResponseWriter: w,
				responseData:   &responseData{},
			}
			next.ServeHTTP(&cw, r.WithContext(ctx))

			select {
			// if deadline exceeded - return 504
			case <-ctx.Done():
				cw.ResponseWriter.WriteHeader(http.StatusGatewayTimeout)
				_, _ = cw.ResponseWriter.Write([]byte(http.StatusText(http.StatusGatewayTimeout)))

			// if deadline not exceeded return handler result
			default:
				cw.ResponseWriter.WriteHeader(cw.responseData.status)
				_, _ = cw.ResponseWriter.Write(cw.responseData.body)
				return
			}
		})
	}
}

type contextResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

// Write does nothing to prevent writing to http-body more than once
func (r *contextResponseWriter) Write(b []byte) (int, error) {
	r.responseData.body = b
	return 0, nil
}

// WriteHeader does nothing to prevent `Http: superfluous response.WriteHeader call` error
func (r *contextResponseWriter) WriteHeader(statusCode int) {
	r.responseData.status = statusCode
}
