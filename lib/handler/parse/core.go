package parse

import (
	"slices"
	"strings"

	errorlist "main/constants/error-list"
	"main/lib/array/filter"
	"main/lib/death"
	"main/model"
)

var (
	payloadPrefixes  = []string{"p/", "-p:"}
	passwordPrefixes = []string{"pass/", "-pass:"}
)

func Parse(args []string, opts model.ParseOptions) (model.Event, error) {
	if len(args) == 0 {
		return model.Event{}, errorlist.ErrEmptyCommand
	}

	event := model.Event{
		Key:      args[0],
		DateTime: death.CurrentTime(),
		Flags:    parseFlags(args, opts),
	}

	if opts.KeyWord {
		if len(args) < 2 {
			return event, errorlist.ErrTooFewArgs
		}
		event.KeyWord = args[1]
	}

	if opts.Password {
		pass, err := getItem(args, passwordPrefixes)
		if err != nil {
			return event, err
		}
		event.Password = pass
	}

	if opts.Payload {
		var stopPrefixes []string
		if opts.Password {
			stopPrefixes = passwordPrefixes
		}
		payload, err := getPayload(args, payloadPrefixes, stopPrefixes)
		if err != nil {
			return event, err
		}
		event.Payload = payload
	}

	if opts.GetSubEvent {
		if len(args) < 2 {
			return event, errorlist.ErrTooFewArgs
		}
		event.SubEvent = &model.Event{
			KeyWord: args[1],
			Key:     opts.SubEvent,
		}
	}

	return event, nil
}

func getItem(args []string, prefixes []string) (string, error) {
	for _, a := range args {
		if v, ok := cutPrefix(a, prefixes); ok {
			return v, nil
		}
	}
	return "", errorlist.ErrNotExist
}

func getPayload(args []string, prefixes, stopPrefixes []string) (string, error) {
	for i, a := range args {
		if v, ok := cutPrefix(a, prefixes); ok {
			rest := make([]string, 0, len(args)-i)
			rest = append(rest, v)
			for _, next := range args[i+1:] {
				if _, stop := cutPrefix(next, stopPrefixes); stop {
					break
				}
				rest = append(rest, next)
			}
			return strings.TrimSpace(strings.Join(rest, " ")), nil
		}
	}
	return "", errorlist.ErrNoPayload
}

func cutPrefix(value string, prefixes []string) (string, bool) {
	for _, prefix := range prefixes {
		if v, ok := strings.CutPrefix(value, prefix); ok {
			return v, true
		}
	}
	return "", false
}

func parseFlags(args []string, opts model.ParseOptions) []string {
	limit := len(args)
	if opts.Payload {
		if idx := indexOfPrefix(args, payloadPrefixes); idx >= 0 {
			limit = idx
		}
	}

	flags := filter.FilterIsFlag(args[:limit])
	return slices.DeleteFunc(flags, func(flag string) bool {
		_, ok := cutPrefix(flag, passwordPrefixes)
		return ok
	})
}

func indexOfPrefix(args, prefixes []string) int {
	for i, arg := range args {
		if _, ok := cutPrefix(arg, prefixes); ok {
			return i
		}
	}
	return -1
}
