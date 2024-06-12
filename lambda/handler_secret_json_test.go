package main

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func Test_FullWorkflow_Create_S3_JSON_Simple(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_secret_json_simple.json")

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}

func Test_FullWorkflow_Create_S3_JSON_Complex(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_secret_json_complex.json")

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}

func Test_FullWorkflow_Create_S3_JSON_Complex_StringifyValues(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_secret_json_complex_stringify.json")

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}
func Test_FullWorkflow_Create_S3_JSON_Complex_Flat(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_secret_json_complex_flat.json")

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}

func Test_FullWorkflow_Create_INLINE_JSON_Simple(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_secret_json_simple.json")
	event = fileToInline(event)

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}

func Test_FullWorkflow_Create_INLINE_JSON_Complex(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_secret_json_complex.json")
	event = fileToInline(event)

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}

func Test_FullWorkflow_Create_INLINE_JSON_Complex_StringifyValues(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_secret_json_complex_stringify.json")
	event = fileToInline(event)

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}
func Test_FullWorkflow_Create_INLINE_JSON_Complex_Flat(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_secret_json_complex_flat.json")
	event = fileToInline(event)

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}
