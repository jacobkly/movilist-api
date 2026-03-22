package middleware

import (
	"context"
	"net/http"
	"time"
)

type Stats struct {
	DurationMs int64 `json:"duration_ms"`
}

type startKey struct{}

func StatsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), startKey{}, time.Now())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func StatsFromContext(ctx context.Context) *Stats {
	start, ok := ctx.Value(startKey{}).(time.Time)
	if !ok {
		return nil
	}

	return &Stats{
		DurationMs: time.Since(start).Milliseconds(),
	}
}
