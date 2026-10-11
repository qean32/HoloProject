package handler

import (
	"main/callstack"
	errorlist "main/constants/error-list"
	"main/constants/literal"
	"main/model"
	"main/terminal/questionnaire"
)

func questionnaireAccess(event model.Event) error {
	if event.SubEvent == nil {
		return errorlist.ErrNoSubEvent
	}

	q := questionnaire.Questionnaire(model.QuestionnairePayload{
		Questions: []model.Question{
			{
				Message:  "",
				Key:      "Response",
				Callback: func(string) bool { return true },
			},
		},
		Title: literal.Titles.Access,
	})

	if !questionnaire.IsAcceptableYea(q["Response"]) {
		return errorlist.ErrAccessDenied
	}

	callstack.PushCallStack(*event.SubEvent)
	return nil
}
