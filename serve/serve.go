package serve

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	redisclient "canto-api/redisclient"
)

func GetSmartContractData(ctx *gin.Context) {
	rdb := redisclient.RDB

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
	ctx.JSON(http.StatusOK, gin.H{
		"message": val,
	})
}
