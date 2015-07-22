package cue

import (
	"fmt"
	"os"
	"testing"
)

func TestPackage(t *testing.T) {
	filename := "test.cue"

	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Failed to open file. %s", err.Error())
	}

	sheet, err := Parse(file)
	if err != nil {
		t.Fatalf("Failed to parse file. %s", err.Error())
	}

	fmt.Printf("Sheet: %V\n", sheet)
}
