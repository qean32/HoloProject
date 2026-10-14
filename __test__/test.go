package test

import (
	"fmt"
	"main/callstack"
	"main/constants"
	eventlist "main/constants/event-list"
	"main/lib/death"
	"main/model"
	"path/filepath"
)

func Test() {
	callstack.PushMany(
		model.Event{Key: eventlist.TESTVIEW},
		model.Event{Key: eventlist.PURE_REMOVE_BAT, KeyWord: "test"},
		model.Event{Key: eventlist.RUNBAT, KeyWord: "test"},
		model.Event{Key: eventlist.BAT, KeyWord: "test", Payload: "start cmd /K echo 67"},
	)
}

func TestView() {
	if err := death.WriteFile(callstack.GetErrorsString(), filepath.Join("tmplog.asc")); err != nil {
		fmt.Println(err)
		return
	}

	defer callstack.PushMany(
		model.Event{Key: eventlist.PURE_REMOVE_BAT, KeyWord: "tmpbat"},
		model.Event{Key: eventlist.RUNBAT, KeyWord: "tmpbat"},
		model.Event{Key: eventlist.BAT, KeyWord: "tmpbat", Payload: "start " + filepath.Join(constants.Root, "tmplog.asc")},
	)
}
