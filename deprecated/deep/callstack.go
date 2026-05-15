package deep

import (
	"main/constants"
	"main/model"
	"strings"
)

var CALLSTACK = []model.Event{}
var TMP_DATA = [][]string{}
var TMP_COMMANDS = [][]string{}
var SETTINGS = model.Settings{}

func DATA() {
	if len(TMP_COMMANDS) == 0 {
		strs := (ReadFile(constants.PATH_COMMAND))
		var commands [][]string

		for i := 0; i < len(strs); i++ {
			tmp := strings.Split(strs[i], " ")
			commands = append(commands, []string{tmp[0], strings.Join(tmp[1:], " ")})
		}
		TMP_COMMANDS = commands
	}
	if len(TMP_DATA) == 0 {
		strs := (ReadFile(constants.PATH_DATA))
		var data [][]string

		for i := 0; i < len(strs); i++ {
			tmp := strings.Split(strs[i], " ")
			data = append(data, []string{tmp[0], strings.Join(tmp[1:], " ")})
		}
		TMP_DATA = data
	}
}
