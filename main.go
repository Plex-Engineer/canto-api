package main

import (
	"context"

	"canto-api/config"
	cqe "canto-api/queryengine/contracts"
	nqe "canto-api/queryengine/native"
	re "canto-api/requestengine"
)

func main() {
	config.NewConfig()
	ctx := context.Background()
	go cqe.Run(ctx) // run contract query engine
	go nqe.Run(ctx) // run native query engine
	re.Run(ctx)     // run request engine
}
