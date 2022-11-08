package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	"golang.org/x/crypto/nacl/secretbox"
)

const (
	keySize   = 32
	nonceSize = 24
)

// Generates and verifies tokens using NaCl SecretBox.
//
// Tokens are URL-safe base64 blobs, with the first `nonceSize`
// bytes containing the randomly-generated nonce used, followed
// by the encrypted payload.
type Auth struct {
	//Pointer so it can be checked for absence with a nil check
	//instead of looping through the key bytes directly (which,
	//while very limited, would leak information about the key
	//in execution timing).
	key *[keySize]byte
}

func allZero(data []byte) bool {
	for _, v := range data {
		if v != 0 {
			return false
		}
	}
	return true
}

func MakeAuth(key [keySize]byte) (*Auth, error) {
	if allZero(key[:]) {
		return nil, fmt.Errorf("All key bytes are zero")
	}
	return &Auth{key: &key}, nil
}

func (a *Auth) checkInit() {
	//key will only be non-nil if created through NewAuth, which
	//also rejects all-0 keys
	if a.key == nil {
		panic("Key not initialized")
	}
}

// Generates a token containing the given data. The data can only be
// read from the token with the secret key used for generating it.
//
// The returned token is safe for use in URLs.
func (a *Auth) GenerateToken(data []byte) string {
	a.checkInit()

	var nonce [nonceSize]byte
	if _, err := io.ReadFull(rand.Reader, nonce[:]); err != nil {
		panic(fmt.Sprintf("Unable to read nonce bytes: %v", err))
	}

	// Allocate properly sized output
	msg := make([]byte, 0, len(nonce)+len(data)+secretbox.Overhead)
	msg = append(msg, nonce[:]...)
	encrypted := secretbox.Seal(msg, data, &nonce, a.key)

	return base64.URLEncoding.EncodeToString(encrypted)
}

// Verifies whether a token is valid, and extracts the data from it
// if so. Invalid tokens (invalid format/length, attempted tampering)
// are detected and rejected.
func (a *Auth) VerifyToken(token string) ([]byte, error) {
	a.checkInit()

	encrypted, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		return nil, fmt.Errorf("Unable to decode token: %v", token)
	}

	if len(encrypted) < nonceSize {
		return nil, fmt.Errorf("Token too short")
	}

	var nonce [nonceSize]byte
	copy(nonce[:], encrypted[:nonceSize])

	decrypted, ok := secretbox.Open(nil, encrypted[nonceSize:], &nonce, a.key)
	if !ok {
		return nil, fmt.Errorf("Unable to decrypt token")
	}

	return decrypted, nil
}
