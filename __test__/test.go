package test

import (
	"main/callstack"
	eventlist "main/constants/event-list"
	"main/model"
)

func Test() error {
	callstack.PushCallStackMany(
		model.Event{Key: eventlist.PURE_REMOVE_BAT, KeyWord: "test"},
		model.Event{Key: eventlist.RUNBAT, KeyWord: "test"},
		model.Event{Key: eventlist.BAT, KeyWord: "test", Payload: "start cmd /K echo 67"},
	)
	// callstack.PushCallStackMany(
	// 	model.Event{Key: eventlist.PURE_REMOVE_BAT, KeyWord: "test2"},
	// 	model.Event{Key: eventlist.RUNBAT, KeyWord: "test2"},
	// 	model.Event{Key: eventlist.BAT, KeyWord: "test2", Payload: "start brave"},
	// )

	return nil
}
