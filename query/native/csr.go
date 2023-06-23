package query

import (
	"context"
	"strconv"

	csr "github.com/Canto-Network/Canto/v6/x/csr/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
)

// CSR
type CSR struct {
	// ID of the CSR
	Id uint64 `json:"id"`
	// all contracts under this csr id
	Contracts []string `json:"contracts"`
	// total number of transactions under this csr id
	Txs uint64 `json:"txs"`
	// The cumulative revenue for this CSR NFT -> represented as a big.Int
	Revenue string `json:"revenue"`
}

// get all CSRS
// will return full response string and mapping of nft id to response string
func GetCSRS(ctx context.Context, queryClient csr.QueryClient) ([]CSR, map[string]string) {
	resp, err := queryClient.CSRs(ctx, &csr.QueryCSRsRequest{Pagination: &query.PageRequest{
		Limit: 500,
	}})
	CheckError(err)
	allCsrs := new([]CSR)
	csrMap := make(map[string]string)
	for _, csr := range resp.GetCsrs() {
		csrResponse := CSR{
			Id:        csr.GetId(),
			Contracts: csr.GetContracts(),
			Txs:       csr.GetTxs(),
			Revenue:   csr.Revenue.String(),
		}
		*allCsrs = append(*allCsrs, csrResponse)
		csrMap[strconv.Itoa(int(csr.GetId()))] = GeneralResultToString(csrResponse)
	}
	return *allCsrs, csrMap
}
