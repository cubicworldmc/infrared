package protocol

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha512"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

type Signer struct {
	key *ecdsa.PrivateKey
}

func NewSigner(path string) (*Signer, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("missing block")
	}

	key, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return &Signer{key}, nil
}

func (s *Signer) Sign(payload string) (string, error) {
	hash := sha512.Sum512([]byte(payload))

	signature, err := ecdsa.SignASN1(rand.Reader, s.key, hash[:])
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s///%s", payload, string(signature)), nil
}
