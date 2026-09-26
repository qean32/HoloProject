package death

import (
	"strings"

	"main/model"
)

var (
	TMP_DATA     = [][]string{}
	TMP_COMMANDS = [][]string{}
	SETTINGS     = model.Settings{}
)

func SETDATA() {
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
