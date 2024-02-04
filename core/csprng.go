package core

import (
	"crypto/rand"
)

type CSPRNG struct{}

func Read(block []byte) error {
	_, err := rand.Read(block)
	return err
}
