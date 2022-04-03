package main

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func Test_FromYAML(t *testing.T) {
	simpleStruct, err := fromYAML(ReadFile("../test-secrets/yaml/sopsfile.yml"))
	check(err)
	snaps.MatchSnapshot(t, simpleStruct)
	complexStruct, err := fromYAML(ReadFile("../test-secrets/yaml/sopsfile-complex.yml"))
	check(err)
	snaps.MatchSnapshot(t, complexStruct)
}

func Test_toJSON(t *testing.T) {
	simpleJSON, err := toJSON(UnmarshalAny(ReadFile("../test-secrets/json/sopsfile.json")))
	check(err)
	snaps.MatchSnapshot(t, string(simpleJSON))
	complexJSON, err := toJSON(UnmarshalAny(ReadFile("../test-secrets/json/sopsfile-complex.json")))
	check(err)
	snaps.MatchSnapshot(t, string(complexJSON))
}
