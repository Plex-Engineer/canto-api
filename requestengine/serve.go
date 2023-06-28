package requestengine

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"canto-api/config"
)

var (
	StatusBadRequest          = fiber.ErrBadRequest          // 400 (required fields are invalid)
	StatusNotFound            = fiber.ErrNotFound            // 404 (resource do not exist)
	StatusInternalServerError = fiber.ErrInternalServerError // 500 (unexpected error)
	StatusOkay                = fiber.StatusOK               // 200 (success)
)

// functions to return status errors
func RedisKeyNotFound(ctx *fiber.Ctx, key string) error {
	//key there are looking for is not in redis
	return ctx.Status(StatusNotFound.Code).SendString(fmt.Sprintf("%s not found", key))
}
func InvalidParameters(ctx *fiber.Ctx, err error) error {
	//invalid parameters
	return ctx.Status(StatusBadRequest.Code).SendString(err.Error())
}

func GetStoreValueFromKey(key string) (string, error) {
	rdb := config.RDB
	val, err := rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
