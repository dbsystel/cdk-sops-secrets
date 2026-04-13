package sops

import (
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
)

func read(filename string) ([]byte, error) {
	filePath := "../../../test-secrets/" + filename
	existingContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return existingContent, nil
}

func TestUsesAgeEncryption(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		format  Format
		want    bool
	}{
		{
			name:    "JSON with age keys",
			format:  JSON,
			content: []byte(`{"sops":{"age":[{"recipient":"age1xxx","enc":"---BEGIN---"}],"kms":null}}`),
			want:    true,
		},
		{
			name:    "JSON with null age",
			format:  JSON,
			content: []byte(`{"sops":{"age":null,"kms":[{"arn":"arn:aws:kms:..."}]}}`),
			want:    false,
		},
		{
			name:    "JSON with empty age array",
			format:  JSON,
			content: []byte(`{"sops":{"age":[]}}`),
			want:    false,
		},
		{
			name:    "YAML with age keys",
			format:  YAML,
			content: []byte("sops:\n    age:\n        - recipient: age1xxx\n          enc: '---'\n"),
			want:    true,
		},
		{
			name:    "YAML with no age key",
			format:  YAML,
			content: []byte("sops:\n    kms:\n        - arn: arn:aws:kms:...\n"),
			want:    false,
		},
		{
			name:    "YAML with empty age list",
			format:  YAML,
			content: []byte("sops:\n    age: []\n"),
			want:    false,
		},
		{
			name:    "dotenv with age metadata",
			format:  DOTENV,
			content: []byte("KEY=ENC[...]\nsops_age__list_0__map_recipient=age1xxx\n"),
			want:    true,
		},
		{
			name:    "dotenv without age metadata",
			format:  DOTENV,
			content: []byte("KEY=ENC[...]\nsops_kms__list_0__map_arn=arn:aws:kms:...\n"),
			want:    false,
		},
		{
			name:    "binary (JSON envelope) with age keys",
			format:  BINARY,
			content: []byte(`{"data":"ENC[...]","sops":{"age":[{"recipient":"age1xxx","enc":"---"}]}}`),
			want:    true,
		},
		{
			name:    "binary (JSON envelope) without age keys",
			format:  BINARY,
			content: []byte(`{"data":"ENC[...]","sops":{"age":null,"kms":[{"arn":"arn:aws:kms:..."}]}}`),
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			secret := EncryptedSopsSecret{Content: tt.content, Format: tt.format, Hash: "testhash"}
			assert.Equal(t, tt.want, secret.UsesAgeEncryption())
		})
	}
}

func TestUsesAgeEncryption_RealFiles(t *testing.T) {
	cases := []struct {
		file   string
		format Format
	}{
		{"testsecret.sops.json", JSON},
		{"testsecret.sops.yaml", YAML},
		{"testsecret.sops.env", DOTENV},
		{"README.sops.binary", BINARY},
	}
	for _, tc := range cases {
		t.Run(tc.file, func(t *testing.T) {
			content, err := read(tc.file)
			if err != nil {
				t.Fatalf("failed to read %s: %v", tc.file, err)
			}
			secret := EncryptedSopsSecret{Content: content, Format: tc.format, Hash: "testhash"}
			assert.True(t, secret.UsesAgeEncryption(), "expected age encryption to be detected in %s", tc.file)
		})
	}
}

func TestDecrypt(t *testing.T) {
	tests := map[Format]string{
		BINARY: "README.sops.binary",
		DOTENV: "testsecret.sops.env",
		JSON:   "testsecret.sops.json",
		YAML:   "testsecret.sops.yaml",
	}

	t.Setenv("SOPS_AGE_KEY", "AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3")

	for format, file := range tests {
		t.Run(file, func(t *testing.T) {

			content, err := read(file)

			if err != nil {
				t.Fatalf("Failed to read file: %v", err)
			}

			encryptedSecret := EncryptedSopsSecret{
				Content: content,
				Format:  format,
			}

			decryptedSecret, err := encryptedSecret.Decrypt()
			if err != nil {
				t.Fatalf("Failed to decrypt: %v", err)
			}

			snaps.WithConfig(
				snaps.Filename(file),
			).MatchSnapshot(t, string(decryptedSecret.content))
		})
	}
}

func TestToData(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		format  Format
	}{
		{
			name:    "Test JSON ToData",
			content: []byte(`{"key": "value"}`),
			format:  JSON,
		},
		{
			name:    "Test YAML ToData",
			content: []byte(`key: value`),
			format:  YAML,
		},
		{
			name:    "Test DotEnv ToData",
			content: []byte(`KEY=value`),
			format:  DOTENV,
		},
		{
			name:    "Test Binary ToData",
			content: []byte{0x00, 0x01, 0x02, 0x03},
			format:  BINARY,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decryptedSecret := DecryptedSopsSecret{
				content: tt.content,
				format:  tt.format,
			}

			data, err := decryptedSecret.ToData()
			if err != nil {
				t.Fatalf("Failed to convert to data: %v", err)
			}

			snaps.MatchSnapshot(t, data)
		})
	}
}

func TestCreateEncryptedSopsSecret(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		format  Format
		hash    string
		wantErr bool
	}{
		{
			name:    "Valid JSON Secret",
			content: []byte(`{"key": "value"}`),
			format:  JSON,
			hash:    "somehash",
			wantErr: false,
		},
		{
			name:    "Empty Content",
			content: []byte{},
			format:  JSON,
			hash:    "somehash",
			wantErr: true,
		},
		{
			name:    "Empty Format",
			content: []byte(`{"key": "value"}`),
			format:  "",
			hash:    "somehash",
			wantErr: true,
		},
		{
			name:    "Empty Hash",
			content: []byte(`{"key": "value"}`),
			format:  JSON,
			hash:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encryptedSecret, err := CreateEncryptedSopsSecret(tt.content, tt.format, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Fatalf("CreateEncryptedSopsSecret() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				snaps.MatchSnapshot(t, encryptedSecret)
			}
		})
	}
}
