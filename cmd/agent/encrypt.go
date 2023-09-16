package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func Encrypt(publicKeyPath string, data []byte) ([]byte, error) {
	publicKeyPEM, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return nil, err
	}
	publicKeyBlock, _ := pem.Decode(publicKeyPEM)
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		return nil, err

	}

	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), data)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil
}
