package client

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go-v2/service/scheduler"
	schedulerTypes "github.com/aws/aws-sdk-go-v2/service/scheduler/types"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssmTypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

type S3Client interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	GetObjectAttributes(ctx context.Context, params *s3.GetObjectAttributesInput, optFns ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error)
}

type SecretsManagerClient interface {
	PutSecretValue(ctx context.Context, params *secretsmanager.PutSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.PutSecretValueOutput, error)
}

type SsmClient interface {
	PutParameter(ctx context.Context, params *ssm.PutParameterInput, optFns ...func(*ssm.Options)) (*ssm.PutParameterOutput, error)
	GetParameter(ctx context.Context, params *ssm.GetParameterInput, optFns ...func(*ssm.Options)) (*ssm.GetParameterOutput, error)
}

type SchedulerClient interface {
	CreateSchedule(ctx context.Context, params *scheduler.CreateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.CreateScheduleOutput, error)
	UpdateSchedule(ctx context.Context, params *scheduler.UpdateScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.UpdateScheduleOutput, error)
	DeleteSchedule(ctx context.Context, params *scheduler.DeleteScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.DeleteScheduleOutput, error)
	GetSchedule(ctx context.Context, params *scheduler.GetScheduleInput, optFns ...func(*scheduler.Options)) (*scheduler.GetScheduleOutput, error)
	ListSchedules(ctx context.Context, params *scheduler.ListSchedulesInput, optFns ...func(*scheduler.Options)) (*scheduler.ListSchedulesOutput, error)
}

type AwsClient interface {
	S3GetObject(file SopsS3File) (data []byte, err error)
	S3GetObjectETAG(file SopsS3File) (*string, error)
	SecretsManagerPutSecretValue(sopsHash string, secretArn string, secretContent *[]byte, binary *bool) (data *secretsmanager.PutSecretValueOutput, err error)
	SsmPutParameter(parameterName string, parameterContent *[]byte, keyId string) (response *ssm.PutParameterOutput, err error)
	SsmGetParameter(parameterName string) (*string, error)
	SchedulerCreateOrUpdateSchedule(name string, groupName string, scheduleExpression string, topicArn string, roleArn string, message string) error
	SchedulerDeleteSchedule(name string, groupName string) error
	SchedulerListSchedules(groupName string) ([]string, error)
}

type Client struct {
	ctx            context.Context
	s3             S3Client
	secretsManager SecretsManagerClient
	ssm            SsmClient
	scheduler      SchedulerClient
}

func CreateAwsClients(context context.Context) AwsClient {
	cfg, err := config.LoadDefaultConfig(context, config.WithRetryer(func() aws.Retryer {
		return retry.AddWithMaxAttempts(retry.NewStandard(), 5)
	}))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return &Client{
		ctx:            context,
		s3:             s3.NewFromConfig(cfg),
		secretsManager: secretsmanager.NewFromConfig(cfg),
		ssm:            ssm.NewFromConfig(cfg),
		scheduler:      scheduler.NewFromConfig(cfg),
	}
}

type SopsS3File struct {
	Bucket string `json:"Bucket,omitempty"`
	Key    string `json:"Key,omitempty"`
}

func (c *Client) S3GetObject(file SopsS3File) (data []byte, err error) {
	logger := slog.With("Package", "client", "Function", "S3GetObject")
	logger.Info("Downloading file", "Bucket", file.Bucket, "Key", file.Key)

	resp, err := c.s3.GetObject(c.ctx, &s3.GetObjectInput{
		Bucket: &file.Bucket,
		Key:    &file.Key,
	})
	if err != nil {
		return nil, fmt.Errorf("S3 get object error:\n%v", err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read buffer error:\n%v", err)
	}
	logger.Info("Downloaded file", "Size", buf.Len())
	return buf.Bytes(), nil
}

func (c *Client) S3GetObjectETAG(file SopsS3File) (*string, error) {
	attr, err := c.s3.GetObjectAttributes(c.ctx, &s3.GetObjectAttributesInput{
		Bucket: &file.Bucket,
		Key:    &file.Key,
		ObjectAttributes: []s3Types.ObjectAttributes{
			"ETag",
		},
	})

	if err != nil {
		return nil, fmt.Errorf("error while getting S3 object attributes:\n%v", err)
	}

	if attr.ETag == nil {
		return nil, fmt.Errorf("no ETag checksum found in S3 object:\n%v", err)
	}

	return attr.ETag, nil
}

func (c *Client) SecretsManagerPutSecretValue(sopsHash string, secretArn string, secretStringData *[]byte, binary *bool) (data *secretsmanager.PutSecretValueOutput, err error) {
	if binary != nil && *binary {
		input := &secretsmanager.PutSecretValueInput{
			SecretId:           &secretArn,
			SecretBinary:       *secretStringData,
			ClientRequestToken: &sopsHash,
		}
		return c.secretsManager.PutSecretValue(c.ctx, input)
	}
	input := &secretsmanager.PutSecretValueInput{
		SecretId:           &secretArn,
		SecretString:       aws.String(string(*secretStringData)),
		ClientRequestToken: &sopsHash,
	}
	return c.secretsManager.PutSecretValue(c.ctx, input)
}

func (c *Client) SsmPutParameter(parameterName string, parameterContent *[]byte, keyId string) (response *ssm.PutParameterOutput, err error) {
	parameterContentString := string(*parameterContent)
	input := &ssm.PutParameterInput{
		Name:      &parameterName,
		Value:     &parameterContentString,
		Type:      ssmTypes.ParameterTypeSecureString,
		Overwrite: aws.Bool(true),
		KeyId:     &keyId,
	}
	return c.ssm.PutParameter(c.ctx, input)
}

func (c *Client) SsmGetParameter(parameterName string) (*string, error) {
	logger := slog.With("Package", "client", "Function", "SsmGetParameter")
	logger.Info("Fetching parameter", "Name", parameterName)

	resp, err := c.ssm.GetParameter(c.ctx, &ssm.GetParameterInput{
		Name:           &parameterName,
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		return nil, fmt.Errorf("SSM get parameter error:\n%v", err)
	}
	if resp.Parameter == nil || resp.Parameter.Value == nil {
		return nil, fmt.Errorf("SSM parameter value is nil for parameter: %s", parameterName)
	}
	return resp.Parameter.Value, nil
}

// SchedulerCreateOrUpdateSchedule creates or updates an EventBridge Scheduler
// one-time schedule that publishes a message to an SNS topic.
func (c *Client) SchedulerCreateOrUpdateSchedule(name string, groupName string, scheduleExpression string, topicArn string, roleArn string, message string) error {
	logger := slog.With("Package", "client", "Function", "SchedulerCreateOrUpdateSchedule")
	logger.Info("Upserting schedule", "Name", name, "Group", groupName, "Expression", scheduleExpression)

	target := schedulerTypes.Target{
		Arn:     aws.String(topicArn),
		RoleArn: aws.String(roleArn),
		Input:   aws.String(message),
	}
	flexibleWindow := schedulerTypes.FlexibleTimeWindow{
		Mode: schedulerTypes.FlexibleTimeWindowModeOff,
	}

	// Try update first; if not found, create.
	_, err := c.scheduler.UpdateSchedule(c.ctx, &scheduler.UpdateScheduleInput{
		Name:               aws.String(name),
		GroupName:          aws.String(groupName),
		ScheduleExpression: aws.String(scheduleExpression),
		Target:             &target,
		FlexibleTimeWindow: &flexibleWindow,
	})
	if err == nil {
		logger.Info("Schedule updated", "Name", name)
		return nil
	}

	// Create if not found.
	_, createErr := c.scheduler.CreateSchedule(c.ctx, &scheduler.CreateScheduleInput{
		Name:               aws.String(name),
		GroupName:          aws.String(groupName),
		ScheduleExpression: aws.String(scheduleExpression),
		Target:             &target,
		FlexibleTimeWindow: &flexibleWindow,
		// Delete the schedule after it fires to avoid accumulation of expired schedules.
		ActionAfterCompletion: schedulerTypes.ActionAfterCompletionDelete,
	})
	if createErr != nil {
		return fmt.Errorf("failed to create schedule %s: %v (update error: %v)", name, createErr, err)
	}
	logger.Info("Schedule created", "Name", name)
	return nil
}

// SchedulerDeleteSchedule deletes a single named schedule from the group.
// Returns nil if the schedule does not exist.
func (c *Client) SchedulerDeleteSchedule(name string, groupName string) error {
	logger := slog.With("Package", "client", "Function", "SchedulerDeleteSchedule")
	logger.Info("Deleting schedule", "Name", name, "Group", groupName)

	_, err := c.scheduler.DeleteSchedule(c.ctx, &scheduler.DeleteScheduleInput{
		Name:      aws.String(name),
		GroupName: aws.String(groupName),
	})
	if err != nil {
		return fmt.Errorf("failed to delete schedule %s: %v", name, err)
	}
	return nil
}

// SchedulerListSchedules returns the names of all schedules in the given group.
func (c *Client) SchedulerListSchedules(groupName string) ([]string, error) {
	logger := slog.With("Package", "client", "Function", "SchedulerListSchedules")
	logger.Info("Listing schedules", "Group", groupName)

	var names []string
	var nextToken *string
	for {
		resp, err := c.scheduler.ListSchedules(c.ctx, &scheduler.ListSchedulesInput{
			GroupName: aws.String(groupName),
			NextToken: nextToken,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to list schedules for group %s: %v", groupName, err)
		}
		for _, s := range resp.Schedules {
			if s.Name != nil {
				names = append(names, *s.Name)
			}
		}
		if resp.NextToken == nil {
			break
		}
		nextToken = resp.NextToken
	}
	return names, nil
}
