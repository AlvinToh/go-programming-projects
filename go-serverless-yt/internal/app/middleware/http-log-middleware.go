package middleware

import (
	"log"
	"net/http"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	length     int
	body       []byte
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	// Copy the response body
	body := make([]byte, len(b))
	copy(body, b)

	// Store the response body
	lrw.body = body

	size, err := lrw.ResponseWriter.Write(b)
	lrw.length += size
	return size, err
}

func HttpLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := &loggingResponseWriter{ResponseWriter: w}
		log.Println("Request:", r.Method, r.RequestURI)
		next.ServeHTTP(lrw, r)
		log.Println("Response:", lrw.statusCode, "Body:", string(lrw.body))
	})
}
