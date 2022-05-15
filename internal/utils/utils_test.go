package utils

import "testing"

func TestUtilityProvider_IsAtLeastTwoChunks(t *testing.T) {
	tests := []struct {
		fileSize   int
		chunkSize  int
		expected   bool
		shouldFail bool
	}{
		{
			fileSize:   100,
			chunkSize:  25,
			expected:   true,
			shouldFail: false,
		},
		{
			fileSize:   100,
			chunkSize:  51,
			expected:   false,
			shouldFail: false,
		},
		{
			fileSize:   100,
			chunkSize:  0,
			expected:   false,
			shouldFail: true,
		},
	}

	utp := &utilityProvider{}

	for index, test := range tests {
		result, err := utp.IsAtLeastTwoChunks(test.fileSize, test.chunkSize)
		if err != nil && !test.shouldFail {
			t.Fatalf("test index %d: expected err nil, got %v", index, err.Error())
		}

		if result != test.expected {
			t.Fatalf("test index %d: expected %v, got %v", index, test.expected, result)
		}
	}
}
