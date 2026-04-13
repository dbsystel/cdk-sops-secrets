package sops

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/getsops/sops/v3/decrypt"
	"github.com/markussiebert/cdk-sops-secrets/internal/data"
	"gopkg.in/yaml.v3"
)

type Format string

const (
	JSON   Format = "json"
	YAML   Format = "yaml"
	DOTENV Format = "dotenv"
	BINARY Format = "binary"
)

type EncryptedSopsSecret struct {
	Content []byte
	Format  Format
	Hash    string
}

func CreateEncryptedSopsSecret(content []byte, format Format, hash string) (*EncryptedSopsSecret, error) {
	if len(content) == 0 {
		return nil, fmt.Errorf("content cannot be empty")
	}
	if format == "" {
		return nil, fmt.Errorf("format cannot be empty")
	}
	if hash == "" {
		return nil, fmt.Errorf("hash cannot be empty")
	}
	return &EncryptedSopsSecret{
		Content: content,
		Format:  format,
		Hash:    hash,
	}, nil
}

func (e EncryptedSopsSecret) UsesAgeEncryption() bool {
	logger := slog.With("Package", "sops", "Function", "UsesAgeEncryption")

	switch e.Format {
	case JSON, BINARY:
		var meta struct {
			SOPS struct {
				Age []json.RawMessage `json:"age"`
			} `json:"sops"`
		}
		if err := json.Unmarshal(e.Content, &meta); err != nil {
			logger.Warn("could not parse SOPS JSON metadata, assuming age is present", "error", err)
			return true
		}
		return len(meta.SOPS.Age) > 0

	case YAML:
		var meta struct {
			SOPS struct {
				Age []interface{} `yaml:"age"`
			} `yaml:"sops"`
		}
		if err := yaml.Unmarshal(e.Content, &meta); err != nil {
			logger.Warn("could not parse SOPS YAML metadata, assuming age is present", "error", err)
			return true
		}
		return len(meta.SOPS.Age) > 0

	case DOTENV:
		return bytes.Contains(e.Content, []byte("sops_age__list_"))

	default:
		return false
	}
}

type DecryptedSopsSecret struct {
	content []byte
	format  Format
	hash    string
}

func (e EncryptedSopsSecret) Decrypt() (*DecryptedSopsSecret, error) {
	// Create a logger with context information
	logger := slog.With("Package", "sops", "Function", "Decrypt")
	logger.Info("Decrypting content", "Format", e.Format)

	// Decrypt the content using the specified format
	cleartext, err := decrypt.Data(e.Content, string(e.Format))
	if err != nil {
		return nil, fmt.Errorf("decryption error:\n%v", err)
	}
	logger.Info("Decryption successful")

	// If the format is JSON, process the cleartext to ensure proper formatting
	if e.Format == JSON {
		var jsonObj interface{}
		var buf bytes.Buffer

		// Unmarshal the JSON content into an interface
		err := json.Unmarshal(cleartext, &jsonObj)
		if err != nil {
			return nil, fmt.Errorf("decoding error:\n%v", err)
		}

		// Create a JSON encoder with specific settings
		encoder := json.NewEncoder(&buf)
		encoder.SetEscapeHTML(false)
		encoder.SetIndent("", "	") // tab inside

		// Encode the JSON object back into bytes
		err = encoder.Encode(jsonObj)
		if err != nil {
			return nil, fmt.Errorf("encoding error:\n%v", err)
		}

		// Trim any extra whitespace from the encoded JSON
		cleartext = bytes.TrimSpace(buf.Bytes())
	}

	return &DecryptedSopsSecret{
		content: cleartext,
		format:  e.Format,
		hash:    e.Hash,
	}, nil
}

func (d DecryptedSopsSecret) ToData() (*data.Data, error) {
	// Generate a data object by parsing the decrypted secret depending on the data input type
	switch d.format {
	case JSON:
		return data.FromJSON(d.content, &d.hash)
	case YAML:
		return data.FromYAML(d.content, &d.hash)
	case DOTENV:
		return data.FromDotEnv(d.content, &d.hash)
	case BINARY:
		return data.FromBinary(d.content, &d.hash)
	default:
		return nil, fmt.Errorf("unsupported format %s", d.format)
	}
}
