package handler

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

func (h *Handler) logging(handler http.Handler) http.Handler {
	logFn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		uri := r.RequestURI
		method := r.Method

		body, err := io.ReadAll(r.Body)
		if err != nil {
			h.logger.Err(err).Msg("failed to read request body")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.Body = io.NopCloser(bytes.NewReader(body))

		h.logger.Info().
			Str("uri", uri).
			Str("method", method).
			Any("headers", r.Header).
			Bytes("body", body).
			Msg("incoming request")

		responseData := &responseData{
			status: 0,
			size:   0,
		}
		lw := loggingResponseWriter{
			ResponseWriter: w,
			responseData:   responseData,
		}

		handler.ServeHTTP(&lw, r)

		duration := time.Since(start)

		body, err = io.ReadAll(r.Body)
		if err != nil {
			h.logger.Err(err).Msg("failed to read request body")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		r.Body = io.NopCloser(bytes.NewReader(body))

		h.logger.Info().
			Str("uri", uri).
			Str("method", method).
			Bytes("body", lw.responseData.body).
			Any("headers", lw.Header()).
			Int("status", lw.responseData.status).
			Dur("duration", duration).
			Int("size", lw.responseData.size).
			Send()
	}

	return http.HandlerFunc(logFn)
}

type responseData struct {
	status int
	size   int
	body   []byte
}

type loggingResponseWriter struct {
	http.ResponseWriter
	responseData *responseData
}

func (r *loggingResponseWriter) Write(b []byte) (int, error) {
	r.responseData.body = b
	size, err := r.ResponseWriter.Write(b)
	r.responseData.size += size
	return size, err
}

func (r *loggingResponseWriter) WriteHeader(statusCode int) {
	r.ResponseWriter.WriteHeader(statusCode)
	r.responseData.status = statusCode
}
