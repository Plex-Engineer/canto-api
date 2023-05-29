package main

import (
	"fmt"
	"net/http"

	"canto-api/query"
	"canto-api/serve"
)

func main() {
	go query.Tick() // run query engine in background

	http.HandleFunc("/", serve.GetSmartContractData)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
