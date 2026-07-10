package questionnaire

import (
	"main/model"
	"main/terminal"
	"main/terminal/field"
)

func askQuestion(question model.Question) {
	terminal.Output(terminal.GetCustomMessage(question.Message))
	terminal.DownAndStart()
	answer := field.Field()
	terminal.OutputTechInfo("", answer, question.Callback(answer))

	if question.Callback(answer) {
		pushAnswer(question.Key, answer)
	}
}

func runQuestionnaire() {
	for _, question := range questionnaire.Questions {
		askQuestion(question)
	}
}
