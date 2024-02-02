package core

import (
	"crypto/rand"
)

type RNG struct{}

func Read(block []byte) error {
	_, err := rand.Read(block)
	return err
}
