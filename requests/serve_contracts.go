package requests

import (
	"canto-api/config"
	"canto-api/rediskeys"
	"context"

	"github.com/gofiber/fiber/v2"
)

// Processed Pairs
func QueryPairs(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey(rediskeys.ProcessedPairs))
}

func QueryPairsByAddress(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), rediskeys.ProcessedPairsMap, ctx.Params("address")).Result()
	if err != nil {
		val = "CSR not found"
	}
	return ctx.SendString(val)
}
