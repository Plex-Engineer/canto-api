// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    contractsConfig, err := UnmarshalContractsConfig(bytes)
//    bytes, err = contractsConfig.Marshal()

package config

import "encoding/json"

type ContractsConfig map[string]ContractsConfigValue

func UnmarshalContractsConfig(data []byte) (ContractsConfig, error) {
	var r ContractsConfig
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContractsConfig) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type ContractsConfigValue struct {
	Comptroller Comptroller `json:"Comptroller"`
	Router      Comptroller `json:"Router"`
	Reservoir   Comptroller `json:"Reservoir"`
}

type Comptroller struct {
	Address string `json:"address"`
}
