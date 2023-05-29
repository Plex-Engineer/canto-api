package query

import (
	"fmt"
	"time"
)

func Tick() {
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println("Call smart contract data")
		}
	}
}
