package lib

import (
	"main/model"
	"os"
	"strings"
	"time"
)

var KEY_FUNCTION = map[string]func(e model.Event){
	"hash":   func(e model.Event) {},             // hash ключ пароль {сообщение} (-nl)
	"g:key":  func(e model.Event) {},             // g:key
	"g:log":  func(e model.Event) {},             // g:log
	"c:log":  func(e model.Event) {},             // c:log (-nl -f)
	"dihash": func(e model.Event) {},             // dihash ключ пароль (-nl)
	"master": func(e model.Event) {},             // master (-nl)
	"drop":   func(e model.Event) {},             // drop (-f)
	"stop":   func(e model.Event) { os.Exit(0) }, // stop
	"help":   func(e model.Event) {},             // help
}

var KEY_PARSE = map[string]func(command string) (event model.Event, _error bool){
	"hash": func(command string) (event model.Event, _error bool) {
		event = model.Event{}
		_error = false
		arr := strings.Split(command, " ")
		payload := getPaylaod(command)

		if len(arr) < 3 || payload == "" {
			_error = true
			return
		}

		event = model.Event{
			Time:     time.Now().Format("2006-01-02 15:04:05"),
			Key:      arr[0],
			Password: arr[1],
			Payload:  payload,
			Flags:    filterIsFlag(arr),
		}
		return
	},
	"g:key": func(command string) (event model.Event, _error bool) {
		event = model.Event{}
		_error = false
		return
	},
	"g:log": func(command string) (event model.Event, _error bool) {
		event = model.Event{}
		_error = false
		return
	},
	"c:log": func(command string) (event model.Event, _error bool) {
		event = model.Event{}
		_error = false
		return
	},
	"dihash": func(command string) (event model.Event, _error bool) {
		event = model.Event{}
		_error = false
		arr := strings.Split(command, " ")

		if len(arr) < 3 {
			_error = true
			return
		}

		event = model.Event{
			Time:     time.Now().Format("2006-01-02 15:04:05"),
			Key:      arr[0],
			Password: arr[1],
			Flags:    filterIsFlag(arr),
		}
		return
	},
	"master": func(command string) (event model.Event, _error bool) {
		event = model.Event{}
		_error = false
		return
	},
	"drop": func(command string) (event model.Event, _error bool) {
		event = model.Event{}
		_error = false
		return
	},
	"stop": func(command string) (event model.Event, _error bool) {
		event = model.Event{}
		_error = false
		return
	},
	"help": func(command string) (event model.Event, _error bool) {
		event = model.Event{}
		_error = false
		return
	},
}

var FLAG = map[string]string{
	"f":     "",
	"force": "",
	"nl":    "",
}
