package usecase

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/paulnune/goexpert-ratelimiter/internal/database"
)

// RateLimiter struct to hold rate limiting data
type RateLimiter struct {
	ctx           context.Context
	requests      map[string]int
	mu            sync.Mutex
	limit         int
	interval      time.Duration
	blockInterval time.Duration
	listTokens    database.TokenLimitList
	client        database.Persist
}

// NewIpRateLimiter creates a new rate limiter
func NewIpRateLimiter(ctx context.Context, limit int, interval time.Duration, blockInterval time.Duration, listTokens database.TokenLimitList, client database.Persist) *RateLimiter {
	rl := &RateLimiter{
		ctx:           ctx,
		requests:      make(map[string]int),
		limit:         limit,
		interval:      interval,
		blockInterval: blockInterval,
		listTokens:    listTokens,
		client:        client,
	}
	go rl.cleanup()
	return rl
}

// Allow checks if the request is allowed
func (rl *RateLimiter) Allow(ip, token string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	var ipRequests database.IpRequests

	result, err := rl.client.Get(rl.ctx, ip)
	if err != nil {
		// log.Println("first req")
		json, _ := json.Marshal(database.NewRequest(ip, 1, 0))
		rl.client.Set(rl.ctx, ip, json)
		return true
	}

	json.Unmarshal([]byte(result), &ipRequests)
	if ipRequests.BlockUntil > 0 {
		timeToRelase := time.Unix(ipRequests.BlockUntil, 0).
			Add(rl.blockInterval)
		if timeToRelase.After(time.Now()) {
			return false
		} else {
			json, _ := json.Marshal(database.NewRequest(ip, 1, 0))
			rl.client.Set(rl.ctx, ip, json)
			return true
		}
	}

	requestsNow := ipRequests.Qty

	if token != "" && rl.listTokens.GetLimit(token) > 0 {
		tokenLimit := rl.listTokens.GetLimit(token)
		// resultToken, _ := rl.client.Get(rl.ctx, token).Result()
		// json.Unmarshal([]byte(resultToken), &tokenLimit)
		if requestsNow >= tokenLimit {
			json, _ := json.Marshal(database.NewRequest(ip, 0, time.Now().Unix()))
			rl.client.Set(rl.ctx, ip, json)
			return false
		} else {
			requestsNow++
			json, _ := json.Marshal(database.NewRequest(ip, requestsNow, 0))
			rl.client.Set(rl.ctx, ip, json)
			return true
		}
	}

	if requestsNow >= rl.limit {
		json, _ := json.Marshal(database.NewRequest(ip, 0, time.Now().Unix()))
		rl.client.Set(rl.ctx, ip, json)
		return false
	}

	requestsNow++
	json, _ := json.Marshal(database.NewRequest(ip, requestsNow, 0))
	rl.client.Set(rl.ctx, ip, json)
	return true
}

// cleanup resets the rate limit counts at regular intervals
func (rl *RateLimiter) cleanup() error {
	for {
		rl.mu.Lock()
		// for k := range rl.requests {
		// 	delete(rl.requests, k)
		// }
		pattern := "*.*.*.*"
		// cursor := uint64(0)

		lista, err := rl.client.Keys(rl.ctx, pattern)
		if err != nil {
			log.Panicln(err.Error())
			return err
		}
		if len(lista) > 0 {
			for _, key := range lista {
				_, err = rl.client.Del(rl.ctx, key)
				if err != nil {
					return err
				}
			}
		}
		rl.mu.Unlock()
		time.Sleep(rl.interval)
	}
}
