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
