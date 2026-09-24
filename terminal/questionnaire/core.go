package questionnaire

import (
	"main/constants/literals"
	"main/model"
	"main/terminal"
)

var questionnaire = model.Questionnaire{
	Questions: []model.Question{},
	Result:    map[string]string{},
}

func Questionnaire(questions []model.Question, title string) map[string]string {
	terminal.Println(terminal.GetCustomMessage("────────────────────────────────", literals.SGR.DIM))
	set(questions, title)
	runQuestionnaire()

	return questionnaire.Result
}

func reset() {
	questionnaire.Questions = []model.Question{}
	questionnaire.Result = map[string]string{}
}

func set(questions []model.Question, title string) {
	questionnaire.Questions = questions
	questionnaire.Title = title
	questionnaire.Result = map[string]string{}
}

func pushAnswer(key, answer string) {
	questionnaire.Result[key] = answer
}
