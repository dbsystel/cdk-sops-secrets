package event

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
)

func TestGenerateSchema(t *testing.T) {
	GenerateSchema()
}

func createTestSchema(t *testing.T, filename string) cfn.Event {
	file, err := os.Open("testdata/" + filename)
	assert.NoError(t, err)
	defer file.Close()
	var i map[string]interface{}
	err = json.NewDecoder(file).Decode(&i)
	assert.NoError(t, err)
	return cfn.Event{
		ResourceProperties: i,
	}
}

func TestFromCfnEvent(t *testing.T) {
	tests := []string{
		"valid_minimal.json",
		"valid_full.json",
		"invalid_empty.json",
	}

	for _, test := range tests {
		mockEvent := createTestSchema(t, test)
		props, err := FromCfnEvent(mockEvent)
		snaps.WithConfig(snaps.Filename(test)).MatchSnapshot(t, props)
		snaps.WithConfig(snaps.Filename(test+"_err")).MatchSnapshot(t, fmt.Sprintf("%v", err))
	}
}
