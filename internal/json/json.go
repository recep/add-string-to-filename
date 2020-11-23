package json

import (
	"encoding/json"
	"io/ioutil"
)

// TODO Refactor this function
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

// TODO function => READ from a json file
