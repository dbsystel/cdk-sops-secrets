package main

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func Test_FromJSON(t *testing.T) {
	simpleStruct, err := fromJSON(ReadFile("../test-secrets/json/sopsfile.json"))
	check(err)
	snaps.MatchSnapshot(t, ">>>Simple", simpleStruct)
	complexStruct, err := fromJSON(ReadFile("../test-secrets/json/sopsfile-complex.json"))
	check(err)
	snaps.MatchSnapshot(t, ">>>Complex", complexStruct)
}

func Test_FromYAML(t *testing.T) {
	simpleStruct, err := fromYAML(ReadFile("../test-secrets/yaml/sopsfile.yml"))
	check(err)
	snaps.MatchSnapshot(t, ">>>Simple", simpleStruct)
	complexStruct, err := fromYAML(ReadFile("../test-secrets/yaml/sopsfile-complex.yml"))
	check(err)
	snaps.MatchSnapshot(t, ">>>Complex", complexStruct)
}

func Test_Flatten(t *testing.T) {
	simpleStruct, err := fromJSON(ReadFile("../test-secrets/json/sopsfile.json"))
	check(err)
	flattenedSimpleStruct := make(map[string]interface{})
	err = flatten("", simpleStruct, flattenedSimpleStruct, ".")
	snaps.MatchSnapshot(t, ">>>Simple", flattenedSimpleStruct)
	complexStruct, err := fromJSON(ReadFile("../test-secrets/json/sopsfile-complex.json"))
	check(err)
	flattenedComplexStruct := make(map[string]interface{})
	err = flatten("", complexStruct, flattenedComplexStruct, ".")
	snaps.MatchSnapshot(t, ">>>Complex", flattenedComplexStruct)
}

func Test_toJSON(t *testing.T) {
	simpleJSON, err := toJSON(UnmarshalAny(ReadFile("../test-secrets/json/sopsfile.json")))
	check(err)
	snaps.MatchSnapshot(t, ">>>Simple", string(simpleJSON))
	complexJSON, err := toJSON(UnmarshalAny(ReadFile("../test-secrets/json/sopsfile-complex.json")))
	check(err)
	snaps.MatchSnapshot(t, ">>>Complex", string(complexJSON))
}

func Test_toYAML(t *testing.T) {
	simpleJSON, err := toYAML(UnmarshalAny(ReadFile("../test-secrets/json/sopsfile.json")))
	check(err)
	snaps.MatchSnapshot(t, ">>>Simple", string(simpleJSON))
	complexJSON, err := toYAML(UnmarshalAny(ReadFile("../test-secrets/json/sopsfile-complex.json")))
	check(err)
	snaps.MatchSnapshot(t, ">>>Complex", string(complexJSON))
}
