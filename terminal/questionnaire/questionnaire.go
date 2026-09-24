package questionnaire

import (
	"main/constants/literals"
	"main/model"
	"main/terminal"
	"main/terminal/field"
)

func start(message string) {
	terminal.Println(terminal.GetCustomMessage("─────────", literals.SGR.DIM))
	terminal.Print(terminal.GetCustomMessage(message, literals.SGR.DIM))
}

func askQuestion(question model.Question) {
	start(question.Message)
	terminal.DownAndStart()
	answer := field.Field()

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
