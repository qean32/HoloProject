package model

type Questionnaire struct {
	Questions []Question
	Result    map[string]string
}
