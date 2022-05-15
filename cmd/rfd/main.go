package main

import (
	"flag"
	"strconv"

	"github.com/acernik/rha-file-differ/internal/adler32"
	"github.com/acernik/rha-file-differ/internal/delta"
	rdr "github.com/acernik/rha-file-differ/internal/reader"
	"github.com/acernik/rha-file-differ/internal/utils"
)

func main() {
	flag.Parse()

	pathFile1 := flag.Arg(0)
	if pathFile1 == "" {
		panic("Please provide a path to the file1.txt as the first argument")
	}

	pathFile2 := flag.Arg(1)
	if pathFile2 == "" {
		panic("Please provide a path to the file2.txt as the second argument")
	}

	chunkSize := 4
	chs := flag.Arg(3)
	if len(chs) > 0 {
		i, err := strconv.Atoi(chs)
		if err != nil {
			panic(err)
		}

		chunkSize = i
	}

	reader := rdr.New()

	file1, file1Size, err := reader.Read(pathFile1)
	if err != nil {
		panic(err)
	}

	utp := utils.New()

	isAtLeastTwoChunks, err := utp.IsAtLeastTwoChunks(int(file1Size), chunkSize)
	if err != nil {
		panic(err)
	}

	if !isAtLeastTwoChunks {
		panic("the original file should contain at least two chunks")
	}

	file2, _, err := reader.Read(pathFile2)
	if err != nil {
		panic(err)
	}

	adler := adler32.New()
	deltaDiffer := delta.New(adler)
	deltaResult, err := deltaDiffer.CalculateDelta(file1, file2, chunkSize)
	if err != nil {
		panic(err)
	}

	deltaResult.DisplayDeltaStats()
}
