package questionnaire

import (
	"main/model"
)

var questionnaire = model.Questionnaire{
	Questions: []model.Question{},
	Result:    map[string]string{},
}

func Questionnaire(payload model.QuestionnairePayload) map[string]string {
	set(payload)
	runQuestionnaire()

	return questionnaire.Result
}

func reset() {
	questionnaire.Questions = []model.Question{}
	questionnaire.Result = map[string]string{}
}

func set(payload model.QuestionnairePayload) {
	questionnaire.Questions = payload.Questions
	questionnaire.Title = payload.Title
	questionnaire.Result = map[string]string{}
}
