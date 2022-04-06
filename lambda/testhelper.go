package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func UnmarshalAny(input []byte) map[string]interface{} {
	var returnValue map[string]interface{}
	err := json.Unmarshal(input, &returnValue)
	check(err)
	return returnValue
}

func ReadFile(path string) []byte {
	ret, err := os.ReadFile(path)
	check(err)
	return ret
}

func ReadJSONFromFile(t *testing.T, inputFile string) []byte {
	inputJSON, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	return inputJSON
}
