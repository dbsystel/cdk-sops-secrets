package main

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func Test_FullWorkflow_Create_S3_ENV_Simple(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_env_simple.json")

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}

func Test_FullWorkflow_Create_S3_ENV_as_JSON_Simple(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_env_as_json_simple.json")
	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}

//func Test_FullWorkflow_Create_S3_YAML_Complex(t *testing.T) {
//	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_yaml_complex.json")
//	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
//	check(err)
//	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
//}
//func Test_FullWorkflow_Create_S3_YAML_as_JSON_Complex(t *testing.T) {
//	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_yaml_as_json_complex.json")
//	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
//	check(err)
//	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
//}
//func Test_FullWorkflow_Create_S3_YAML_Complex_Flat(t *testing.T) {
//	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_yaml_complex_flat.json")
//	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
//	check(err)
//	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
//}
//func Test_FullWorkflow_Create_S3_YAML_as_JSON_Complex_Flat(t *testing.T) {
//	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_yaml_as_json_complex_flat.json")
//	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
//	check(err)
//	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
//}
//func Test_FullWorkflow_Create_INLINE_YAML_Simple(t *testing.T) {
//	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_yaml_simple.json")
//	event = fileToInline(event)
//	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
//	check(err)
//	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
//}
//func Test_FullWorkflow_Create_INLINE_YAML_as_JSON_Simple(t *testing.T) {
//	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_yaml_as_json_simple.json")
//	event = fileToInline(event)
//	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
//	check(err)
//	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
//}
//func Test_FullWorkflow_Create_INLINE_YAML_Complex(t *testing.T) {
//	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_yaml_complex.json")
//	event = fileToInline(event)
//	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
//	check(err)
//	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
//}
//func Test_FullWorkflow_Create_INLINE_YAML_as_JSON_Complex(t *testing.T) {
//	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_yaml_as_json_complex.json")
//	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
//	check(err)
//	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
//}
//func Test_FullWorkflow_Create_INLINE_YAML_Complex_Flat(t *testing.T) {
//	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_yaml_complex_flat.json")
//	event = fileToInline(event)
//	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
//	check(err)
//	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
//}
//func Test_FullWorkflow_Create_INLINE_YAML_as_JSON_Complex_Flat(t *testing.T) {
//	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_yaml_as_json_complex_flat.json")
//	event = fileToInline(event)
//	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
//	check(err)
//	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
//}
