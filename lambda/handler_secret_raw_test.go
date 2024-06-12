package main

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func Test_FullWorkflow_Create_S3_RAW_Simple(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_create_s3_secret_raw_simple.json")

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}
