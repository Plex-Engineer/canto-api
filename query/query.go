package query

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// QueryEngine queries smart contracts directly from a node
// and stores the data in a Redis database on a regular interval.
type QueryEngine struct {
	client   *redis.Client
	interval time.Duration
}

func NewQueryEngine() *QueryEngine {
	return &QueryEngine{
		client: redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0, // use default DB
		}),
		interval: 5,
	}
}

func (qe *QueryEngine) StartQueryEngine(ctx context.Context) {
	x := 0

	ticker := time.NewTicker(qe.interval * time.Second)
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
