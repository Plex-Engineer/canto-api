package requestengine

import (
	"strings"

	"canto-api/config"
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func GetGeneralContractDataFiber(ctx *fiber.Ctx) error {
	// assemble key from route
	var key string
	route := strings.Split(ctx.Route().Path, `/`)

	for index, part := range route {
		if index > 1 {
			key += ":" + part
		} else if index == 1 {
			key += part
		}
	}

	rdb := config.RDB
	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		log.Error().
			Err(err).
			Msgf("Error getting key '%s' from redis", key)
	}
	return ctx.SendString(val)
}

// Processed Pairs
func QueryPairs(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(config.ProcessedPairs)
	if err != nil {
		return RedisKeyNotFound(ctx, config.ProcessedPairs)
	}
	return ctx.Status(StatusOkay).SendString(val)
}

func QueryPairsByAddress(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), config.ProcessedPairsMap, ctx.Params("address")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, config.ProcessedPairsMap)
	}
	return ctx.SendString(val)
}

func QueryCTokens(ctx *fiber.Ctx) error {
	val, err := GetStoreValueFromKey(config.ProcessedCTokens)
	if err != nil {
		return RedisKeyNotFound(ctx, config.ProcessedCTokens)
	}
	return ctx.SendString(val)
}

func QueryCTokenByAddress(ctx *fiber.Ctx) error {
	val, err := config.RDB.HGet(context.Background(), config.ProcessedCTokensMap, ctx.Params("address")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, config.ProcessedCTokensMap)
	}
	return ctx.SendString(val)
}
