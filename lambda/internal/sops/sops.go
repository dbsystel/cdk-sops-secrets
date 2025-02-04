package sops

import (
	"fmt"
	"log"

	"github.com/getsops/sops/v3/decrypt"
)

type EncryptedSopsSecret struct {
	content []byte
	format  string
}

type DecryptedSopsSecret struct {
	content []byte
	format  string
}

func (e EncryptedSopsSecret) Decrypt() (*DecryptedSopsSecret, error) {
	log.Printf("Decrypting content with format %s\n", e.format)
	cleartext, err := decrypt.Data(e.content, e.format)
	if err != nil {
		return nil, fmt.Errorf("decryption error:\n%v", err)
	}
	log.Println("Decrypted")
	return &DecryptedSopsSecret{
		content: cleartext,
		format:  e.format,
	}, nil
}
