package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/paulnune/goexpert-ratelimiter/internal/database"
	"github.com/paulnune/goexpert-ratelimiter/internal/usecase"
	"github.com/paulnune/goexpert-ratelimiter/internal/web/handler"
	"github.com/paulnune/goexpert-ratelimiter/internal/web/middleware"
)

func main() {

	limit, err := strconv.Atoi(os.Getenv("RATELIMIT"))
	if err != nil {
		panic("RATE LIMIT not defined or invalid")
	}

	interval, err := strconv.Atoi(os.Getenv("RATELIMIT_CLEANUP_INTERVAL"))
	if err != nil {
		panic("RATE LIMIT INTERVAL not defined or invalid")
	}

	blockInterval, err := strconv.Atoi(os.Getenv("RATELIMIT_BLOCK_TIME"))
	if err != nil {
		panic("RATELIMIT BLOCK TIME not defined or invalid")
	}

	listTokens := database.NewTokenLimitList(os.Getenv("RATELIMIT_TOKEN_LIST"))

	redis_url := os.Getenv("RATELIMIT_REDIS_URL")
	if redis_url == "" {
		panic("RATELIMIT_REDIS_URL not defined or invalid")
	}

	ctx := context.Background()

	options := make(map[string]string)
	options["addr"] = redis_url
	options["password"] = ""
	options["db"] = "0"

	redisClient, err := database.NewRedisClient(ctx, options)
	if err != nil {
		panic("client error: " + err.Error())
	}

	limiter := usecase.NewIpRateLimiter(
		ctx,
		limit,
		time.Millisecond*time.Duration(interval),
		time.Millisecond*time.Duration(blockInterval),
		listTokens,
		&redisClient)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler.HelloWorldHandler)

	log.Println(":8080 started")
	if err = http.ListenAndServe(":8080", middleware.RateLimitMiddleware(mux, limiter)); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
