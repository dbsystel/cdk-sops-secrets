package sops

import (
	"fmt"
	"os"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
)

func read(filename string) ([]byte, error) {
	filePath := "../../../test-secrets/" + filename
	existingContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return existingContent, nil
}

func TestDecrypt(t *testing.T) {
	tests := map[Format]string{
		BINARY: "sopsfile.enc-age.binary",
		DOTENV: "encrypted-best-secret.env",
		JSON:   "sopsfile-complex.enc-age.json",
		YAML:   "sopsfile.enc-age.yaml",
	}

	t.Setenv("SOPS_AGE_KEY", "AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3")

	for format, file := range tests {
		t.Run(file, func(t *testing.T) {

			content, err := read(fmt.Sprintf("%s/%s", format, file))

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
