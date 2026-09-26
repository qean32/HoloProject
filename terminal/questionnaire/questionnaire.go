package questionnaire

import (
	"main/constants/literals"
	"main/model"
	"main/terminal"
	"main/terminal/field"
)

func askQuestion(question model.Question) {
	terminal.TopLine()
	answer := field.Field(model.FieldPayload{
		Prefix:    "│ " + question.Message + " -> ",
		PrefixSRG: []int{literals.SGR.DIM},
	})

	if question.Callback(answer) {
		pushAnswer(question.Key, answer)
	}
}

func runQuestionnaire() {
	terminal.Println(questionnaire.Title)
	for _, question := range questionnaire.Questions {
		askQuestion(question)
	}
}
