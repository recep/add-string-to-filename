package tools

import (
	"encoding/json"
	"io/ioutil"
)

func Writer(data interface{}, path string) error {
	jsonFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, jsonFile, 0644)
	if err != nil {
		return err
	}

	return nil
}

func Reader(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}
