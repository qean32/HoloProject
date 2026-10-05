package handler

import (
	"main/callstack"
	"main/constants/literal"
	"main/model"
	"main/terminal/questionnaire"
)

func questionnaireAccess(event model.Event) error {
	if event.SubEvent == nil {
		return ErrNoSubEvent
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
		return ErrAccessDenied
	}

	callstack.PushCallStack(*event.SubEvent)
	return nil
}
