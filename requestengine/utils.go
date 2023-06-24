package requestengine

import (
	"canto-api/config"
	"context"
	"fmt"
	"strconv"
	"strings"

	cantoConfig "github.com/Canto-Network/Canto/v6/cmd/config"
	"github.com/gofiber/fiber/v2"
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

// CheckValidatorAddress checks if the given address is a valid validator address
func CheckValidatorAddress(address string) error {
	if !(strings.HasPrefix(address, cantoConfig.Bech32PrefixValAddr)) {
		return fmt.Errorf("invalid bech32 validator address: %s", address)
	}
	return nil
}

// CheckIdString checks if the given id is a valid string uint64 id
func CheckIdString(id string) error {
	if _, err := strconv.Atoi(id); err != nil {
		return fmt.Errorf("invalid id: %s", id)
	}
	return nil
}
