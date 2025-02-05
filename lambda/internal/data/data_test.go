package data

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

var testData = map[string]interface{}{
	"1.simple": map[string]interface{}{
		"string": "examplestring",
		"number": 1.1,
		"bool":   true,
	},
	"2.nested": map[string]interface{}{
		"string": "examplestring",
		"number": 1.1,
		"bool":   true,
		"nested": map[string]interface{}{
			"string": "examplestring",
			"number": 1.1,
			"bool":   true,
		},
	},
	"3.array": []interface{}{
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

func read(filename string) ([]byte, error) {
	filePath := "__snapshots__/" + filename
	existingContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	lines := bytes.Split(existingContent, []byte("\n"))
	if len(lines) > 2 {
		existingContent = bytes.Join(lines[2:len(lines)-2], []byte("\n"))
	}
	return existingContent, nil
}

func Test_ToJSON(t *testing.T) {
	for tName, tData := range testData {
		t.Run(tName, func(t *testing.T) {
			data := &Data{
				parsed: &tData,
			}
			jsonData, err := data.ToJSON()
			if err != nil {
				t.Errorf("ToJson() error = %v", err)
				return
			}
			snaps.WithConfig(
				snaps.Filename(tName+".toJson"),
			).MatchSnapshot(t, string(*jsonData))
		})
	}
}

func Test_ToJSON_fail(t *testing.T) {
	data := &Data{}
	_, err := data.ToJSON()
	if err != ErrorUnparsedData {
		t.Errorf("Expected error %v, got %v", ErrorUnparsedData, err)
	}
}

func Test_ToYAML(t *testing.T) {
	for tName, tData := range testData {
		t.Run(tName, func(t *testing.T) {
			data := &Data{
				parsed: &tData,
			}
			jsonData, err := data.ToYAML()
			if err != nil {
				t.Errorf("toYAML() error = %v", err)
				return
			}
			snaps.WithConfig(
				snaps.Filename(tName+".toYaml"),
			).MatchSnapshot(t, string(*jsonData))
		})
	}
}

func Test_ToYAML_fail(t *testing.T) {
	data := &Data{}
	_, err := data.ToYAML()
	if err != ErrorUnparsedData {
		t.Errorf("Expected error %v, got %v", ErrorUnparsedData, err)
	}
}

func Test_ToStringMap(t *testing.T) {
	for tName, tData := range testData {
		t.Run(tName, func(t *testing.T) {
			data := &Data{
				parsed: &tData,
			}
			stringMap, err := data.ToStringMap()

			if tName == "3.array" {
				errExp := fmt.Errorf("parsed data is not a map[string]interface{}")
				if err.Error() != errExp.Error() {
					t.Errorf("Expected error %v, got %v", errExp, err)
				}
				return
			}
			if tName == "2.nested" {
				errExp := fmt.Errorf(`value for key nested is a complex type (map or slice), maybe run "Flatten" first`)
				if err.Error() != errExp.Error() {
					t.Errorf("Expected error %v, got %v", errExp, err)
				}
				return
			}
			if err != nil {
				t.Errorf("ToStringMap() error = %v", err)
				return
			}
			snaps.WithConfig(
				snaps.Filename(tName+".toStringMap"),
			).MatchSnapshot(t, stringMap)
		})
	}
}

func Test_FromJSON(t *testing.T) {
	for tName := range testData {
		t.Run(tName, func(t *testing.T) {
			jsonData, err := read(tName + ".toJson.snap")
			if err != nil {
				t.Errorf("Read() error = %v", err)
				return
			}
			data, err := FromJSON(jsonData)
			if err != nil {
				t.Errorf("FromJSON() error = %v", err)
				return
			}
			snaps.WithConfig(
				snaps.Filename(tName+".fromJson"),
			).MatchSnapshot(t, data)
		})
	}
}

func Test_FromYAML(t *testing.T) {
	for tName := range testData {
		t.Run(tName, func(t *testing.T) {
			jsonData, err := read(tName + ".toYaml.snap")
			if err != nil {
				t.Errorf("Read() error = %v", err)
				return
			}
			data, err := FromYAML(jsonData)
			if err != nil {
				t.Errorf("FromYAML() error = %v", err)
				return
			}
			snaps.WithConfig(
				snaps.Filename(tName+".fromYaml"),
			).MatchSnapshot(t, data)
		})
	}
}

func Test_FromDotEnv(t *testing.T) {
	dotEnvData := []byte(`
		# This is a comment
		STRING=examplestring
		NUMBER=1.1
		BOOL=true
		NESTED_STRING=examplestring
		NESTED_NUMBER=1.1
		NESTED_BOOL=true
	`)

	expectedData := map[string]string{
		"STRING":        "examplestring",
		"NUMBER":        "1.1",
		"BOOL":          "true",
		"NESTED_STRING": "examplestring",
		"NESTED_NUMBER": "1.1",
		"NESTED_BOOL":   "true",
	}

	data, err := FromDotEnv(dotEnvData)
	if err != nil {
		t.Errorf("FromDotEnv() error = %v", err)
		return
	}

	if data.parsed == nil {
		t.Errorf("FromDotEnv() parsed data is nil")
		return
	}

	parsedData, ok := (*data.parsed).(map[string]string)
	if !ok {
		t.Errorf("FromDotEnv() parsed data is not of type map[string]string")
		return
	}

	for key, expectedValue := range expectedData {
		if parsedValue, exists := parsedData[key]; !exists || parsedValue != expectedValue {
			t.Errorf("FromDotEnv() key %s: expected %s, got %s", key, expectedValue, parsedValue)
		}
	}
}

func Test_Stringify(t *testing.T) {
	for tName, tData := range testData {
		t.Run(tName, func(t *testing.T) {
			data := &Data{
				parsed: &tData,
			}
			err := data.StringifyValues()
			if err != nil {
				t.Errorf("StringifyValues() error = %v", err)
				return
			}
			snaps.WithConfig(
				snaps.Filename(tName+".stringify"),
			).MatchSnapshot(t, data)
		})
	}
}

func Test_Stringify_fail(t *testing.T) {
	data := &Data{}
	_, err := data.ToYAML()
	if err != ErrorUnparsedData {
		t.Errorf("Expected error %v, got %v", ErrorUnparsedData, err)
	}
}

func Test_Flatten(t *testing.T) {
	for tName, tData := range testData {
		t.Run(tName, func(t *testing.T) {
			data := &Data{
				parsed: &tData,
			}
			err := data.Flatten(".")
			if err != nil {
				t.Errorf("Flatten() error = %v", err)
				return
			}
			snaps.WithConfig(
				snaps.Filename(tName+".flatten"),
			).MatchSnapshot(t, data)
		})
	}
}

func Test_Flatten_fail(t *testing.T) {
	data := &Data{}
	_, err := data.ToYAML()
	if err != ErrorUnparsedData {
		t.Errorf("Expected error %v, got %v", ErrorUnparsedData, err)
	}
}
