// Package cue implement CUE-SHEET files parser.
// For CUE documentation see: http://digitalx.org/cue-sheet/syntax/
package cue

import (
	"io"
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// commandParser is the function for parsing one command.
type commandParser func(params []string, sheet *CueSheet) os.Error

// commandParserDesctiptor describes command parser.
type commandParserDescriptor struct {
	// -1 -- zero or more parameters.
	paramsCount int
	parser      commandParser
}

// parsersMap used for commands and parser functions correspondence.
var parsersMap = map[string]commandParserDescriptor{
	"CATALOG": {1, parseCatalog},
	//	"CDTEXTFILE": parseCdTextFile,
	//	"FILE":       parseFile,
	//	"FLAGS":      parseFlags,
	//	"INDEX":      parseIndex,
	//	"ISRC":       parseIsrc,
	//	"PERFORMER":  parsePerformer,
	//	"POSTGAP":    parsePostgap,
	//	"PREGAP":     parsePregap,
	"REM": {-1, parseRem},
	//	"SONGWRITER": parseSongWriter,
	//	"TITLE":      parseTitle, 
}

// Parse parses cue-sheet data (file) and returns filled CueSheet struct.
func Parse(reader io.Reader) (sheet *CueSheet, err os.Error) {
	sheet = new(CueSheet)

	rd := bufio.NewReader(reader)
	lineNumber := 0

	for buf, _, err := rd.ReadLine(); err != os.EOF; buf, _, err = rd.ReadLine() {
		if err != nil {
			return nil, err
		}

		cmd, params, err := parseCommand(string(buf))
		if err != nil {
			return nil, fmt.Errorf("Line %d. %s", err.String())
		}

		lineNumber++

		parserDescriptor, ok := parsersMap[cmd]
		if !ok {
			return nil, fmt.Errorf("Line %d. Unknown command '%s'", lineNumber, cmd)
		}

		paramsExpected := parserDescriptor.paramsCount
		paramsRecieved := len(params)
		if paramsExpected != -1 && paramsExpected != paramsRecieved {
			return nil, fmt.Errorf("Line %d. Command %s: recieved %d parameters but %d expected",
				lineNumber, cmd, paramsRecieved, paramsExpected)
		}

		err = parserDescriptor.parser(params, sheet)
		if err != nil {
			return nil, fmt.Errorf("Line %d. Failed to parse %s command. %s", lineNumber, cmd, err.String())
		}
	}

	return sheet, nil
}

// parseCatalog parsers CATALOG command.
func parseCatalog(params []string, sheet *CueSheet) os.Error {
	num := params[0]

	matched, _ := regexp.MatchString("^[0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9][0-9]$", num)
	if !matched {
		return fmt.Errorf("%s is not valid catalog number", params)
	}

	sheet.Catalog = num

	return nil
}

// parseCdTextFile parsers CDTEXTFILE command.
func parseCdTextFile(params []string, sheet *CueSheet) os.Error {
	return nil
}

// parseFile parsers FILE command.
func parseFile(params []string, sheet *CueSheet) os.Error {
	return nil
}

// parseFlags parsers FLAGS command.
func parseFlags(params []string, sheet *CueSheet) os.Error {
	return nil
}

// parseIndex parsers INDEX command.
func parseIndex(params []string, sheet *CueSheet) os.Error {
	return nil
}

// parseIsrc parsers ISRC command.
func parseIsrc(params []string, sheet *CueSheet) os.Error {
	return nil
}

// parsePerformer parsers PERFORMER command.
func parsePerformer(params []string, sheet *CueSheet) os.Error {
	return nil
}

// parsePostgap parsers POSTGAP command.
func parsePostgap(params []string, sheet *CueSheet) os.Error {
	return nil
}

// parsePregap parsers PREGAP command.
func parsePregap(params []string, sheet *CueSheet) os.Error {
	return nil
}

// parseRem parsers REM command.
func parseRem(params []string, sheet *CueSheet) os.Error {
	sheet.Comments = append(sheet.Comments, strings.Join(params, " "))

	return nil
}

// parseSongWriter parsers SONGWRITER command.
func parseSongWriter(params []string, sheet *CueSheet) os.Error {
	return nil
}

// parseTitle parsers TITLE command.
func parseTitle(params []string, sheet *CueSheet) os.Error {
	return nil
}
