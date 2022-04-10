package main

import (
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func Test_FullWorkflow_Delete(t *testing.T) {
	mocks, ctx, event := prepareHandler(t, "events/event_delete.json")

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}
