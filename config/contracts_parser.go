// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    rawContracts, err := UnmarshalRawContracts(bytes)
//    bytes, err = rawContracts.Marshal()

package config

import "encoding/json"

type RawContracts map[string]RawContractsValue

func UnmarshalRawContracts(data []byte) (RawContracts, error) {
	var r RawContracts
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RawContracts) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RawContractsValue struct {
	Comptroller Comptroller `json:"Comptroller"`
	Router      Comptroller `json:"Router"`
	Reservoir   Comptroller `json:"Reservoir"`
}

type Comptroller struct {
	Address string `json:"address"`
}
