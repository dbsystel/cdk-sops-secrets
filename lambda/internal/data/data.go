package data

import (
	"encoding/json"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type Data struct {
	parsed *interface{}
	raw    *[]byte
	hash   *string
}

var (
	ErrorUnparsedData = fmt.Errorf("Data does not contain parsed data")
)

func FromBinary(in []byte) (*Data, error) {
	return &Data{
		raw: &in,
	}, nil
}

func FromYAML(in []byte) (*Data, error) {
	var content interface{}
	err := yaml.Unmarshal(in, &content)
	if err != nil {
		return nil, err
	}
	return &Data{
		parsed: &content,
		raw:    &in,
	}, nil
}

func FromJSON(in []byte) (*Data, error) {
	var content interface{}
	err := json.Unmarshal(in, &content)
	if err != nil {
		return nil, err
	}
	return &Data{
		parsed: &content,
		raw:    &in,
	}, nil
}

func FromDotEnv(in []byte) (*Data, error) {
	var dotEnvMap = make(map[string]string)
	dotenvLines := strings.Split(string(in), "\n")
	for _, line := range dotenvLines {
		if line != "" && !strings.HasPrefix(line, "#") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				dotEnvMap[key] = value
			}
		}
	}
	var content interface{} = dotEnvMap
	return &Data{
		parsed: &content,
		raw:    &in,
	}, nil
}

func (d *Data) ToJSON() (*[]byte, error) {
	if d.parsed == nil {
		return nil, ErrorUnparsedData
	}
	ret, err := json.MarshalIndent(d.parsed, "", "  ")
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (d *Data) ToYAML() (*[]byte, error) {
	if d.parsed == nil {
		return nil, ErrorUnparsedData
	}
	ret, err := yaml.Marshal(d.parsed)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}

func (d *Data) ToDotEnv() (*[]byte, error) {
	if d.parsed == nil {
		return nil, ErrorUnparsedData
	}
	stringMap, err := d.ToStringMap()
	if err != nil {
		return nil, err
	}
	var sb strings.Builder
	for k, v := range stringMap {
		sb.WriteString(fmt.Sprintf("%s=%s\n", k, v))
	}
	ret := []byte(sb.String())
	return &ret, nil
}

func (d *Data) ToStringMap() (map[string][]byte, error) {
	if d.parsed == nil {
		return nil, ErrorUnparsedData
	}
	if result, ok := (*d.parsed).(map[string]interface{}); ok {
		stringMap := make(map[string][]byte)
		for k, v := range result {
			switch v.(type) {
			case string, float32, float64, int, int32, int64, uint, uint32, uint64, bool, []byte:
				stringMap[k] = []byte(fmt.Sprintf("%f", v))
			default:
				return nil, fmt.Errorf(`value for key %s is a complex type (map or slice), maybe run "Flatten" first`, k)
			}
		}
		return stringMap, nil
	}
	return nil, fmt.Errorf("parsed data is not a map[string]interface{}")
}

func (d *Data) GetRaw() (*[]byte, error) {
	if d.raw == nil {
		return nil, ErrorUnparsedData
	}
	return d.raw, nil
}

func (d *Data) StringifyValues() error {
	if d.parsed == nil {
		return ErrorUnparsedData
	}
	stringified, _, err := stringifyValues(*d.parsed)
	if err != nil {
		return err
	}
	d.parsed = &stringified
	return nil
}

func (d *Data) Flatten(separator string) error {
	if d.parsed == nil {
		return ErrorUnparsedData
	}
	output := make(map[string]interface{})
	err := flatten("", *d.parsed, output, separator)
	if err != nil {
		return err
	}
	var content interface{} = output
	d.parsed = &content
	return nil
}

func stringifyValues(input any) (interface{}, string, error) {
	switch child := input.(type) {
	case map[string]interface{}:
		{
			output := make(map[string]interface{})
			for k, v := range child {
				object, val, err := stringifyValues(v)
				if err != nil {
					return nil, "", err
				}
				if object != nil {
					output[k] = object
				} else {
					output[k] = val
				}
			}
			return output, "", nil
		}
	case []interface{}:
		{
			output := []interface{}{}
			for _, v := range child {
				object, val, err := stringifyValues(v)
				if err != nil {
					return nil, "", err
				}
				if object != nil {
					output = append(output, object)
				} else {
					output = append(output, val)
				}
			}
			return output, "", nil
		}
	default:
		{
			return nil, fmt.Sprint(input), nil
		}
	}
}

func flatten(parentkey string, input any, output map[string]interface{}, separator string) error {
	switch child := input.(type) {
	case map[string]interface{}:
		{
			for k, v := range child {
				if parentkey == "" {
					flatten(k, v, output, separator)
				} else {
					flatten(fmt.Sprintf("%s%s%s", parentkey, separator, k), v, output, separator)
				}
			}
		}
	case []interface{}:
		{
			for i, v := range child {
				if parentkey == "" {
					flatten(fmt.Sprintf("[%d]", i), v, output, separator)
				} else {
					flatten(fmt.Sprintf("%s[%d]", parentkey, i), v, output, separator)
				}
			}
		}
	default:
		{
			output[parentkey] = input
		}
	}
	return nil
}
