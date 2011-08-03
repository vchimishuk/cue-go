package cue

import (
	"testing"
)

type expected struct {
	Cmd    string
	Params []string
}

type test struct {
	Input  string
	Etalon expected
}

func TestParseCommand(t *testing.T) {
	var tests = []test{
		{"COMMAND",
			expected{"COMMAND",
				[]string{}}},

		{"COMMAND \t PARAM1   PARAM2\tPARAM3",
			expected{"COMMAND",
				[]string{"PARAM1", "PARAM2", "PARAM3"}}},
		{"COMMAND 'P A R A M 1' \"PA RA M2\" PA\"RAM'3",
			expected{"COMMAND",
				[]string{"P A R A M 1", "PA RA M2", "PARAM3"}}},
	}

	for _, tt := range tests {
		cmd, params, err := parseCommand(tt.Input)
		if err != nil {
			t.Fatalf(err.String())
		}

		if cmd != tt.Etalon.Cmd {
			t.Fatalf("Parsed command '%s' but '%s' expected", cmd, tt.Etalon.Cmd)
		}

		if len(params) != len(tt.Etalon.Params) {
			t.Fatalf("Parsed %d params but %d expected", len(params), len(tt.Etalon.Params))
		}

		for i := 0; i < len(params); i++ {
			if params[i] != tt.Etalon.Params[i] {
				t.Fatalf("Parsed '%s' parameter but '%s' expected", params[i], tt.Etalon.Params[i])
			}
		}
	}
}
