package low

import (
	"strings"

	"main/constants"
	"main/model"
)

var (
	CALLSTACK    = []model.Event{}
	TMP_DATA     = [][]string{}
	TMP_COMMANDS = [][]string{}
	SETTINGS     = model.Settings{}
)

func SETDATA() {
	if len(TMP_COMMANDS) == 0 {
		TMP_COMMANDS = parsePairs(ReadFile(constants.PATH_COMMAND))
	}
	if len(TMP_DATA) == 0 {
		TMP_DATA = parsePairs(ReadFile(constants.PATH_DATA))
	}
}

func parsePairs(lines []string) [][]string {
	result := make([][]string, 0, len(lines))
	for _, line := range lines {
		tmp := strings.Split(line, " ")
		if len(tmp) == 0 {
			continue
		}
		result = append(result, []string{tmp[0], strings.Join(tmp[1:], " ")})
	}
	return result
}
