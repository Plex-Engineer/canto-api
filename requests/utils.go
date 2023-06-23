package requests

import (
	"fmt"
	"strconv"
	"strings"

	cantoConfig "github.com/Canto-Network/Canto/v6/cmd/config"
)

// CheckValidatorAddress checks if the given address is a valid validator address
func CheckValidatorAddress(address string) error {
	if !(strings.HasPrefix(address, cantoConfig.Bech32PrefixValAddr)) {
		return fmt.Errorf("invalid bech32 validator address: %s", address)
	}
	return nil
}

// CheckIdString checks if the given id is a valid string uint64 id
func CheckIdString(id string) error {
	if _, err := strconv.Atoi(id); err != nil {
		return fmt.Errorf("invalid id: %s", id)
	}
	return nil
}
