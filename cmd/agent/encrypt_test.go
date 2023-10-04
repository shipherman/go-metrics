package main

import (
	"io/fs"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncrypt(t *testing.T) {
	// Read wrong key path
	_, err := Encrypt("no/such/path", []byte("some data"))

	assert.ErrorIs(t, err, fs.ErrNotExist)

	// Read broken key
	keyPath := "brokenkey"
	f, _ := os.Create(keyPath)
	_, err = f.Write([]byte("brokey key"))
	if err != nil {
		panic(err)
	}
	f.Close()

	_, err = Encrypt(keyPath, []byte("some data"))

	assert.ErrorIs(t, err, ErrBrokenKeyFile)

	// REmove temp file
	os.Remove(keyPath)

}
