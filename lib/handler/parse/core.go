package parse

import (
	"strings"

	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/lib/array/filter"
	"main/lib/death"
	"main/model"
)

const (
	prefixPassword = "-pass:"
	prefixPayload  = "-p:"
)

type Options struct {
	KeyWord     bool
	Payload     bool
	Password    bool
	Flags       bool
	GetSubEvent bool
	SubEvent    string
}

var registry = map[string]Options{
	eventlist.CRYPTO:   {KeyWord: true, Payload: true, Password: true},
	eventlist.DECRYPTO: {KeyWord: true, Password: true},

	eventlist.BAT:       {KeyWord: true, Payload: true},
	eventlist.RUNBAT:    {KeyWord: true},
	eventlist.REMOVEBAT: {GetSubEvent: true, SubEvent: eventlist.PUREREMOVEBAT, KeyWord: true},
}

func Dispatch(key string, args []string) (model.Event, error) {
	return parse(args, registry[key])
}

func parse(args []string, opts Options) (model.Event, error) {
	if len(args) == 0 {
		return model.Event{}, errorlist.ErrNotExist
	}

	event := model.Event{
		Key:      args[0],
		DateTime: death.CurrentTime(),
		Flags:    filter.FilterIsFlag(args),
	}

	if args == nil {
		return event, nil
	}

	if opts.KeyWord {
		if len(args) < 2 {
			return event, errorlist.ErrNotExist
		}
		event.KeyWord = args[1]
	}

	if opts.Password {
		pass, err := getItem(args, prefixPassword)
		if err != nil {
			return event, err
		}
		event.Password = pass
	}

	if opts.Payload {
		payload, err := getPayload(args, prefixPayload)
		if err != nil {
			return event, err
		}
		event.Payload = payload
	}

	if opts.GetSubEvent {
		if len(args) < 2 && opts.KeyWord {
			return event, errorlist.ErrNotExist
		}
		event.SubEvent = &model.Event{
			KeyWord: args[1],
			Key:     opts.SubEvent,
		}
	}

	return event, nil
}

func getItem(args []string, prefix string) (string, error) {
	for _, a := range args {
		if v, ok := strings.CutPrefix(a, prefix); ok {
			return v, nil
		}
	}
	return "", errorlist.ErrNotExist
}

func getPayload(args []string, prefix string) (string, error) {
	for i, a := range args {
		if v, ok := strings.CutPrefix(a, prefix); ok {
			rest := append([]string{v}, args[i+1:]...)
			return strings.TrimSpace(strings.Join(rest, " ")), nil
		}
	}
	return "", errorlist.ErrNoPayload
}
