package query

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type QueryEngine struct {
	client  *redis.Client
	seconds time.Duration
}

func NewQueryEngine() *QueryEngine {
	return &QueryEngine{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0, // use default DB
		}),
		seconds: 5,
	}
}

func (qe *QueryEngine) StartQueryEngine(ctx context.Context) {
	x := 0

	ticker := time.NewTicker(qe.seconds * time.Second)
	for range ticker.C {
		fmt.Println("Call smart contract data", x)
		err := qe.client.Set(ctx, "key", x, 0).Err()
		if err != nil {
			panic(err)
		}
		x++
	}
}

func Run(ctx context.Context) {
	qe := NewQueryEngine()
	qe.StartQueryEngine(ctx)
}
