//go:build !integration

package main

import (
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/markussiebert/cdk-sops-secrets/internal/client"
	"github.com/markussiebert/cdk-sops-secrets/internal/event"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockExpirationClient records scheduler calls for assertion.
type mockExpirationClient struct {
	upsertCalls []schedulerUpsertCall
	deleteCalls []schedulerDeleteCall
	listResult  []string
	listErr     error
}

type schedulerUpsertCall struct {
	name               string
	groupName          string
	scheduleExpression string
	topicArn           string
	roleArn            string
	message            string
}

type schedulerDeleteCall struct {
	name      string
	groupName string
}

func (m *mockExpirationClient) S3GetObject(file client.SopsS3File) ([]byte, error) {
	return nil, nil
}
func (m *mockExpirationClient) S3GetObjectETAG(file client.SopsS3File) (*string, error) {
	s := ""
	return &s, nil
}
func (m *mockExpirationClient) SecretsManagerPutSecretValue(sopsHash string, secretArn string, secretContent *[]byte, binary *bool) (*secretsmanager.PutSecretValueOutput, error) {
	return nil, nil
}
func (m *mockExpirationClient) SsmPutParameter(parameterName string, parameterContent *[]byte, keyId string) (*ssm.PutParameterOutput, error) {
	return nil, nil
}
func (m *mockExpirationClient) SsmGetParameter(parameterName string) (*string, error) {
	empty := ""
	return &empty, nil
}
func (m *mockExpirationClient) SchedulerCreateOrUpdateSchedule(name string, groupName string, scheduleExpression string, topicArn string, roleArn string, message string) error {
	m.upsertCalls = append(m.upsertCalls, schedulerUpsertCall{
		name:               name,
		groupName:          groupName,
		scheduleExpression: scheduleExpression,
		topicArn:           topicArn,
		roleArn:            roleArn,
		message:            message,
	})
	return nil
}
func (m *mockExpirationClient) SchedulerDeleteSchedule(name string, groupName string) error {
	m.deleteCalls = append(m.deleteCalls, schedulerDeleteCall{name: name, groupName: groupName})
	return nil
}
func (m *mockExpirationClient) SchedulerListSchedules(groupName string) ([]string, error) {
	return m.listResult, m.listErr
}

// ---- parseExpirationDate tests ----

func TestParseExpirationDate_YYYYMMDD(t *testing.T) {
	t.Parallel()
	ts, err := parseExpirationDate("2030-12-31")
	require.NoError(t, err)
	assert.Equal(t, 2030, ts.Year())
	assert.Equal(t, time.December, ts.Month())
	assert.Equal(t, 31, ts.Day())
}

func TestParseExpirationDate_RFC3339(t *testing.T) {
	t.Parallel()
	ts, err := parseExpirationDate("2030-06-15T12:00:00Z")
	require.NoError(t, err)
	assert.Equal(t, 2030, ts.Year())
	assert.Equal(t, time.June, ts.Month())
}

func TestParseExpirationDate_Invalid(t *testing.T) {
	t.Parallel()
	_, err := parseExpirationDate("not-a-date")
	assert.Error(t, err)
}

func TestParseExpirationDate_WhitespaceStripped(t *testing.T) {
	t.Parallel()
	ts, err := parseExpirationDate("  2030-01-01  ")
	require.NoError(t, err)
	assert.Equal(t, 2030, ts.Year())
}

// ---- sanitizeScheduleName tests ----

func TestSanitizeScheduleName_AllowedChars(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "my-key_name", sanitizeScheduleName("my-key_name"))
}

func TestSanitizeScheduleName_Dots(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "my-key-name", sanitizeScheduleName("my.key.name"))
}

func TestSanitizeScheduleName_Truncation(t *testing.T) {
	t.Parallel()
	long := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789xx" // 66 chars
	result := sanitizeScheduleName(long)
	assert.Len(t, result, 64)
}

// ---- scanExpirationKeys tests ----

func TestScanExpirationKeys_DefaultSuffix(t *testing.T) {
	t.Parallel()
	secretMap := map[string]string{
		"gitlab_token":            "some-token",
		"gitlab_token_expiration": "2099-12-31",
		"other_key":               "value",
	}
	result, err := scanExpirationKeys(secretMap, "_expiration")
	require.NoError(t, err)
	require.Len(t, result, 1)
	assert.Contains(t, result, "gitlab_token")
}

func TestScanExpirationKeys_CustomSuffix(t *testing.T) {
	t.Parallel()
	secretMap := map[string]string{
		"api_key":       "secret",
		"api_key_exp":   "2099-06-01",
		"other_key_exp": "2099-07-01",
	}
	result, err := scanExpirationKeys(secretMap, "_exp")
	require.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestScanExpirationKeys_InvalidDateReturnsError(t *testing.T) {
	t.Parallel()
	secretMap := map[string]string{
		"token_expiration": "not-a-date",
		"key_expiration":   "2099-01-01",
	}
	result, err := scanExpirationKeys(secretMap, "_expiration")
	require.Error(t, err)
	assert.Nil(t, result)
	assert.Contains(t, err.Error(), "token_expiration")
	assert.Contains(t, err.Error(), "not-a-date")
}

func TestScanExpirationKeys_NoMatches(t *testing.T) {
	t.Parallel()
	secretMap := map[string]string{
		"foo": "bar",
		"baz": "qux",
	}
	result, err := scanExpirationKeys(secretMap, "_expiration")
	require.NoError(t, err)
	assert.Empty(t, result)
}

// ---- handleExpirationUpsert tests ----

func TestHandleExpirationUpsert_CreatesSchedules(t *testing.T) {
	t.Parallel()
	mockClient := &mockExpirationClient{}

	suffix := "_expiration"
	days := 14
	props := &event.SopsSyncResourceProperties{
		Expiration: &event.Expiration{
			TopicArn:             "arn:aws:sns:eu-central-1:123456789:my-topic",
			SchedulerRoleArn:     "arn:aws:iam::123456789:role/my-role",
			ScheduleGroupName:    "my-group",
			ExpirationSuffix:     &suffix,
			DaysBeforeExpiration: &days,
		},
	}

	// Use a date far enough in the future so notification time is not in the past.
	futureDate := time.Now().UTC().AddDate(1, 0, 0).Format("2006-01-02")
	secretMap := map[string]string{
		"gitlab_token":            "some-value",
		"gitlab_token_expiration": futureDate,
	}

	err := handleExpirationUpsert(props, mockClient, secretMap, "arn:aws:secretsmanager:eu-central-1:123456789:secret:my-secret")
	require.NoError(t, err)

	require.Len(t, mockClient.upsertCalls, 1)
	call := mockClient.upsertCalls[0]
	assert.Equal(t, "gitlab_token", call.name)
	assert.Equal(t, "my-group", call.groupName)
	assert.Contains(t, call.scheduleExpression, "at(")
	assert.Equal(t, "arn:aws:sns:eu-central-1:123456789:my-topic", call.topicArn)
}

func TestHandleExpirationUpsert_SkipsPastDates(t *testing.T) {
	t.Parallel()
	mockClient := &mockExpirationClient{}

	suffix := "_expiration"
	days := 14
	props := &event.SopsSyncResourceProperties{
		Expiration: &event.Expiration{
			TopicArn:             "arn:aws:sns:eu-central-1:123456789:my-topic",
			SchedulerRoleArn:     "arn:aws:iam::123456789:role/my-role",
			ScheduleGroupName:    "my-group",
			ExpirationSuffix:     &suffix,
			DaysBeforeExpiration: &days,
		},
	}

	// Expiration is only 5 days away but we want 14 days notice, so notification time is in the past.
	soonDate := time.Now().UTC().AddDate(0, 0, 5).Format("2006-01-02")
	secretMap := map[string]string{
		"token_expiration": soonDate,
	}

	err := handleExpirationUpsert(props, mockClient, secretMap, "arn:aws:secretsmanager:eu-central-1:123456789:secret:my-secret")
	require.NoError(t, err)
	assert.Empty(t, mockClient.upsertCalls)
}

func TestHandleExpirationUpsert_DefaultSuffixAndDays(t *testing.T) {
	t.Parallel()
	mockClient := &mockExpirationClient{}

	// nil suffix and days → use defaults
	props := &event.SopsSyncResourceProperties{
		Expiration: &event.Expiration{
			TopicArn:          "arn:aws:sns:eu-central-1:123456789:my-topic",
			SchedulerRoleArn:  "arn:aws:iam::123456789:role/my-role",
			ScheduleGroupName: "my-group",
		},
	}

	futureDate := time.Now().UTC().AddDate(1, 0, 0).Format("2006-01-02")
	secretMap := map[string]string{
		"api_key_expiration": futureDate,
	}

	err := handleExpirationUpsert(props, mockClient, secretMap, "arn:aws:secretsmanager:eu-central-1:123456789:secret:my-secret")
	require.NoError(t, err)
	require.Len(t, mockClient.upsertCalls, 1)
	assert.Equal(t, "api_key", mockClient.upsertCalls[0].name)
}

func TestHandleExpirationUpsert_InvalidDateReturnsError(t *testing.T) {
	t.Parallel()
	mockClient := &mockExpirationClient{}

	props := &event.SopsSyncResourceProperties{
		Expiration: &event.Expiration{
			TopicArn:          "arn:aws:sns:eu-central-1:123456789:my-topic",
			SchedulerRoleArn:  "arn:aws:iam::123456789:role/my-role",
			ScheduleGroupName: "my-group",
		},
	}

	secretMap := map[string]string{
		"api_key_expiration": "not-a-date",
	}

	err := handleExpirationUpsert(props, mockClient, secretMap, "arn:aws:secretsmanager:eu-central-1:123456789:secret:my-secret")
	require.Error(t, err)
	assert.Contains(t, err.Error(), "api_key_expiration")
	assert.Contains(t, err.Error(), "not-a-date")
	assert.Empty(t, mockClient.upsertCalls)
}

// ---- handleExpirationDelete tests ----

func TestHandleExpirationDelete_DeletesAllSchedules(t *testing.T) {
	t.Parallel()
	mockClient := &mockExpirationClient{
		listResult: []string{"sched-a", "sched-b"},
	}

	props := &event.SopsSyncResourceProperties{
		Expiration: &event.Expiration{
			TopicArn:          "arn:aws:sns:eu-central-1:123456789:my-topic",
			SchedulerRoleArn:  "arn:aws:iam::123456789:role/my-role",
			ScheduleGroupName: "my-group",
		},
	}

	err := handleExpirationDelete(props, mockClient)
	require.NoError(t, err)
	assert.Len(t, mockClient.deleteCalls, 2)
}

func TestHandleExpirationDelete_NoSchedules(t *testing.T) {
	t.Parallel()
	mockClient := &mockExpirationClient{
		listResult: []string{},
	}

	props := &event.SopsSyncResourceProperties{
		Expiration: &event.Expiration{
			TopicArn:          "arn:aws:sns:eu-central-1:123456789:my-topic",
			SchedulerRoleArn:  "arn:aws:iam::123456789:role/my-role",
			ScheduleGroupName: "my-group",
		},
	}

	err := handleExpirationDelete(props, mockClient)
	require.NoError(t, err)
	assert.Empty(t, mockClient.deleteCalls)
}
