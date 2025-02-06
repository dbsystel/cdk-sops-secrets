package data

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"gopkg.in/yaml.v3"
)

type Data struct {
	parsed *any
	raw    *[]byte
	hash   *string
}

var (
	ErrorUnparsedData = fmt.Errorf("Data does not contain parsed data")
)

func FromBinary(in []byte, hash *string) (*Data, error) {
	return &Data{
		raw:  &in,
		hash: hash,
	}, nil
}

func FromYAML(in []byte, hash *string) (*Data, error) {
	var content any
	err := yaml.Unmarshal(in, &content)
	if err != nil {
		return nil, err
	}
	return &Data{
		parsed: &content,
		raw:    &in,
		hash:   hash,
	}, nil
}

func FromJSON(in []byte, hash *string) (*Data, error) {
	var content any
	err := json.Unmarshal(in, &content)
	if err != nil {
		return nil, err
	}
	return &Data{
		parsed: &content,
		raw:    &in,
		hash:   hash,
	}, nil
}

func FromDotEnv(in []byte, hash *string) (*Data, error) {
	var dotEnvMap any = map[string]string{}
	dotenvLines := strings.Split(string(in), "\n")
	for _, line := range dotenvLines {
		if line != "" && !strings.HasPrefix(line, "#") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])
				dotEnvMap.(map[string]string)[key] = value
			}
		}
	}
	return &Data{
		parsed: &dotEnvMap,
		raw:    &in,
		hash:   hash,
	}, nil
}

func (d *Data) ToJSON() (*[]byte, error) {
	if d.parsed == nil {
		return nil, ErrorUnparsedData
	}
	content := *d.parsed
	ret, err := json.MarshalIndent(content, "", "  ")
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
				stringMap[k] = []byte(fmt.Sprintf("%v", v))
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
	stringified, err := stringifyValues(*d.parsed)
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

func (d *Data) stripEmptyKey() error {
	if d.parsed == nil {
		return ErrorUnparsedData
	}
	if parsedMap, ok := (*d.parsed).(map[string]interface{}); ok {
		if val, exists := parsedMap[""]; exists && len(parsedMap) == 1 {
			if nestedMap, ok := val.(map[string]interface{}); ok {
				var nestedAny any = nestedMap
				d.parsed = &nestedAny
			}
		}
	}
	return nil
}

// stringifyValues recursively stringifies primitive values in complex types
func stringifyValues(input any) (any, error) {
	// Get the type of the input
	v := reflect.ValueOf(input)

	// If the value is a pointer, dereference it to get the actual value
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// Handle based on the kind of the value
	switch v.Kind() {
	case reflect.Map:
		// If it's a map, iterate through the map and recursively stringify the values
		output := make(map[string]interface{})
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			strVal, err := stringifyValues(val.Interface())
			if err != nil {
				return nil, err
			}
			output[key.String()] = strVal
		}
		return output, nil
	case reflect.Array, reflect.Slice:
		// If it's a slice or array, recursively stringify each element
		output := []interface{}{}
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			strElem, err := stringifyValues(elem.Interface())
			if err != nil {
				return nil, err
			}
			output = append(output, strElem)
		}
		return output, nil
	case reflect.Struct:
		// If it's a struct, iterate through its fields and recursively stringify them
		output := make(map[string]interface{})
		for i := 0; i < v.NumField(); i++ {
			field := v.Field(i)
			strField, err := stringifyValues(field.Interface())
			if err != nil {
				return nil, err
			}
			output[v.Type().Field(i).Name] = strField
		}
		return output, nil
	case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64, reflect.Bool, reflect.Complex64, reflect.Complex128:
		// If it's a primitive type, convert it to a string
		return fmt.Sprintf("%v", input), nil
	default:
		// If it's an unsupported type, return an error
		return nil, fmt.Errorf("unsupported type: %v", v.Kind())
	}
}

func flatten(parentKey string, input any, output map[string]interface{}, separator string) error {
	v := reflect.ValueOf(input)

	// If the value is a pointer, dereference it to get the actual value
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			val := v.MapIndex(key)
			newKey := key.String()
			if parentKey != "" {
				newKey = parentKey + separator + newKey
			}
			if err := flatten(newKey, val.Interface(), output, separator); err != nil {
				return err
			}
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			newKey := fmt.Sprintf("%s[%d]", parentKey, i)
			if err := flatten(newKey, elem.Interface(), output, separator); err != nil {
				return err
			}
		}
	default:
		output[parentKey] = input
	}
	return nil
}
