package sops

import (
	"fmt"
	"log"

	"github.com/getsops/sops/v3/decrypt"
	"github.com/markussiebert/cdk-sops-secrets/internal/data"
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

type DecryptedSopsSecret struct {
	content []byte
	format  Format
	hash    string
}

func (e EncryptedSopsSecret) Decrypt() (*DecryptedSopsSecret, error) {
	log.Printf("Decrypting content with format %s\n", e.Format)
	cleartext, err := decrypt.Data(e.Content, string(e.Format))
	if err != nil {
		return nil, fmt.Errorf("decryption error:\n%v", err)
	}
	log.Println("Decrypted")
	return &DecryptedSopsSecret{
		content: cleartext,
		format:  e.Format,
		hash:    e.Hash,
	}, nil
}

func (d DecryptedSopsSecret) ToData() (*data.Data, error) {
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
