package parse

import (
	"strings"

	"main/constants/literals"
	"main/lib/filter"
	"main/lib/low"
	"main/model"
)

var MAP = map[string]model.FnReturnEvent{
	literals.COMMANDS_LIST.CRIPTO:  ParseCripto,
	literals.COMMANDS_LIST.ECRIPTO: ParseEcripto,
	literals.COMMANDS_LIST.DECLARE: ParseDeclare,

	literals.COMMANDS_LIST.RUN_COMMAND:           ShortEventWithKeyword,
	literals.COMMANDS_LIST.RUN_MULTIPLE_COMMANDS: ShortEventWithKeyword,
	literals.COMMANDS_LIST.REMOVE_COMMAND:        ShortEventWithKeyword,
	literals.COMMANDS_LIST.MENU:                  ShortEvent,
}

func ParseCripto(arr []string) (e model.Event, err bool) {
	payload := getPayload(arr)

	if len(arr) < 3 || payload == "" {
		err = true
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
}

func ParseEcripto(arr []string) (e model.Event, err bool) {
	if len(arr) < 3 {
		err = true
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
}

func ParseDeclare(arr []string) (e model.Event, err bool) {
	payload := getPayload(arr)

	if len(arr) < 3 || payload == "" {
		err = true
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
}

func ShortEvent(arr []string) (e model.Event, err bool) {
	if len(arr) == 0 {
		err = true
		return
	}

	e = model.Event{
		DateTime: low.CurrentTime(),
		Key:      arr[0],
		Flags:    filter.FilterIsFlag(arr),
	}
	return
}

func ShortEventWithKeyword(arr []string) (e model.Event, err bool) {
	if len(arr) < 2 {
		err = true
		return
	}

	e = model.Event{
		DateTime: low.CurrentTime(),
		Key:      arr[0],
		KeyWord:  arr[1],
		Flags:    filter.FilterIsFlag(arr),
	}
	return
}

func ParseEvent(command string, key string) (e model.Event, err bool) {
	fn := MAP[key]

	arr := strings.Split(command, " ")

	if fn == nil {
		return ShortEvent(arr)
	}
	return fn(arr)
}

func getPayload(arr []string) string {
	command := strings.Join(arr, " ")
	start := strings.Index(command, "{")
	if start == -1 {
		return ""
	}

	end := strings.Index(command[start+1:], "}")
	if end == -1 {
		return ""
	}

	return command[start+1 : start+1+end]
}
