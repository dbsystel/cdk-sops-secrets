package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/gkampitakis/go-snaps/snaps"
)

type SecretsManagerMockClient struct {
	secretsmanageriface.SecretsManagerAPI
	t *testing.T
}

type PutSecretValueInputNotSecure struct {

	// A unique identifier for the new version of the secret.
	//
	// If you use the Amazon Web Services CLI or one of the Amazon Web Services
	// SDKs to call this operation, then you can leave this parameter empty. The
	// CLI or SDK generates a random UUID for you and includes it as the value for
	// this parameter in the request.
	//
	// If you generate a raw HTTP request to the Secrets Manager service endpoint,
	// then you must generate a ClientRequestToken and include it in the request.
	//
	// This value helps ensure idempotency. Secrets Manager uses this value to prevent
	// the accidental creation of duplicate versions if there are failures and retries
	// during a rotation. We recommend that you generate a UUID-type (https://wikipedia.org/wiki/Universally_unique_identifier)
	// value to ensure uniqueness of your versions within the specified secret.
	//
	//    * If the ClientRequestToken value isn't already associated with a version
	//    of the secret then a new version of the secret is created.
	//
	//    * If a version with this value already exists and that version's SecretString
	//    or SecretBinary values are the same as those in the request then the request
	//    is ignored. The operation is idempotent.
	//
	//    * If a version with this value already exists and the version of the SecretString
	//    and SecretBinary values are different from those in the request, then
	//    the request fails because you can't modify a secret version. You can only
	//    create new versions to store new secret values.
	//
	// This value becomes the VersionId of the new version.
	ClientRequestToken *string `min:"32" type:"string" idempotencyToken:"true"`

	// The binary data to encrypt and store in the new version of the secret. To
	// use this parameter in the command-line tools, we recommend that you store
	// your binary data in a file and then pass the contents of the file as a parameter.
	//
	// You must include SecretBinary or SecretString, but not both.
	//
	// You can't access this value from the Secrets Manager console.
	//
	// SecretBinary is a sensitive parameter and its value will be
	// replaced with "sensitive" in string returned by PutSecretValueInput's
	// String and GoString methods.
	//
	// SecretBinary is automatically base64 encoded/decoded by the SDK.
	SecretBinary []byte `min:"1" type:"blob"`

	// The ARN or name of the secret to add a new version to.
	//
	// For an ARN, we recommend that you specify a complete ARN rather than a partial
	// ARN. See Finding a secret from a partial ARN (https://docs.aws.amazon.com/secretsmanager/latest/userguide/troubleshoot.html#ARN_secretnamehyphen).
	//
	// If the secret doesn't already exist, use CreateSecret instead.
	//
	// SecretId is a required field
	SecretId *string `min:"1" type:"string" required:"true"`

	// The text to encrypt and store in the new version of the secret.
	//
	// You must include SecretBinary or SecretString, but not both.
	//
	// We recommend you create the secret string as JSON key/value pairs, as shown
	// in the example.
	//
	// SecretString is a sensitive parameter and its value will be
	// replaced with "sensitive" in string returned by PutSecretValueInput's
	// String and GoString methods.
	SecretString *string `min:"1" type:"string"`

	// A list of staging labels to attach to this version of the secret. Secrets
	// Manager uses staging labels to track versions of a secret through the rotation
	// process.
	//
	// If you specify a staging label that's already associated with a different
	// version of the same secret, then Secrets Manager removes the label from the
	// other version and attaches it to this version. If you specify AWSCURRENT,
	// and it is already attached to another version, then Secrets Manager also
	// moves the staging label AWSPREVIOUS to the version that AWSCURRENT was removed
	// from.
	//
	// If you don't include VersionStages, then Secrets Manager automatically moves
	// the staging label AWSCURRENT to this version.
	VersionStages []*string `min:"1" type:"list"`
}

type putSecretValueInputNotSecure secretsmanager.PutSecretValueInput

func (m *SecretsManagerMockClient) PutSecretValue(input *secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error) {
	versionId := fmt.Sprintf("%x", sha256.Sum256([]byte(*input.SecretString)))

	snaps.MatchSnapshot(m.t, ">>>SecretsManagerMockClient.PutSecretValue.Input", putSecretValueInputNotSecure(*input))

	name := fmt.Sprintf("%x", sha256.Sum256([]byte(*input.SecretId)))

	return &secretsmanager.PutSecretValueOutput{
		ARN:           input.SecretId,
		Name:          &name,
		VersionStages: []*string{&name},
		VersionId:     &versionId,
	}, nil
}

type S3ManagerMockClient struct {
	s3manageriface.DownloaderAPI
	t *testing.T
}

func (d S3ManagerMockClient) Download(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (n int64, err error) {
	dat := ReadFile(*input.Key)
	snaps.MatchSnapshot(d.t, ">>>S3MAnagerMockClient.Download.Input", input)
	_, err = w.WriteAt(dat, int64(0))
	check(err)
	return int64(len(dat)), err
}
