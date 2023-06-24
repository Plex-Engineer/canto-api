package queryengine

import (
	"context"
	"errors"

	"canto-api/config"
)

// SetJsonToCache will take key, result and sets the resulte as a json string to redis
func (qe *QueryEngine) SetJsonToCache(ctx context.Context, key string, result interface{}) error {
	// convert result to json string
	ret := ResultToString(result)
	err := qe.redisclient.Set(ctx, key, ret, 0).Err()
	if err != nil {
		return errors.New("SetJsonToCache:" + err.Error())
	}
	return nil
}

// SetMapToCache will take key, result map and sets to redis using HSet()
func (qe *QueryEngine) SetMapToCache(ctx context.Context, key string, result map[string]string) error {
	//set key in redis
	err := qe.redisclient.HSet(ctx, key, result).Err()
	if err != nil {
		return errors.New("SetMapToCache:" + err.Error())
	}
	return nil
}

// SetCacheWithFpi sets the result of ctokens and pairs in Redis
// and returns an error if any occurs.
func (qe *QueryEngine) SetCacheWithFpi(ctx context.Context, ctokens TokensMap, pairs PairsMap) error {
	// set ctokens as a json string in redis
	err := qe.SetJsonToCache(ctx, config.CTokens, ctokens)
	if err != nil {
		return errors.New("SetCacheWithFpi:" + err.Error())
	}

	// set pairs as a json string in redis
	err = qe.SetJsonToCache(ctx, config.Pairs, pairs)
	if err != nil {
		return errors.New("SetCacheWithFpi:" + err.Error())
	}

	return nil
}

// SetCacheWithResult sets the result of a multicall query in Redis
// and returns an error if any occur.
func (qe *QueryEngine) SetCacheWithGeneral(ctx context.Context, results map[string][]interface{}) error {
	// iterate others map and set keys in redis
	for key, value := range results {
		// convert result slice to string
		ret := ResultToString(value)
		// set key in redis
		err := qe.redisclient.Set(ctx, key, string(ret), 0).Err()
		if err != nil {
			return errors.New("SetCacheWithResult:" + err.Error())
		}
	}
	return nil
}

// This function gets the pairs data from redis, processes it and sets the processed pairs data to redis
func (qe *QueryEngine) SetCacheWithProcessedPairs(ctx context.Context, pairs PairsMap) error {
	// get processed pairs data
	processedPairs, processedPairsMap := GetProcessedPairs(ctx, pairs)

	// set processed pairs as a json string to redis
	err := qe.SetJsonToCache(ctx, config.ProcessedPairs, processedPairs)
	if err != nil {
		return errors.New("SetCacheWithProcessedPairs:" + err.Error())
	}

	// set processed pairs map as a json string to redis
	err = qe.SetMapToCache(ctx, config.ProcessedPairsMap, processedPairsMap)
	if err != nil {
		return errors.New("SetCacheWithProcessedPairs:" + err.Error())
	}

	return nil
}

func (qe *QueryEngine) SetCacheWithProcessedCTokens(ctx context.Context, ctokens TokensMap) error {
	// get processed ctokens data
	processedCTokens, processedCTokensMap := GetProcessedCTokens(ctx, ctokens)

	// set processed ctokens as a json string to redis
	err := qe.SetJsonToCache(ctx, config.ProcessedCTokens, processedCTokens)
	if err != nil {
		return errors.New("SetCacheWithProcessedCTokens:" + err.Error())
	}

	// set processed ctokens map as a json string to redis
	err = qe.SetMapToCache(ctx, config.ProcessedCTokensMap, processedCTokensMap)
	if err != nil {
		return errors.New("SetCacheWithProcessedCTokens:" + err.Error())
	}

	return nil
}
