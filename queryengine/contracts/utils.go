package queryengine

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	"time"

	"canto-api/config"
	"canto-api/multicall"

	"errors"
	"regexp"
)

var SecondsPerBlock float64 = 5.8
var BlocksPerDay float64 = 86400 / SecondsPerBlock
var DaysPerYear float64 = 365
var BlocksPerYear float64 = BlocksPerDay * DaysPerYear

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

// This function takes  an interface value and returns a boolean
func InterfaceToBool(value interface{}) (bool, error) {
	//Convert interface{} type to bool
	if boolean, ok := value.(bool); ok {
		return boolean, nil
	}
	return false, errors.New("QueryEngine::InterfaceToBool - Interface value is not a boolean")
}

// FormatUnits() scales down big Int value by 1e(decimals) and returns a float64 value
func FormatUnits(value *big.Int, decimals int64) float64 {
	//divide  value with 1e(decimals) using Quo() method
	formattedValue := new(big.Float).Quo(new(big.Float).SetInt(value), new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)))
	// convert formattedValue to float
	formattedFloatValue, _ := formattedValue.Float64()
	return formattedFloatValue
}

// This function takes  an interface value and returns a bigInt number
func InterfaceToBigInt(value interface{}) (*big.Int, error) {
	// Declare a num variable to store bigInt
	var num = new(big.Int)

	// convert value to string type
	str, err := InterfaceToString(value)
	if err != nil {
		return num, errors.New("InterfaceToBigInt" + err.Error())
	}

	// assign bigInt value to num from its string representation
	num.SetString(str, 10)

	return num, nil
}

// Takes reserve1 and reserve2 bigInt inputs and returns their ratio scaled by 1e18 and boolean a value which is true if reserves1 >= reserves2, false otherwise
func GetLpPairRatio(reserve1 *big.Int, reserve2 *big.Int) (*big.Int, bool) {
	// Check if either reserve1 or reserve2 is nil/zero
	if reserve1 == nil || reserve2 == nil || reserve1.Cmp(big.NewInt(0)) == 0 || reserve2.Cmp(big.NewInt(0)) == 0 {
		return new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil), true // Return [1e18, true] if either is nil/zero
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

// convert big.Int to big.float
func BigIntToFloat64(value *big.Int) *big.Float {
	return new(big.Float).SetInt(value)
}

// holiday mapping for day to days interest update
var holidayMap = map[string]float64{
	"30" + time.December.String(): 4,
	"13" + time.January.String():  4,
	"17" + time.February.String(): 4,
	"25" + time.May.String():      4,
}

// get amount of days passed for hashnote interest update (includes holidy schedule)
func GetInterestDaysPassed(updatedAt int64) float64 {
	// get time from UNIX timestamp
	unixTime := time.Unix(updatedAt, 0)
	// check for holidays
	daysOffset := holidayMap[fmt.Sprintf("%d%s", unixTime.Day(), unixTime.Month().String())]
	if daysOffset != 0 {
		return daysOffset
	}
	// check if saturday reporting (for three days)
	if unixTime.Weekday() == time.Saturday || unixTime.Weekday() == time.Sunday {
		return 3
	}
	return 1
}

// get hashnote apy
func HashnoteAPY(balance *big.Int, interest *big.Int, updatedAt *big.Int) float64 {
	prevBalance := new(big.Float).Sub(BigIntToFloat64(balance), BigIntToFloat64(interest))
	spotAPY, _ := new(big.Float).Quo(BigIntToFloat64(interest), prevBalance).Float64()
	dayOffset := GetInterestDaysPassed(updatedAt.Int64())
	return (spotAPY * 365 * 100) / dayOffset
}

// APY takes the block rate, calculates APY and returns
func APY(blockRate *big.Int) float64 {
	// format blockRate by 1e18
	formattedBlockRate := FormatUnits(blockRate, 18)
	return (math.Pow(formattedBlockRate*BlocksPerDay+1, float64(DaysPerYear)) - 1) * 100
}

// APR takes the block rate, calculates APR and returns
func APR(blockRate *big.Int) float64 {
	// format blockRate by 1e18
	formattedBlockRate := FormatUnits(blockRate, 18)
	return (formattedBlockRate * BlocksPerYear) * 100
}

// distributionAPY takes the block rate, calculates distAPY and returns
func distributionAPY(compSupplySpeed float64, tokenSupply float64, tokenPrice float64, cantoPrice float64) float64 {
	if tokenSupply == 0 || tokenPrice == 0 {
		return 0
	}
	return ((compSupplySpeed * (BlocksPerDay * DaysPerYear)) / tokenSupply) * (cantoPrice / tokenPrice) * 100
}

// This function takes unprocessed pairs data, calculates, adds additional required data and returns the processed pair data
func GetProcessedPairs(ctx context.Context, blocknumber string, pairs PairsMap) ([]ProcessedPair, map[string]string) {
	processedPairs := []ProcessedPair{}
	processedPairsMap := make(map[string]string)

	processedPairsMap["blocknumber"] = blocknumber

	// loop over all pairs
	// key is address of lp pair and value is a map of pair data
	for address, pair := range pairs {
		// get all the data and process
		reserve1, _ := InterfaceToBigInt(pair["reserves"][0])
		reserve2, _ := InterfaceToBigInt(pair["reserves"][1])
		totalSupply, _ := InterfaceToBigInt(pair["totalSupply"][0])
		price1, _ := InterfaceToString(pair["underlyingPriceTokenA"][0])
		price2, _ := InterfaceToString(pair["underlyingPriceTokenB"][0])
		lpPrice, _ := InterfaceToBigInt(pair["underlyingPriceLp"][0])

		// calculate total value locked by multiplying lp price and total supply using Mul() method
		var tvl = new(big.Float).Mul(new(big.Float).SetInt(lpPrice), new(big.Float).SetInt(totalSupply))

		// divide tvl with 1e18 using Quo() method
		tvl.Quo(tvl, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)))

		// get ratio of reserve and reserve2
		ratio, aTob := GetLpPairRatio(reserve1, reserve2)

		// get lp pair data
		symbol, decimals, token1, token2, stable, cDecimals, cLpAddress, logoURI := config.GetLpPairData(address)
		processedPair := ProcessedPair{
			Address:     address,
			Symbol:      symbol,
			Decimals:    decimals,
			Token1:      token1,
			Token2:      token2,
			Stable:      stable,
			CDecimal:    cDecimals,
			CLpAddress:  cLpAddress,
			TotalSupply: totalSupply.String(),
			Tvl:         fmt.Sprintf("%.2f", tvl),
			Ratio:       ratio.String(),
			AToB:        aTob,
			Price1:      price1,
			Price2:      price2,
			LpPrice:     lpPrice.String(),
			Reserve1:    reserve1.String(),
			Reserve2:    reserve2.String(),
			LogoURI:     logoURI,
		}

		processedPairs = append(processedPairs, processedPair)
		processedPairsMap[address] = ResultToString(processedPair)
	}
	return processedPairs, processedPairsMap
}

// This function takes unprocessed ctokens data, calculates, adds additional required data and returns the processed ctokens data
func GetProcessedCTokens(ctx context.Context, cTokens TokensMap) ([]ProcessedCToken, map[string]string) {
	processedCTokens := []ProcessedCToken{}
	processedCTokensMap := make(map[string]string)

	// get ccanto address by symbol
	cCantoAddress := config.GetCTokenAddressBySymbol("cCANTO")
	// get canto price from cTokens data
	cantoPrice, _ := InterfaceToBigInt(cTokens[cCantoAddress]["underlyingPrice"][0])

	// format canto price by 1e18
	formattedCantoPrice := FormatUnits(cantoPrice, 18)

	// loop over all cTokens
	// key is address of cToken and value is a map of cToken data
	for address, cToken := range cTokens {
		// get cToken data
		symbol, name, decimals, tags, underlying := config.GetCTokenData(address)

		// process data
		cash, _ := InterfaceToBigInt(cToken["cash"][0])
		exchangeRate, _ := InterfaceToString(cToken["exchangeRateStored"][0])
		isListed, _ := InterfaceToBool(cToken["markets"][0])
		collateralFactor, _ := InterfaceToString(cToken["markets"][1])
		price, _ := InterfaceToBigInt(cToken["underlyingPrice"][0])
		borrowCap, _ := InterfaceToBigInt(cToken["borrowCaps"][0])
		compSupplyState, _ := InterfaceToString(cToken["compSupplyState"][0])

		if borrowCap.Cmp(big.NewInt(0)) == 0 {
			borrowCap = big.NewInt(math.MaxInt64 - 1)
		}

		// calculate liquidity by multiplying cash and price using Mul() method
		liquidity := new(big.Float).Mul(new(big.Float).SetInt(cash), new(big.Float).SetInt(price))

		// divide liquidity by 1e36 using Quo() method
		liquidity.Quo(liquidity, new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(36), nil)))

		// get supplyApy using APY()
		supplyBlockRate, _ := InterfaceToBigInt(cToken["supplyRatePerBlock"][0])
		supplyApy := APY(supplyBlockRate)
		supplyApr := APR(supplyBlockRate)

		// check tags that may affect this supply rate number
		for _, tag := range tags {
			if tag == "hashnote" {
				// use latest round details to calculate supplyApy
				balance, _ := InterfaceToBigInt(cToken["latestRoundDetails"][1])
				interest, _ := InterfaceToBigInt(cToken["latestRoundDetails"][2])
				updatedAt, _ := InterfaceToBigInt(cToken["latestRoundDetails"][4])
				supplyApy = HashnoteAPY(balance, interest, updatedAt) - 0.5
			}
			if tag == "fbill" {
				supplyApy = 4.90
			}
		}

		// get borrowApy using APY()
		borrowBlockRate, _ := InterfaceToBigInt(cToken["borrowRatePerBlock"][0])
		borrowApy := APY(borrowBlockRate)
		borrowApr := APR(borrowBlockRate)
		compSupplySpeed, _ := InterfaceToBigInt(cToken["compSupplySpeeds"][0])
		// format compSupplySpeed by 1e18
		formattedCompSupplySpeed := FormatUnits(compSupplySpeed, 18)

		// format cash by 1e(decimals)
		formattedTokenSupply := FormatUnits(cash, underlying.Decimals)

		// format token price by 1e(36-decimals)
		formattedTokenPrice := FormatUnits(price, int64(36)-underlying.Decimals)

		distApy := distributionAPY(formattedCompSupplySpeed, formattedTokenSupply, formattedTokenPrice, formattedCantoPrice)

		// Set price of cNOTE, cUSDC, cUSDT to exactly 1USD scaled by 1e(36-decimals)
		if symbol == "cNOTE" || symbol == "cUSDC" || symbol == "cUSDT" {
			price.Exp(big.NewInt(10), big.NewInt(36-underlying.Decimals), nil)
		}

		// get underlying total supply
		underlyingTotalSupply, _ := InterfaceToString(cToken["underlyingSupply"][0])

		processedCToken := ProcessedCToken{
			Address:               address,
			Symbol:                symbol,
			Name:                  name,
			Decimals:              decimals,
			Underlying:            underlying,
			Cash:                  cash.String(),
			ExchangeRate:          exchangeRate,
			IsListed:              isListed,
			CollateralFactor:      collateralFactor,
			Price:                 price.String(),
			BorrowCap:             borrowCap.String(),
			Liquidity:             fmt.Sprintf("%.2f", liquidity),
			SupplyApy:             fmt.Sprintf("%.2f", supplyApy),
			SupplyApr:             fmt.Sprintf("%.2f", supplyApr),
			BorrowApy:             fmt.Sprintf("%.2f", borrowApy),
			BorrowApr:             fmt.Sprintf("%.2f", borrowApr),
			DistApy:               fmt.Sprintf("%.2f", distApy),
			CompSupplyState:       compSupplyState,
			UnderlyingTotalSupply: underlyingTotalSupply,
		}

		processedCTokens = append(processedCTokens, processedCToken)
		processedCTokensMap[address] = ResultToString(processedCToken)
	}
	return processedCTokens, processedCTokensMap
}
