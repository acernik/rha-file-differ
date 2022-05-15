package reader

import "testing"

func TestReader_Read(t *testing.T) {
	tests := []struct {
		filePath   string
		shouldFail bool
	}{
		{
			filePath:   "../../file1.txt",
			shouldFail: false,
		},

		{
			filePath:   "../../file2.txt",
			shouldFail: false,
		},
		{
			filePath:   "../../file3.txt",
			shouldFail: true,
		},
	}

	testReader := &reader{}

	for index, test := range tests {
		_, _, err := testReader.Read(test.filePath)
		if err != nil && !test.shouldFail {
			t.Fatalf("test index %d: expected err nil, got %s", index, err.Error())
		}
	}
}
