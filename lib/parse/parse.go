package parse

import (
	"strings"

	"main/constants/literals"
	"main/lib/filter"
	"main/lib/low"
	"main/model"
)

var MAP = map[string]model.FnReturnEvent{
	literals.COMMANDLIST.CRYPTO:   parseCripto,
	literals.COMMANDLIST.DECRYPTO: parseEcrypto,
	literals.COMMANDLIST.DECLARE:  parseDeclare,

	literals.COMMANDLIST.RUNCOMMAND:         eventWithKeyword,
	literals.COMMANDLIST.RUNMULTIPLECOMMAND: eventWithKeyword,
	literals.COMMANDLIST.REMOVECOMMAND:      eventWithKeyword,
	literals.COMMANDLIST.MENU:               event,
	literals.COMMANDLIST.LOGS:               event,
}

func parseCripto(arr []string) (model.Event, bool) {
	payload := getPayload(arr)
	if len(arr) < 3 || payload == "" {
		return model.Event{}, true
	}

	return model.Event{
		DateTime: low.CurrentTime(),
		Key:      arr[0],
		KeyWord:  arr[1],
		Password: arr[2],
		Payload:  payload,
		Flags:    filter.FilterIsFlag(arr),
	}, false
}

func parseEcrypto(arr []string) (model.Event, bool) {
	if len(arr) < 3 {
		return model.Event{}, true
	}

	return model.Event{
		DateTime: low.CurrentTime(),
		Key:      arr[0],
		KeyWord:  arr[1],
		Password: arr[2],
		Flags:    filter.FilterIsFlag(arr),
	}, false
}

func parseDeclare(arr []string) (model.Event, bool) {
	payload := getPayload(arr)
	if len(arr) < 3 || payload == "" {
		return model.Event{}, true
	}

	return model.Event{
		DateTime: low.CurrentTime(),
		Key:      arr[0],
		KeyWord:  arr[1],
		Payload:  payload,
		Flags:    filter.FilterIsFlag(arr),
	}, false
}

func event(arr []string) (model.Event, bool) {
	if len(arr) == 0 {
		return model.Event{}, true
	}

	return model.Event{
		DateTime: low.CurrentTime(),
		Key:      arr[0],
		Flags:    filter.FilterIsFlag(arr),
	}, false
}

func eventWithKeyword(arr []string) (model.Event, bool) {
	if len(arr) < 2 {
		return model.Event{}, true
	}

	return model.Event{
		DateTime: low.CurrentTime(),
		Key:      arr[0],
		KeyWord:  arr[1],
		Flags:    filter.FilterIsFlag(arr),
	}, false
}

func ParseEvent(command, key string) (model.Event, bool) {
	arr := strings.Split(command, " ")

	if fn := MAP[key]; fn != nil {
		return fn(arr)
	}
	return event(arr)
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
