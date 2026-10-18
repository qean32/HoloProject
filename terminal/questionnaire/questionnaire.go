package questionnaire

import (
	"main/constants/sgr"
	"main/model"
	"main/terminal"
	"main/terminal/field"
)

func Questionnaire(payload model.QuestionnairePayload) map[string]string {
	result := make(map[string]string, len(payload.Questions))

	terminal.Println(payload.Title)
	for _, question := range payload.Questions {
		if answer := askQuestion(question); question.Callback(answer) {
			result[question.Key] = answer
		} else {
			newQuestion := question
			newQuestion.Message = "invalid response " + question.Message
			askQuestion(newQuestion)
		}
	}
	return result
}

func askQuestion(question model.Question) string {
	return field.Field(model.FieldPayload{
		Prefix:    question.Message + " > ",
		PrefixSRG: []int{sgr.DIM},
	})
}
