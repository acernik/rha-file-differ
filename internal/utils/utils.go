package utils

import "fmt"

const (
	minNumChunks = 2
)

// UtilityProvider interface defines utility methods.
type UtilityProvider interface {
	IsAtLeastTwoChunks(fileSize int, chunkSize int) (bool, error)
}

// utilityProvider is the type that implements UtilityProvider interface.
type utilityProvider struct{}

// New returns a new value of type that implements UtilityProvider interface.
func New() UtilityProvider {
	return &utilityProvider{}
}

// IsAtLeastTwoChunks checks if file size is at least two chunks.
func (p *utilityProvider) IsAtLeastTwoChunks(fileSize int, chunkSize int) (bool, error) {
	if chunkSize <= 0 {
		return false, fmt.Errorf("chunkSize invalid: %d", chunkSize)
	}

	return fileSize/chunkSize >= minNumChunks, nil
}
