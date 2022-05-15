package adler32

import (
	"testing"
)

func TestAdler32_RollIn(t *testing.T) {
	rolling := New()

	expected := rolling.Write([]byte("Hello World!")).Sum()
	actual := rolling.
		RollIn('H').
		RollIn('e').
		RollIn('l').
		RollIn('l').
		RollIn('o').
		RollIn(' ').
		RollIn('W').
		RollIn('o').
		RollIn('r').
		RollIn('l').
		RollIn('d').
		RollIn('!').
		Sum()

	if expected != actual {
		t.Errorf("expected hash: %d, got %d\n", expected, actual)
	}
}

func TestAdler32_RollOut(t *testing.T) {
	rolling := New()

	expected := rolling.Write([]byte("ello World!")).Sum()
	actual := rolling.
		RollIn('H').
		RollIn('e').
		RollIn('l').
		RollIn('l').
		RollIn('o').
		RollIn(' ').
		RollIn('W').
		RollIn('o').
		RollIn('r').
		RollIn('l').
		RollIn('d').
		RollIn('!').
		RollOut().
		Sum()

	if expected != actual {
		t.Errorf("expected hash: %d, got %d\n", expected, actual)
	}
}
