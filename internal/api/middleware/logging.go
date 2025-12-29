package middleware

import (
	"log"
	"net/http"
	"time"

	chimw "github.com/go-chi/chi/v5/middleware"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rw := &responseWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		reqID := chimw.GetReqID(r.Context())

		ip := r.RemoteAddr
		if realIP := r.Header.Get("X-Forwarded-For"); realIP != "" {
			ip = realIP
		}

		log.Printf(
			"[REQUEST] id=%s method=%s path=%s query=%s ip=%s ua=%q",
			reqID,
			r.Method,
			r.URL.Path,
			r.URL.RawQuery,
			ip,
			r.UserAgent(),
		)

		// DEV ONLY — DO NOT ENABLE IN PROD
		// log.Printf("[HEADERS] %v", r.Header)

		// DEV ONLY — body logging is dangerous
		// log.Printf("[REQUEST BODY] %s", body)

		next.ServeHTTP(rw, r)

		log.Printf(
			"[RESPONSE] id=%s status=%d duration_ms=%d",
			reqID,
			rw.status,
			time.Since(start).Milliseconds(),
		)
	})
}
