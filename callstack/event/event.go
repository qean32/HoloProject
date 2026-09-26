package event

import (
	"main/callstack"
	"main/constants/literals"
	"main/lib/utils"
	"main/model"
)

func Response(message ...any) {
	callstack.PushCallStack(model.Event{
		Key:     literals.EventList.RESPONSE,
		Payload: utils.ConcatMessage(message...),
	})
}
