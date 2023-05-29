package serve

import (
	"fmt"
	"io"
	"net/http"
)

func GetSmartContractData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "Hello World")
}
