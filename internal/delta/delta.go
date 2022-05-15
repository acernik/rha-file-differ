package delta

import (
	"bufio"
	"fmt"
	"io"

	"github.com/acernik/rha-file-differ/internal/adler32"
)

// Differ interface defines methods used to calculate the delta between two files.
type Differ interface {
	CalculateDelta(currentFile *bufio.Reader, updatedFile *bufio.Reader, chunkSize int) (Delta, error)
}

// Chunk holds data for a chunk of a file.
type Chunk struct {
	Data   []byte // actual chunk data
	Start  int    // starting position of the chunk
	Offset int    // offset of the chunk
}

// Delta holds slices of data chunks that are the same and that are different compared existing and updated file.
type Delta struct {
	Same []Chunk
	Diff []Chunk
}

// differ is the type that implements Differ interface.
type differ struct {
	adler adler32.Adler32
}

// New returns a new value of type that implements Differ interface.
func New(adler adler32.Adler32) Differ {
	return differ{
		adler: adler,
	}
}

// DisplayDeltaStats prints out the number of matching and different chunks.
func (d Delta) DisplayDeltaStats() {
	fmt.Println("=====  Delta Results =====")
	fmt.Printf("Matching chunks: %d\n", len(d.Same))
	fmt.Printf("Different chunks: %d\n", len(d.Diff))
	fmt.Println("==========================")
}

// getCurrentFileHashes goes through all the chunks for current file and
// puts together hash to chunk map.
func getCurrentFileHashes(currentFile *bufio.Reader, chunkSize int) (map[uint32][]byte, error) {
	result := make(map[uint32][]byte)

	if currentFile == nil {
		return result, fmt.Errorf("current file is nil")
	}

	if chunkSize <= 0 {
		return result, fmt.Errorf("chunkSize is invalid")
	}

	currentAdler := adler32.New()

	index := 0

	for {
		currentFileByte, err := currentFile.ReadByte()
		if err == io.EOF {
			break
		}

		if err != nil {
			return result, err
		}

		currentAdler = currentAdler.RollIn(currentFileByte)
		if currentAdler.Count() < chunkSize {
			continue
		}

		result[currentAdler.Sum()] = currentAdler.Window()

		if currentAdler.Count() > chunkSize {
			currentAdler = currentAdler.RollOut()
		}

		index++
	}

	return result, nil
}

// CalculateDelta calculates difference between current and updated file and constructs value of type Delta
// containing chunks that are different between the current and updated file and chunks that are the same.
func (d differ) CalculateDelta(currentFile *bufio.Reader, updatedFile *bufio.Reader, chunkSize int) (Delta, error) {
	result := Delta{
		Same: make([]Chunk, 0),
		Diff: make([]Chunk, 0),
	}

	if currentFile == nil {
		return result, fmt.Errorf("current file is nil")
	}

	if updatedFile == nil {
		return result, fmt.Errorf("updated file is nil")
	}

	if chunkSize <= 0 {
		return result, fmt.Errorf("chunkSize is invalid")
	}

	currentFileHashes, err := getCurrentFileHashes(currentFile, chunkSize)
	if err != nil {
		return result, err
	}

	updatedAdler := adler32.New()

	index := 0

	for {
		updatedFileByte, err := updatedFile.ReadByte()
		if err == io.EOF {
			break
		}

		if err != nil {
			return result, err
		}

		updatedAdler = updatedAdler.RollIn(updatedFileByte)
		if updatedAdler.Count() < chunkSize {
			continue
		}

		currentFileHashChunkBytes, ok := currentFileHashes[updatedAdler.Sum()]
		if ok {
			result.Same = append(result.Same, Chunk{
				Data:   currentFileHashChunkBytes,
				Start:  index * chunkSize,
				Offset: (index * chunkSize) + chunkSize,
			})
		} else {
			result.Diff = append(result.Diff, Chunk{
				Data:   updatedAdler.Window(),
				Start:  index * chunkSize,
				Offset: (index * chunkSize) + chunkSize,
			})
		}

		if updatedAdler.Count() > chunkSize {
			updatedAdler = updatedAdler.RollOut()
		}

		index++
	}

	return result, nil
}
