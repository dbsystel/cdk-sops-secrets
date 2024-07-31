package main

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func Test_FullWorkflow_Create_S3_Parameter_YAML_multi_Simple(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_parameter_yaml_complex.json")

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}

func Test_FullWorkflow_Create_S3_Parameter_YAML_multi_Simple_custom_keys(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_parameter_yaml_complex_custom_keys.json")

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}
