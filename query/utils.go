package query

import (
	"canto-api/config"
	"canto-api/multicall"
	"encoding/json"

	"errors"
	"regexp"
)

func ProcessContractCalls(contracts []config.Contract) (multicall.ViewCalls, error) {

	vcs := multicall.ViewCalls{}

	for _, contract := range contracts {
		for index, method := range contract.Methods {
			// validate address
			if err := validateAddress(contract.Address); err != nil {
				return nil, err
			}
			vc := multicall.NewViewCall(
				contract.Names[index],
				contract.Address,
				method,
				contract.Args[index],
			)

			if err := vc.Validate(); err != nil {
				return nil, errors.New("QueryEngine::ProcessContractCalls - " + err.Error())
			}

			vcs = append(vcs, vc)
		}
	}

	return vcs, nil
}

func ResultToString(results *multicall.Result) string {
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
		return errors.New("QueryEngine::ValidateAddress - invalid address")
	}
	return nil
}
