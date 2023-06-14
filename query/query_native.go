package query

import (
	"context"
	"fmt"
	"time"

	"canto-api/config"

	"github.com/redis/go-redis/v9"

	"github.com/Canto-Network/Canto/v6/x/csr/types"

	"google.golang.org/grpc"

	"io/ioutil"
	"net/http"
)

type NativeQueryEngine struct {
	redisclient *redis.Client
	GrpcClient  *grpc.ClientConn
	interval    time.Duration
}

// Returns a NativeQueryEngine instance
func NewNativeQueryEngine() *NativeQueryEngine {
	return &NativeQueryEngine{
		redisclient: config.RDB,
		GrpcClient:  config.GrpcClient,
		interval:    time.Duration(config.QueryInterval),
	}
}

// StartNativeQueryEngine starts the query engine and runs the ticker
// on the interval specified in config
func (nqe *NativeQueryEngine) StartQueryEngine(ctx context.Context) {
	ticker := time.NewTicker(nqe.interval * time.Second)
	for range ticker.C {
		// test calls
		csrQueryHandler := types.NewQueryClient(nqe.GrpcClient)
		
		startGRPC := time.Now();
		resp, err := csrQueryHandler.CSRs(ctx, &types.QueryCSRsRequest{})
		fmt.Println("GRPC call took: ", time.Since(startGRPC))

		//REST
		startRest := time.Now();
		resp1, err := http.Get("https://mainnode.plexnode.org:1317/canto/v1/csr/csrs")
		fmt.Println("REST call took: ", time.Since(startRest))

		if err != nil {
			fmt.Println(err)
		}
		_, err = ioutil.ReadAll(resp1.Body)
		if err != nil {
			fmt.Println(resp)
		}
		// fmt.Println(string(body))
	}
}

// RunNative initializes a NativeQueryEngine and starts it
func RunNative(ctx context.Context) {
	nqe := NewNativeQueryEngine()
	nqe.StartQueryEngine(ctx)
}
