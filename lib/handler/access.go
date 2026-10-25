package handler

import (
	errorlist "main/constants/error-list"
	"main/model"
	"main/terminal/questionnaire"
)

func QAccess(event model.Event, warning string) error {
	if event.SubEvent == nil {
		return errorlist.ErrNoSubEvent
	}

	q := questionnaire.Questionnaire(model.QuestionnairePayload{
		Questions: []model.Question{
			{
				Message:  "Подтвердите действие (yea/no) " + warning,
				Key:      "Response",
				Callback: func(string) bool { return true },
			},
		},
		Title: "",
	})

	if !questionnaire.IsAcceptableYea(q["Response"]) {
		return errorlist.ErrAccessDenied
	}

	return nil
}
