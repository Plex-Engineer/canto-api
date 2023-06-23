package main

import (
	"context"

	"canto-api/config"
	cqe "canto-api/query/contracts"
	nqe "canto-api/query/native"
	re "canto-api/requests"
)

func main() {
	config.NewConfig()
	ctx := context.Background()
	go cqe.Run(ctx) // run contract query engine
	go nqe.Run(ctx) // run native query engine
	re.Run(ctx)  // run request engine
}
