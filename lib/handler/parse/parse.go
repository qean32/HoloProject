package parse

import (
	"strings"

	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/constants/literal"
	"main/lib/array/filter"
	"main/lib/death"
	"main/model"
)

var MAP = map[string]model.ParseEvent{
	eventlist.CRYPTO:   parseCripto,
	eventlist.DECRYPTO: parseEcrypto,
	eventlist.BAT:      parseDeclare,

	eventlist.RUNBAT:    eventWithKeyword,
	eventlist.REMOVEBAT: eventWithKeywordRemoveBat,
	eventlist.MENU:      event,
	eventlist.OPENLOG:   event,
}

func parseCripto(arr []string) (model.Event, error) {
	if len(arr) < 3 {
		return model.Event{}, errorlist.WithContext("crypto", errorlist.ErrTooFewArgs)
	}
	payload := getPayload(arr)
	if payload == "" {
		return model.Event{}, errorlist.WithContext("crypto", errorlist.ErrNoPayload)
	}

	return model.Event{
		DateTime: death.CurrentTime(),
		Key:      arr[0],
		KeyWord:  arr[1],
		Password: arr[2],
		Payload:  payload,
		Flags:    filter.FilterIsFlag(arr),
	}, nil
}

func parseEcrypto(arr []string) (model.Event, error) {
	if len(arr) < 3 {
		return model.Event{}, errorlist.WithContext("decrypt", errorlist.ErrTooFewArgs)
	}

	return model.Event{
		DateTime: death.CurrentTime(),
		Key:      arr[0],
		KeyWord:  arr[1],
		Password: arr[2],
		Flags:    filter.FilterIsFlag(arr),
	}, nil
}

func parseDeclare(arr []string) (model.Event, error) {
	if len(arr) < 3 {
		return model.Event{}, errorlist.WithContext("bat", errorlist.ErrTooFewArgs)
	}
	payload := getPayload(arr)
	if payload == "" {
		return model.Event{}, errorlist.WithContext("bat", errorlist.ErrNoPayload)
	}

	return model.Event{
		DateTime: death.CurrentTime(),
		Key:      arr[0],
		KeyWord:  arr[1],
		Payload:  payload,
		Flags:    filter.FilterIsFlag(arr),
	}, nil
}

func event(arr []string) (model.Event, error) {
	if len(arr) == 0 {
		return model.Event{}, errorlist.ErrEmptyCommand
	}

	return model.Event{
		DateTime: death.CurrentTime(),
		Key:      arr[0],
		Flags:    filter.FilterIsFlag(arr),
	}, nil
}

func eventWithKeyword(arr []string) (model.Event, error) {
	if len(arr) < 2 {
		return model.Event{}, errorlist.ErrTooFewArgs
	}

	return model.Event{
		DateTime: death.CurrentTime(),
		Key:      arr[0],
		KeyWord:  arr[1],
		Flags:    filter.FilterIsFlag(arr),
	}, nil
}

func eventWithKeywordRemoveBat(arr []string) (model.Event, error) {
	if len(arr) < 2 {
		return model.Event{}, errorlist.ErrTooFewArgs
	}

	return model.Event{
		DateTime: death.CurrentTime(),
		Key:      arr[0],
		KeyWord:  arr[1],
		SubEvent: &model.Event{
			Key:     eventlist.PURE_REMOVE_BAT,
			KeyWord: arr[1] + literal.Extension.Bat,
		},
		Flags: filter.FilterIsFlag(arr),
	}, nil
}

func ParseEvent(command, key string) (model.Event, error) {
	arr := strings.Split(command, " ")

	if fn, ok := MAP[key]; ok {
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
