package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Contract struct {
	Name    string
	Address string
	Keys    []string
	Methods []string
	Args    [][]interface{}
}

// parses contracts.json and returns calls
func getContractsFromJson(pathToContractsJsonFile string) ([]Contract, error) {
	var calls []Contract

	// Open the JSON file
	file, err := os.Open(pathToContractsJsonFile)
	if err != nil {
		return nil, errors.New("Config::getContractsFromJson - " + err.Error())
	}
	defer file.Close()

	// Read the file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("Config::getContractsFromJson - " + err.Error())
	}

	//Umarshall the data and store in calls
	unmarshallErr := json.Unmarshal(fileContent, &calls)

	if unmarshallErr != nil {
		fmt.Println("Error:", unmarshallErr)
		return nil, errors.New("Config::getContractsFromJson - " + err.Error())
	}
	return calls, nil

}
