package serve

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"

	"canto-api/config"
)

func GetSmartContractDataGin(ctx *gin.Context) {
	rdb := config.RDB

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	ctx.JSON(http.StatusOK, gin.H{
		"message": val,
	})
}

func GetSmartContractDataFiber(ctx *fiber.Ctx) error {
	rdb := config.RDB

	val, err := rdb.Get(context.Background(), "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	return ctx.SendString(val)
}
