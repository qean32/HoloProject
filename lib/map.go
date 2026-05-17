package lib

import (
	"fmt"
	"main/constants"
	"main/deep"
	"main/lib/filter"
	"main/model"
	"os"
	"slices"
	"strings"
)

var MAP_HANDLER = map[string]model.EventFunction{
	constants.COMMANDS.CRIPTO: func(e model.Event) {
		deep.PushToFile(constants.PATH_DATA, fmt.Sprintf("%s %s", e.KeyWord, e.Payload))
		deep.TMP_DATA = append(deep.TMP_DATA, []string{e.KeyWord, e.Payload})
	},
	constants.COMMANDS.ECRIPTO: func(e model.Event) {
		index := slices.IndexFunc(deep.TMP_DATA, func(item []string) bool {
			return item[0] == e.KeyWord
		})

		if index != -1 {
			deep.CONSOLE_RESPONSE(deep.TMP_DATA[index][1], false)
		} else {
			deep.CONSOLE_RESPONSE(constants.UNDEFINED_KEYWORD, false)
		}
	},
	constants.COMMANDS.GKEY: func(e model.Event) {},
	constants.COMMANDS.CLOG: func(e model.Event) { deep.ClearFile(constants.PATH_LOG) },
	constants.COMMANDS.GMASTER: func(e model.Event) {
		deep.CreateFile(constants.PATH_LOG)
		deep.CreateFile(constants.PATH_COMMAND)
		deep.CreateFile(constants.PATH_DATA)
	},
	constants.COMMANDS.DROP: func(e model.Event) {
		os.RemoveAll(constants.Root)
		os.Mkdir(constants.Root, 0755)
	},
	constants.COMMANDS.STOP: func(e model.Event) { os.Exit(0) },
	constants.COMMANDS.HELP: func(e model.Event) { deep.CONSOLE_RESPONSE(constants.HelpMessage, false) },
	constants.COMMANDS.DECLARE: func(e model.Event) {
		deep.PushToFile(constants.PATH_COMMAND, fmt.Sprintf("%s %s", e.KeyWord, e.Payload))
		deep.TMP_COMMANDS = append(deep.TMP_COMMANDS, []string{e.KeyWord, e.Payload})
	},
	constants.COMMANDS.RUN: func(e model.Event) {
		index := slices.IndexFunc(deep.TMP_COMMANDS, func(item []string) bool {
			return item[0] == strings.TrimSpace(e.KeyWord)
		})

		if index != -1 {
			deep.RUN_CMD(deep.TMP_COMMANDS[index][1])
		} else {
			deep.CONSOLE_RESPONSE(constants.UNDEFINED_KEYWORD, false)
		}
	},
	constants.COMMANDS.RUNM: func(e model.Event) {
		index := slices.IndexFunc(deep.TMP_COMMANDS, func(item []string) bool {
			return item[0] == strings.TrimSpace(e.KeyWord)
		})

		if slices.IndexFunc(deep.TMP_COMMANDS, func(item []string) bool {
			return item[0] == strings.TrimSpace(e.KeyWord)
		}) != -1 {
			tmp := strings.Split(deep.TMP_COMMANDS[index][1], ";")

			for i := 0; i < len(tmp); i++ {
				deep.RUN_CMD(tmp[i])
			}
		} else {
			deep.CONSOLE_RESPONSE(constants.UNDEFINED_KEYWORD, false)
		}
	},
	constants.COMMANDS.COMMANDS: func(e model.Event) {
		fmt.Println(constants.OUTPUT_MESSAGE)
		deep.CONSOLE(strings.Join(deep.ReadFile(constants.PATH_COMMAND), "\n~ "))
	},
	constants.COMMANDS.RMC: deep.DECORATOR_ACCESS_ACTION(
		func(e model.Event) {
			if slices.IndexFunc(deep.TMP_COMMANDS, func(item []string) bool {
				return item[0] == e.KeyWord
			}) != -1 {
				filtered := filter.FILTER(deep.TMP_COMMANDS, func(item []string) bool { return item[0] != e.KeyWord })
				deep.TMP_COMMANDS = filtered
				deep.WriteFile(strings.Join(matrixToArrayString(filtered), "\n"), constants.PATH_COMMAND)
			} else {
				deep.CONSOLE_RESPONSE(constants.UNDEFINED_KEYWORD, false)
			}
		},
	),
	constants.COMMANDS.NOTES: func(e model.Event) {},
	constants.COMMANDS.NOTE:  func(e model.Event) {},
	constants.COMMANDS.DNOTE: func(e model.Event) {},
}

var MAP_PARSE = map[string]model.FnReturnEvent{
	constants.COMMANDS.CRIPTO: func(arr []string) (e model.Event, _error bool) {
		payload := getPaylaod(arr)

		if len(arr) < 3 || payload == "" {
			_error = true
			return
		}

		e = model.Event{
			DateTime: deep.CurrentTime(),
			Key:      arr[0],
			KeyWord:  arr[1],
			Password: arr[2],
			Payload:  payload,
			Flags:    filter.FilterIsFlag(arr),
		}
		return
	},
	constants.COMMANDS.ECRIPTO: func(arr []string) (e model.Event, _error bool) {
		if len(arr) < 3 {
			_error = true
			return
		}

		e = model.Event{
			DateTime: deep.CurrentTime(),
			Key:      arr[0],
			KeyWord:  arr[1],
			Password: arr[2],
			Flags:    filter.FilterIsFlag(arr),
		}
		return
	},
	constants.COMMANDS.DECLARE: func(arr []string) (e model.Event, _error bool) {
		payload := getPaylaod(arr)

		if len(arr) < 3 || payload == "" {
			_error = true
			return
		}

		e = model.Event{
			DateTime: deep.CurrentTime(),
			Key:      arr[0],
			KeyWord:  arr[1],
			Payload:  payload,
			Flags:    filter.FilterIsFlag(arr),
		}
		return
	},
	constants.COMMANDS.RUN:   SHORT_EVENT_WordKey,
	constants.COMMANDS.RUNM:  SHORT_EVENT_WordKey,
	constants.COMMANDS.RMC:   SHORT_EVENT_WordKey,
	constants.COMMANDS.NOTES: SHORT_EVENT_WordKey,
	constants.COMMANDS.NOTE:  SHORT_EVENT_WordKey,
	constants.COMMANDS.DNOTE: SHORT_EVENT_WordKey,
}

func SHORT_EVENT(arr []string) (e model.Event, _error bool) {
	e = model.Event{
		DateTime: deep.CurrentTime(),
		Key:      arr[0],
		Flags:    filter.FilterIsFlag(arr),
	}
	return
}

func SHORT_EVENT_WordKey(arr []string) (e model.Event, _error bool) {
	e = model.Event{
		DateTime: deep.CurrentTime(),
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
