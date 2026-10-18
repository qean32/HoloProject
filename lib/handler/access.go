package handler

import (
	"main/callstack"
	errorlist "main/constants/error-list"
	"main/model"
	"main/terminal/questionnaire"
)

func qAccess(event model.Event) error {
	if event.SubEvent == nil {
		return errorlist.ErrNoSubEvent
	}

	q := questionnaire.Questionnaire(model.QuestionnairePayload{
		Questions: []model.Question{
			{
				Message:  "Подтвердите действие (yea/no)",
				Key:      "Response",
				Callback: func(string) bool { return true },
			},
		},
		Title: "",
	})

	if !questionnaire.IsAcceptableYea(q["Response"]) {
		return errorlist.ErrAccessDenied
	}

	callstack.Push(*event.SubEvent)
	return nil
}
