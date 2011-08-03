package cue

import (
	"testing"
	"os"
	"fmt"
)

func TestPackage(t *testing.T) {
	filename := "foo.cue"

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open file. %s", err.String())
	}

	sheet, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse file. %s", err.String())
	}

	fmt.Printf("Sheet: %V\n", sheet)
}
