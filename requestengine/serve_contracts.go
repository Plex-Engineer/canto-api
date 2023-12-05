package requestengine

import (
	"context"
	"encoding/json"
	"strings"

	"canto-api/config"
	queryengine "canto-api/queryengine/contracts"

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

// QueryPairs godoc
// @Summary      Query all pairs in Canto dex
// @Description  return json array of all pairs in Canto dex
// @Accept       json
// @Produce      json
// @Success      200  {object}  Pairs
// @Router       /dex/pairs [get]
func QueryPairs(ctx *fiber.Ctx) error {
	// get pairs json string from cache
	pairsString, err := GetStoreValueFromKey(config.ProcessedPairs)
	if err != nil {
		return RedisKeyNotFound(ctx, config.ProcessedPairs)
	}

	return ctx.Status(StatusOkay).SendString(pairsString)
}

// QueryPairByAddress godoc
// @Summary      Query a pair by address
// @Description  return json array of all pairs in Canto dex
// @Accept       json
// @Produce      json
// @Success      200  {object}  Pairs
// @Router       /dex/pairs/{address} [get]
func QueryPairByAddress(ctx *fiber.Ctx) error {
	// get block number from cache
	blockNumber, err := GetBlockNumber()
	if err != nil {
		return RedisKeyNotFound(ctx, config.BlockNumber)
	}

	// get pair json string from cache
	pairString, err := config.RDB.HGet(context.Background(), config.ProcessedPairsMap, ctx.Params("address")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, config.ProcessedPairsMap)
	}

	// unmarhsall pair
	var pair queryengine.ProcessedPair
	json.Unmarshal([]byte(pairString), &pair)

	// generate json result string
	result := queryengine.ResultToString(map[string]interface{}{
		"block": blockNumber,
		"results":  pair,
	})
	return ctx.Status(StatusOkay).SendString(result)
}

// QueryCTokens godoc
// @Summary      Query all cTokens in CLM
// @Description  return json array of all pairs in CLM
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /lending/ctokens [get]
func QueryCTokens(ctx *fiber.Ctx) error {
	// get cTokens json string from cache
	cTokensString, err := GetStoreValueFromKey(config.ProcessedCTokens)
	if err != nil {
		return RedisKeyNotFound(ctx, config.ProcessedCTokens)
	}

	return ctx.Status(StatusOkay).SendString(cTokensString)
}

// QueryCTokenByAddress godoc
// @Summary      Query cToken by address
// @Description  return json object of cToken
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Router       /lending/ctokens/{address} [get]
func QueryCTokenByAddress(ctx *fiber.Ctx) error {
	// get block number from cache
	blockNumber, err := GetBlockNumber()
	if err != nil {
		return RedisKeyNotFound(ctx, config.BlockNumber)
	}

	// get cToken json string from cache
	cTokenString, err := config.RDB.HGet(context.Background(), config.ProcessedCTokensMap, ctx.Params("address")).Result()
	if err != nil {
		return RedisKeyNotFound(ctx, config.ProcessedCTokensMap)
	}

	// unmarhsall cToken
	var cToken queryengine.ProcessedCToken
	json.Unmarshal([]byte(cTokenString), &cToken)

	// generate json result string
	result := queryengine.ResultToString(map[string]interface{}{
		"block":  blockNumber,
		"results": cToken,
	})
	return ctx.Status(StatusOkay).SendString(result)
}
