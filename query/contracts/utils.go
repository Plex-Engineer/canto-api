package query

import (
	"canto-api/multicall"
	"encoding/json"

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
