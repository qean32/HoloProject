package questionnaire

import (
	"main/constants/literals"
	"main/model"
	"main/terminal"
	"main/terminal/field"
)

func askQuestion(question model.Question) {
	terminal.PrintTopLine()
	_answer := field.Field(model.FieldPayload{
		Prefix:    "│ " + question.Message + " -> ",
		PrefixSRG: []int{literals.SGR.DIM},
	})

	if question.Callback(_answer) {
		answer(question.Key, _answer)
	}
}

func runQuestionnaire() {
	terminal.Println(questionnaire.Title)
	for _, question := range questionnaire.Questions {
		askQuestion(question)
	}
}

func answer(key, answer string) {
	questionnaire.Result[key] = answer
	terminal.DownAndStart()
}
