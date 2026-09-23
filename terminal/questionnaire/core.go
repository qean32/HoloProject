package questionnaire

import "main/model"

var questionnaire = model.Questionnaire{
	Questions: []model.Question{},
	Result:    map[string]string{},
}

func Questionnaire(questions []model.Question) map[string]string {
	set(questions)
	runQuestionnaire()

	return questionnaire.Result
}

func reset() {
	questionnaire.Questions = []model.Question{}
	questionnaire.Result = map[string]string{}
}

func set(questions []model.Question) {
	questionnaire.Questions = questions
	questionnaire.Result = map[string]string{}
}

func pushAnswer(key, answer string) {
	questionnaire.Result[key] = answer
}
