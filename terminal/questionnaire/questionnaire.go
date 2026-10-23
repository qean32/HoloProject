package questionnaire

import (
	errorlist "main/constants/error-list"
	"main/constants/sgr"
	"main/model"
	"main/terminal"
	"main/terminal/field"
)

const maxAttempts = 3

type prompter func(model.Question) (string, error)

func Questionnaire(payload model.QuestionnairePayload) (map[string]string, error) {
	return askAll(payload, askQuestion)
}

func askAll(payload model.QuestionnairePayload, ask prompter) (map[string]string, error) {
	result := make(map[string]string, len(payload.Questions))
	if payload.Title != "" {
		terminal.Println(payload.Title)
	}

	for _, question := range payload.Questions {
		answer, err := askValid(question, ask)
		if err != nil {
			return result, err
		}
		result[question.Key] = answer
	}
	return result, nil
}

func askValid(question model.Question, ask prompter) (string, error) {
	for attempt := 0; attempt < maxAttempts; attempt++ {
		answer, err := ask(question)
		if err != nil {
			return "", err
		}
		if question.Callback == nil || question.Callback(answer) {
			return answer, nil
		}
		terminal.Println("Некорректный ответ, попробуйте ещё раз.")
	}
	return "", errorlist.ErrTooManyAttempts
}

func askQuestion(question model.Question) (string, error) {
	return field.Field(model.FieldPayload{
		Prefix:    question.Message + " > ",
		PrefixSRG: []int{sgr.DIM},
	})
}
