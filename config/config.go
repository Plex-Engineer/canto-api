package config

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
)

var (
	RDB       *redis.Client
	EthClient *ethclient.Client
)

func NewConfig() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,
	})
	EthClient, _ = ethclient.Dial("https://mainnode.plexnode.org:8545")
}
