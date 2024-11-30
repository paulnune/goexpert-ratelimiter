package middleware

import (
	"net"
	"net/http"

	"github.com/paulnune/goexpert-ratelimiter/internal/usecase"
)

// rateLimitMiddleware applies rate limiting to incoming requests
func RateLimitMiddleware(next http.Handler, limiter *usecase.RateLimiter) http.Handler {

	result := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ip := r.RemoteAddr
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		api_token := r.Header["Api_key"]
		s_token := ""
		if len(api_token) > 0 {
			s_token = api_token[0]
		}
		if !limiter.Allow(ip, s_token) {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})

	return result
}
