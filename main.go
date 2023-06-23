package main

import (
	"context"
	"fmt"
	"strings"

	"canto-api/config"
	cqe "canto-api/query/contracts"
	nqe "canto-api/query/native"
	re "canto-api/requests"
)

// getGeneralContractRoutes returns a slice of routes for general contracts
func getGeneralContractRoutes() []string {
	routes := []string{}
	for _, contract := range config.ContractCalls {
		for index, method := range contract.Methods {
			// check if the contract has keys
			if len(contract.Keys) == 0 {
				// generate route from name, method and argument of contracts
				route := contract.Name + "/" + strings.Split(method, "(")[0]

				if len(contract.Args[index]) != 0 {
					route += "/" + fmt.Sprintf("%v", contract.Args[index][0])
				}

				routes = append(routes, route)
			}
		}
	}

	return routes
}

func main() {
	config.NewConfig()
	ctx := context.Background()
	go cqe.Run(ctx) // run contract query engine
	go nqe.Run(ctx) // run native query engine
	re.Run(ctx)  // run request engine
}
