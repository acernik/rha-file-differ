package main

import (
	"flag"
	"testing"

	"github.com/acernik/rha-file-differ/internal/adler32"
	"github.com/acernik/rha-file-differ/internal/delta"
	rdr "github.com/acernik/rha-file-differ/internal/reader"
	"github.com/acernik/rha-file-differ/internal/utils"
)

var (
	pathFile1 = ""
	pathFile2 = ""
	chunkSize = 0
)

func TestMain(m *testing.M) {
	pathFile1 = *flag.String("pathFile1", "../../file1.txt", "file1.txt path")
	pathFile2 = *flag.String("pathFile2", "../../file2.txt", "file2.txt path")
	chunkSize = *flag.Int("chunkSize", 4, "chunk size")
	flag.Parse()

	m.Run()
}

func Test_main(t *testing.T) {
	reader := rdr.New()

	file1, file1Size, err := reader.Read(pathFile1)
	if err != nil {
		t.Fatal(err)
	}

	utp := utils.New()

	isAtLeastTwoChunks, err := utp.IsAtLeastTwoChunks(int(file1Size), chunkSize)
	if err != nil {
		t.Fatal(err)
	}

	if !isAtLeastTwoChunks {
		panic("the original file should contain at least two chunks")
	}

	file2, _, err := reader.Read(pathFile2)
	if err != nil {
		t.Fatal(err)
	}

	adler := adler32.New()
	deltaDiffer := delta.New(adler)
	_, err = deltaDiffer.CalculateDelta(file1, file2, chunkSize)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_mainInvalidPathFile1(t *testing.T) {
	reader := rdr.New()

	_, _, err := reader.Read("../file1.txt")
	if err == nil {
		t.Fatal("was expecting error when reading file1.txt, got nil")
	}
}

func Test_mainInvalidChunkSize(t *testing.T) {
	reader := rdr.New()

	_, file1Size, err := reader.Read(pathFile1)
	if err != nil {
		t.Fatal(err)
	}

	utp := utils.New()

	_, err = utp.IsAtLeastTwoChunks(int(file1Size), 0)
	if err == nil {
		t.Fatal("was expecting error when checking if file size is at least two chunks, got nil")
	}
}

func Test_mainLessThenTwoChunks(t *testing.T) {
	reader := rdr.New()

	_, _, err := reader.Read(pathFile1)
	if err != nil {
		t.Fatal(err)
	}

	utp := utils.New()

	isAtLeastTwoChunks, err := utp.IsAtLeastTwoChunks(3, chunkSize)
	if err != nil {
		t.Fatal(err)
	}

	if isAtLeastTwoChunks {
		t.Fatal("was expecting false for isAtLeastTwoChunks when checking if file size is at least two chunks, got true")
	}
}

func Test_mainInvalidPathFile2(t *testing.T) {
	reader := rdr.New()

	_, file1Size, err := reader.Read(pathFile1)
	if err != nil {
		t.Fatal(err)
	}

	utp := utils.New()

	isAtLeastTwoChunks, err := utp.IsAtLeastTwoChunks(int(file1Size), chunkSize)
	if err != nil {
		t.Fatal(err)
	}

	if !isAtLeastTwoChunks {
		panic("the original file should contain at least two chunks")
	}

	_, _, err = reader.Read("../file2.txt")
	if err == nil {
		t.Fatal("was expecting error when reading file2.txt, got nil")
	}
}
