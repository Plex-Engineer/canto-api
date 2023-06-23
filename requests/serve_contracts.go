package requests

import (
	"canto-api/config"
	"canto-api/rediskeys"
	"context"

	"github.com/gofiber/fiber/v2"
)

// Processed Pairs
func QueryPairs(ctx *fiber.Ctx) error {
	val, err := getStoreValueFromKey(rediskeys.ProcessedPairs)
	if err != nil {
		return redisKeyNotFound(ctx, rediskeys.ProcessedPairs)
	}
	return ctx.Status(StatusOkay).SendString(val)
}

func QueryPairsByAddress(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), rediskeys.ProcessedPairsMap, ctx.Params("address")).Result()
	if err != nil {
		val = "Pair not found"
	}
	return ctx.SendString(val)
}
