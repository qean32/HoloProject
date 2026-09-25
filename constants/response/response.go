package response

import (
	"main/constants"
	"main/constants/literals"
	"main/model"
)

func UndefinedCommand() model.Event {
	return model.Event{
		Key:     literals.COMMANDLIST.RESPONSE,
		Payload: constants.UndefinedCommand,
	}
}

func UndefinedKeyword() model.Event {
	return model.Event{
		Key:     literals.COMMANDLIST.RESPONSE,
		Payload: constants.UndefinedKeyword,
	}
}

func SyntaxError() model.Event {
	return model.Event{
		Key:     literals.COMMANDLIST.RESPONSE,
		Payload: constants.SyntaxError,
	}
}

func Success() model.Event {
	return model.Event{
		Key:     literals.COMMANDLIST.RESPONSE,
		Payload: constants.Success,
	}
}
