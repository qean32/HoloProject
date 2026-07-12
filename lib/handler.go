package lib

import (
	"fmt"
	"main/constants"
	"main/constants/literals"
	"main/lib/array"
	"main/lib/filter"
	"main/lib/low"
	"main/model"
	"main/terminal"
	"main/terminal/list"
	"os"
	"slices"
	"strings"
)

var MAP_HANDLER = map[string]model.EventFunction{
	literals.COMMANDS.CRIPTO: func(e model.Event) {
		low.PushToFile(constants.PATH_DATA, fmt.Sprintf("%s %s", e.KeyWord, e.Payload))
		low.TMP_DATA = append(low.TMP_DATA, []string{e.KeyWord, e.Payload})
	},
	literals.COMMANDS.ECRIPTO: func(e model.Event) {
		index := slices.IndexFunc(low.TMP_DATA, func(item []string) bool {
			return item[0] == e.KeyWord
		})

		if index != -1 {
			terminal.Output(low.TMP_DATA[index][1])
		} else {
			terminal.Output(constants.UNDEFINED_KEYWORD)
		}
	},
	literals.COMMANDS.GKEY: func(e model.Event) {},
	literals.COMMANDS.CLOG: func(e model.Event) { low.ClearFile(constants.PATH_LOG) },
	literals.COMMANDS.GMASTER: func(e model.Event) {
		low.CreateFile(constants.PATH_LOG)
		low.CreateFile(constants.PATH_COMMAND)
		low.CreateFile(constants.PATH_DATA)
	},
	literals.COMMANDS.DROP: func(e model.Event) {
		os.RemoveAll(constants.Root)
		os.Mkdir(constants.Root, 0755)
	},
	literals.COMMANDS.STOP: func(e model.Event) { low.StopProcess() },
	literals.COMMANDS.HELP: func(e model.Event) { terminal.Output(constants.HelpMessage) },
	literals.COMMANDS.DECLARE: func(e model.Event) {
		low.PushToFile(constants.PATH_COMMAND, fmt.Sprintf("%s %s", e.KeyWord, e.Payload))
		low.TMP_COMMANDS = append(low.TMP_COMMANDS, []string{e.KeyWord, e.Payload})
	},
	literals.COMMANDS.RUN: func(e model.Event) {
		index := slices.IndexFunc(low.TMP_COMMANDS, func(item []string) bool {
			return item[0] == strings.TrimSpace(e.KeyWord)
		})

		if index != -1 {
			low.RUN_CMD(low.TMP_COMMANDS[index][1])
		} else {
			terminal.Output(constants.UNDEFINED_KEYWORD)
		}
	},
	literals.COMMANDS.RUNM: func(e model.Event) {
		index := slices.IndexFunc(low.TMP_COMMANDS, func(item []string) bool {
			return item[0] == strings.TrimSpace(e.KeyWord)
		})

		if slices.IndexFunc(low.TMP_COMMANDS, func(item []string) bool {
			return item[0] == strings.TrimSpace(e.KeyWord)
		}) != -1 {
			tmp := strings.Split(low.TMP_COMMANDS[index][1], ";")

			for i := 0; i < len(tmp); i++ {
				low.RUN_CMD(tmp[i])
			}
		} else {
			terminal.Output(constants.UNDEFINED_KEYWORD)
		}
	},
	literals.COMMANDS.COMMANDS: func(e model.Event) {
		terminal.Output(strings.Join(low.ReadFile(constants.PATH_COMMAND), "\n~ "))
	},
	literals.COMMANDS.RMC: func(e model.Event) {
		// need add yes/no
		if slices.IndexFunc(low.TMP_COMMANDS, func(item []string) bool {
			return item[0] == e.KeyWord
		}) != -1 {
			filtered := filter.FILTER(low.TMP_COMMANDS, func(item []string) bool { return item[0] != e.KeyWord })
			low.TMP_COMMANDS = filtered
			low.WriteFile(strings.Join(array.MatrixToArrayString(filtered), "\n"), constants.PATH_COMMAND)
		} else {
			terminal.Output(constants.UNDEFINED_KEYWORD)
		}
	},
	literals.COMMANDS.NOTES: func(e model.Event) {},
	literals.COMMANDS.MENU:  func(e model.Event) { list.List(Menu) },
	literals.COMMANDS.NOTE:  func(e model.Event) {},
	literals.COMMANDS.DNOTE: func(e model.Event) {},
}

var MAP_PARSE = map[string]model.FnReturnEvent{
	literals.COMMANDS.CRIPTO: func(arr []string) (e model.Event, _error bool) {
		payload := getPaylaod(arr)

		if len(arr) < 3 || payload == "" {
			_error = true
			return
		}

		e = model.Event{
			DateTime: low.CurrentTime(),
			Key:      arr[0],
			KeyWord:  arr[1],
			Password: arr[2],
			Payload:  payload,
			Flags:    filter.FilterIsFlag(arr),
		}
		return
	},
	literals.COMMANDS.ECRIPTO: func(arr []string) (e model.Event, _error bool) {
		if len(arr) < 3 {
			_error = true
			return
		}

		e = model.Event{
			DateTime: low.CurrentTime(),
			Key:      arr[0],
			KeyWord:  arr[1],
			Password: arr[2],
			Flags:    filter.FilterIsFlag(arr),
		}
		return
	},
	literals.COMMANDS.DECLARE: func(arr []string) (e model.Event, _error bool) {
		payload := getPaylaod(arr)

		if len(arr) < 3 || payload == "" {
			_error = true
			return
		}

		e = model.Event{
			DateTime: low.CurrentTime(),
			Key:      arr[0],
			KeyWord:  arr[1],
			Payload:  payload,
			Flags:    filter.FilterIsFlag(arr),
		}
		return
	},
	literals.COMMANDS.RUN:   SHORT_EVENT_WordKey,
	literals.COMMANDS.RUNM:  SHORT_EVENT_WordKey,
	literals.COMMANDS.RMC:   SHORT_EVENT_WordKey,
	literals.COMMANDS.NOTES: SHORT_EVENT_WordKey,
	literals.COMMANDS.NOTE:  SHORT_EVENT_WordKey,
	literals.COMMANDS.DNOTE: SHORT_EVENT_WordKey,
	literals.COMMANDS.MENU:  SHORT_EVENT_WordKey,
}

func SHORT_EVENT(arr []string) (e model.Event, _error bool) {
	e = model.Event{
		DateTime: low.CurrentTime(),
		Key:      arr[0],
		Flags:    filter.FilterIsFlag(arr),
	}
	return
}

func SHORT_EVENT_WordKey(arr []string) (e model.Event, _error bool) {
	e = model.Event{
		DateTime: low.CurrentTime(),
		Key:      arr[0],
		KeyWord:  arr[1],
		Flags:    filter.FilterIsFlag(arr),
	}
	return
}

func PARSE_EVENT(command string, key string) (e model.Event, _error bool) {
	fn := MAP_PARSE[key]

	if fn == nil {
		e, _error = SHORT_EVENT(strings.Split(command, " "))
		return
	}
	e, _error = fn(strings.Split(command, " "))
	return e, _error
}

func getPaylaod(arr []string) string {
	var payload string
	command := strings.Join(arr, " ")
	start := strings.IndexAny(command, "{")
	if start != -1 {
		end := strings.IndexAny(command, "}")
		payload = command[start+1 : end]
	}

	return payload
}
