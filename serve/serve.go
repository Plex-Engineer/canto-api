package serve

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"canto-api/config"
)

func GetSmartContractDataFiber(ctx *fiber.Ctx) error {
	rdb := config.RDB

	val, err := rdb.Get(context.Background(), "key").Result()
	if err != nil {
		panic(err)
	}
	return ctx.SendString(val)
}

func getStoreValueFromKey(key string) string {
	rdb := config.RDB
	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		panic(err)
	}
	return val
}

func QueryStakingAPR(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey("stakingApr"))
}

func QueryValidators(ctx *fiber.Ctx) error {
	return ctx.SendString(getStoreValueFromKey("validators"))
}
