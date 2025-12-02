package second_day

import (
	"os"
	"testing"
)

func TestClearInput(t *testing.T) {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		t.Fatal(err)
	}

	ClearInput(string(input))
}
