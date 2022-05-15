package reader

import (
	"bufio"
	"os"
)

// Reader interface defines methods used to read data from file.
type Reader interface {
	Read(path string) (*bufio.Reader, int64, error)
}

// reader is the type that implements Reader interface.
type reader struct {
}

// New returns a new value of type that implements Reader interface.
func New() Reader {
	return &reader{}
}

// Read reads out data from a file.
func (r *reader) Read(path string) (*bufio.Reader, int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}

	fileStats, err := f.Stat()
	if err != nil {
		return nil, 0, err
	}

	return bufio.NewReader(f), fileStats.Size(), nil
}
