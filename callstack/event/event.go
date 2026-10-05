package event

import (
	"main/callstack"
	"main/constants/event-list"
	"main/lib/utils"
	"main/model"
)

func Response(message ...any) {
	callstack.PushCallStack(model.Event{
		Key:     eventlist.RESPONSE,
		Payload: utils.ConcatMessage(message...),
	})
}
