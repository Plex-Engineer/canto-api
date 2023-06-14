package query

import (
	"canto-api/config"
	"canto-api/multicall"
	"errors"
)

// gets viewcalls from the contracts

func ProcessContractCalls(contracts []config.Contract) (multicall.ViewCalls, error) {

	vcs := multicall.ViewCalls{}

	for _, contract := range contracts {
		for index, method := range contract.Methods {
			// validate address
			if err := validateAddress(contract.Address); err != nil {
				return nil, err
			}
			vc := multicall.NewViewCall(
				contract.Keys[index],
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
