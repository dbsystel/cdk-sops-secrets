package data

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = map[string]interface{}{
	"simple": map[string]interface{}{
		"string": "examplestring",
		"number": 1.1,
		"bool":   true,
	},
	"nested": map[string]interface{}{
		"string": "examplestring",
		"number": 1.1,
		"bool":   true,
		"nested": map[string]interface{}{
			"string": "examplestring",
			"number": 1.1,
			"bool":   true,
		},
	},
	"array": []interface{}{
		map[string]interface{}{
			"string": "examplestring",
			"number": 1.1,
			"bool":   true,
		},
		map[string]interface{}{
			"string": "examplestring",
			"number": 1.1,
			"bool":   true,
		},
	},
}

func compare(content []byte, filename string) error {
	filePath := "__snapshots__/" + filename
	if os.Getenv("UPDATE_SNAPS") == "true" {
		err := os.WriteFile(filePath, content, 0644)
		if err != nil {
			return err
		}
	}
	existingContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	if string(existingContent) != string(content) {
		return fmt.Errorf("content does not match for file: %s", filename)
	}
	return nil
}

func read(filename string) ([]byte, error) {
	filePath := "__snapshots__/" + filename
	existingContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return existingContent, nil
}

func Test_ToJSON(t *testing.T) {
	for tName, tData := range testData {
		t.Run(tName, func(t *testing.T) {
			data := &Data{
				content: tData,
			}
			jsonData, err := data.toJSON()
			if err != nil {
				t.Errorf("ToJson() error = %v", err)
				return
			}
			filePath := tName + ".json"
			err = compare(jsonData, filePath)
			if err != nil {
				t.Errorf("Compare() error = %v", err)
				return
			}
		})
	}
}

func Test_ToYAML(t *testing.T) {
	for tName, tData := range testData {
		t.Run(tName, func(t *testing.T) {
			data := &Data{
				content: tData,
			}
			yamlData, err := data.toYAML()
			if err != nil {
				t.Errorf("ToYaml() error = %v", err)
				return
			}
			filePath := tName + ".yaml"
			err = compare(yamlData, filePath)
			if err != nil {
				t.Errorf("Compare() error = %v", err)
				return
			}
		})
	}
}

func Test_FromJSON(t *testing.T) {
	for tName, tData := range testData {
		t.Run(tName, func(t *testing.T) {
			jsonData, err := read(tName + ".json")
			if err != nil {
				t.Errorf("Read() error = %v", err)
				return
			}
			data, err := FromJSON(jsonData)
			if err != nil {
				t.Errorf("FromJSON() error = %v", err)
				return
			}
			assert.EqualValues(t, data, &Data{content: tData})
		})
	}
}

func Test_FromYAML(t *testing.T) {
	for tName, tData := range testData {
		t.Run(tName, func(t *testing.T) {
			jsonData, err := read(tName + ".yaml")
			if err != nil {
				t.Errorf("Read() error = %v", err)
				return
			}
			data, err := FromYAML(jsonData)
			if err != nil {
				t.Errorf("FromYAML() error = %v", err)
				return
			}
			assert.EqualValues(t, data, &Data{content: tData})
		})
	}
}
