package data

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Data struct {
	content interface{}
}

func FromYAML(in []byte) (*Data, error) {
	var content interface{}
	err := yaml.Unmarshal(in, &content)
	if err != nil {
		return nil, err
	}
	return &Data{
		content: content,
	}, nil
}

func FromJSON(in []byte) (*Data, error) {
	var content interface{}
	err := json.Unmarshal(in, &content)
	if err != nil {
		return nil, err
	}
	return &Data{
		content: content,
	}, nil
}

func (d *Data) toJSON() ([]byte, error) {
	ret, err := json.MarshalIndent(d.content, "", "  ")
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (d *Data) toYAML() ([]byte, error) {
	ret, err := yaml.Marshal(d.content)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (d *Data) StringifyValues() error {

	stringified, _, err := stringifyValues(d.content)
	if err != nil {
		return err
	}
	d.content = stringified
	return nil
}

func (d *Data) Flatten(separator string) error {
	output := make(map[string]interface{})
	err := flatten("", d.content, output, separator)
	if err != nil {
		return err
	}
	d.content = output
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
