// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    contractsConfig, err := UnmarshalContractsConfig(bytes)
//    bytes, err = contractsConfig.Marshal()

package config

import "encoding/json"

type ContractsInfo map[string]ContractsConfigValue

func UnmarshalContractsConfig(data []byte) (ContractsInfo, error) {
	var r ContractsInfo
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *ContractsInfo) Marshal() ([]byte, error) {
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
