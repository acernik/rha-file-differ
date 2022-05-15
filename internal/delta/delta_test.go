package delta

import (
	"testing"

	"github.com/acernik/rha-file-differ/internal/adler32"
	"github.com/acernik/rha-file-differ/internal/reader"
)

func TestDiffer_CalculateDelta(t *testing.T) {
	chuckSize := 4

	rdr := reader.New()

	file1, _, err := rdr.Read("../../file1.txt")
	if err != nil {
		t.Fatal(err)
	}

	file2, _, err := rdr.Read("../../file2.txt")
	if err != nil {
		t.Fatal(err)
	}

	adler := adler32.New()
	deltaDiffer := New(adler)
	_, err = deltaDiffer.CalculateDelta(file1, file2, chuckSize)
	if err != nil {
		t.Fatal(err)
	}
}

func TestDiffer_CalculateDeltaInvalidCurrentFile(t *testing.T) {
	chuckSize := 4

	rdr := reader.New()

	file2, _, err := rdr.Read("../../file2.txt")
	if err != nil {
		t.Fatal(err)
	}

	adler := adler32.New()
	deltaDiffer := New(adler)
	_, err = deltaDiffer.CalculateDelta(nil, file2, chuckSize)
	if err == nil {
		t.Fatal("current file is nil, error should have been returned")
	}
}

func TestDiffer_CalculateDeltaInvalidUpdatedFile(t *testing.T) {
	chuckSize := 4

	rdr := reader.New()

	file1, _, err := rdr.Read("../../file1.txt")
	if err != nil {
		t.Fatal(err)
	}

	adler := adler32.New()
	deltaDiffer := New(adler)
	_, err = deltaDiffer.CalculateDelta(file1, nil, chuckSize)
	if err == nil {
		t.Fatal("updated file is nil, error should have been returned")
	}
}

func TestDiffer_CalculateDeltaInvalidChunkSize(t *testing.T) {
	chuckSize := 0

	rdr := reader.New()

	file1, _, err := rdr.Read("../../file1.txt")
	if err != nil {
		t.Fatal(err)
	}

	file2, _, err := rdr.Read("../../file2.txt")
	if err != nil {
		t.Fatal(err)
	}

	adler := adler32.New()
	deltaDiffer := New(adler)
	_, err = deltaDiffer.CalculateDelta(file1, file2, chuckSize)
	if err == nil {
		t.Fatal("chuckSize is 0, error should have been returned")
	}
}

func Test_getCurrentFileHashes(t *testing.T) {
	chuckSize := 4

	rdr := reader.New()

	file1, _, err := rdr.Read("../../file1.txt")
	if err != nil {
		t.Fatal(err)
	}

	_, err = getCurrentFileHashes(file1, chuckSize)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_getCurrentFileHashesInvalidCurrentFile(t *testing.T) {
	chuckSize := 4

	_, err := getCurrentFileHashes(nil, chuckSize)
	if err == nil {
		t.Fatal("current file is nil, error should have been returned")
	}
}

func Test_getCurrentFileHashesInvalidChunkSize(t *testing.T) {
	chuckSize := 0

	rdr := reader.New()

	file1, _, err := rdr.Read("../../file1.txt")
	if err != nil {
		t.Fatal(err)
	}

	_, err = getCurrentFileHashes(file1, chuckSize)
	if err == nil {
		t.Fatal("chuckSize is 0, error should have been returned")
	}
}
