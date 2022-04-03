package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/go-test/deep"
)

func TestMain(t *testing.M) {

	log.SetOutput(ioutil.Discard)

	v := t.Run()
	// After all tests have run `go-snaps` can check for not used snapshots
	snaps.Clean()

	os.Exit(v)
}

func Test_GetS3FileContent(t *testing.T) {
	mocks := &AWS{
		secretsmanager: &SecretsManagerMockClient{},
		s3downlaoder:   &S3ManagerMockClient{},
	}
	data, err := mocks.getS3FileContent(SOPSS3File{
		Bucket: "..",
		Key:    "../test-secrets/json/sopsfile.enc-age.json",
		Format: "..",
	})
	check(err)
	snaps.MatchSnapshot(t, strings.ReplaceAll(string(data), "-----", ""))
}

func Test_UpdateSecret(t *testing.T) {
	mocks := &AWS{
		secretsmanager: &SecretsManagerMockClient{},
		s3downlaoder:   &S3ManagerMockClient{},
	}
	inputArn := "arn:${Partition}:secretsmanager:${Region}:${Account}:secret:${SecretId}"
	secretValue := []byte("some-secret-data")

	response, err := mocks.updateSecret(inputArn, secretValue)
	check(err)

	snaps.MatchSnapshot(t, response)
}
func Test_DecryptSopsFileContent(t *testing.T) {

	os.Setenv("SOPS_AGE_KEY", "AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3")

	sopsEncrypted, err := os.ReadFile("../test-secrets/json/sopsfile.enc-age.json")
	check(err)
	sopsDecrypted, err := decryptSopsFileContent(sopsEncrypted, SOPSS3File{Bucket: "", Key: "", Format: "json"})
	check(err)
	sopsExpected, err := os.ReadFile("../test-secrets/json/sopsfile.json")
	check(err)

	sopsDecryptedJ := UnmarshalAny(sopsDecrypted)
	sopsExpectedJ := UnmarshalAny(sopsExpected)

	if diff := deep.Equal(sopsDecryptedJ, sopsExpectedJ); diff != nil {
		t.Error(diff)
	}
}
