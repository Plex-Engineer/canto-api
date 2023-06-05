package query

import (
	"canto-api/config"
	"canto-api/multicall"
	"context"
	"encoding/json"
	"log"

	"github.com/redis/go-redis/v9"
)

func ProcessContractCalls(contracts []config.Contract) (multicall.ViewCalls, error) {

	vcs := multicall.ViewCalls{}

	for _, contract := range contracts {
		for index, method := range contract.Methods {
			vc := multicall.NewViewCall(
				contract.Address,
				method,
				contract.Args[index],
			)

			if err := vc.Validate(); err != nil {
				return nil, err
			}

			vcs = append(vcs, vc)
		}
	}

	return vcs, nil
}

func SetCacheWithResult(ctx context.Context, redisclient *redis.Client, results *multicall.Result) error {

	ret := ResultToString(results)

	// set key in redis
	err := redisclient.Set(ctx, "key", string(ret), 0).Err()
	if err != nil {
		panic(err)
	}

	return nil
}

func ResultToString(results *multicall.Result) string {
	ret, _ := json.Marshal(results)
	return string(ret)
}

func GetCallData(vcs multicall.ViewCalls) []multicall.Multicall3Call {
	payload, err := vcs.GetCallData()
	if err != nil {
		log.Fatal(err)
	}
	return payload
}
