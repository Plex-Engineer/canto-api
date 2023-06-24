package query

import (
	"context"
	"encoding/json"
	"math/big"

	"canto-api/config"
	"canto-api/multicall"

	"errors"
	"regexp"
)

func ResultToString(results interface{}) string {
	ret, err := json.Marshal(results)
	if err != nil {
		return "QueryEngine::ResultToString - " + err.Error()
	}
	return string(ret)
}

func GetCallData(vcs multicall.ViewCalls) ([]multicall.Multicall3Call, error) {
	payload, err := vcs.GetCallData()
	if err != nil {
		return nil, errors.New("QueryEngine::GetCallData - " + err.Error())
	}
	return payload, nil
}

func validateAddress(address string) error {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(address) {
		return errors.New("QueryEngine::ValidateAddress - invalid address" + address)
	}
	return nil
}

func GeneralResultToString(results interface{}) string {
	ret, err := json.Marshal(results)
	if err != nil {
		return "QueryEngine::GeneralResultToString - " + err.Error()
	}
	return string(ret)
}

// This function takes an interface value, does a type assertion to string and returns it
func InterfaceToString(value interface{}) (string, error) {
	//Convert interface{} type to string
	if str, ok := value.(string); ok {
		return str, nil
	}
	return "", errors.New("QueryEngine::InterfaceToString - Interface value is not a string")
}

// This function takes  an interface value and returns a bigInt number
func InterfaceToBigInt(value interface{}) (*big.Int, error) {
	// Declare a num variable to store bigInt
	var num = new(big.Int)

	// convert value to string type
	str, err := InterfaceToString(value)
	if err != nil {
		return num, errors.New("QueryEngine::InterfaceToBigInt -" + err.Error())
	}

	// assign bigInt value to num from its string representation
	num.SetString(str, 10)

	return num, nil
}

// Takes reserve1 and reserve2 bigInt inputs and returns their ratio scaled by 1e18 and boolean a value which is true if reserves1 >= reserves2, false otherwise
func GetLpPairRatio(reserve1 *big.Int, reserve2 *big.Int) (*big.Int, bool) {
	// Check if either reserve1 or reserve2 is nil
	if reserve1 == nil || reserve2 == nil {
		return big.NewInt(1), true // Return [1, true] if either is nil
	}

	// check if reserve1 is greater than or equal to reserve2
	if reserve1.Cmp(reserve2) >= 0 {
		// Calculate reserve1 to reserve2 ratio and scale it by 1e18
		ratio := new(big.Int).Mul(reserve1, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
		ratio.Div(ratio, reserve2)
		return ratio, true
	} else {
		// Calculate reserve2 to reserve1 ratio and scale it by 1e18
		ratio := new(big.Int).Mul(reserve2, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
		ratio.Div(ratio, reserve1)
		return ratio, false
	}
}

// This function takes symbol and data of a pair(ex:CantoAtomLP), calculates and adds additional required data and returns the processed pair data
func GetProcessedPairs(ctx context.Context, pairs PairsMap) ([]ProcessedPair, map[string]string) {
	processedPairs := []ProcessedPair{}
	processedPairsMap := make(map[string]string)

	// loop over all pairs
	// key is address of lp pair and value is a map of pair data
	for key, value := range pairs {
		// get all the data and process
		reserve1, _ := InterfaceToBigInt(value["reserves"][0])
		reserve2, _ := InterfaceToBigInt(value["reserves"][1])
		totalSupply, _ := InterfaceToBigInt(value["totalSupply"][0])
		price1, _ := InterfaceToString(value["underlyingPriceTokenA"][0])
		price2, _ := InterfaceToString(value["underlyingPriceTokenB"][0])
		lpPrice, _ := InterfaceToBigInt(value["underlyingPriceLp"][0])

		// calculate total value locked by multiplying lp price and total supply using Mul() method
		var tvl = new(big.Int).Mul(lpPrice, totalSupply)

		// divide tvl with 1e18 using Div() method
		tvl.Div(tvl, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))

		// get ratio of reserve and reserve2
		ratio, aTob := GetLpPairRatio(reserve1, reserve2)

		// get lp pair data
		symbol, decimals, token1, token2, stable, cDecimals, cLpAddress := config.GetLpPairData(key)
		processedPair := ProcessedPair{
			Address:     key,
			Symbol:      symbol,
			Decimals:    decimals,
			Token1:      token1,
			Token2:      token2,
			Stable:      stable,
			CDecimal:    cDecimals,
			CLpAddress:  cLpAddress,
			TotalSupply: totalSupply.String(),
			Tvl:         tvl.String(),
			Ratio:       ratio.String(),
			AToB:        aTob,
			Price1:      price1,
			Price2:      price2,
			LpPrice:     lpPrice.String(),
			Reserve1:    reserve1.String(),
			Reserve2:    reserve2.String(),
		}

		processedPairs = append(processedPairs, processedPair)
		processedPairsMap[key] = ResultToString(processedPair)
	}
	return processedPairs, processedPairsMap
}
