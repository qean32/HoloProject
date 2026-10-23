package questionnaire

import (
	"errors"
	"testing"

	errorlist "main/constants/error-list"
	"main/model"
)

func TestAskAllRetriesInvalidAnswer(t *testing.T) {
	answers := []string{"no", "yes"}
	calls := 0
	ask := func(model.Question) (string, error) {
		answer := answers[calls]
		calls++
		return answer, nil
	}

	result, err := askAll(model.QuestionnairePayload{
		Questions: []model.Question{
			{
				Message:  "?",
				Key:      "answer",
				Callback: func(s string) bool { return s == "yes" },
			},
		},
	}, ask)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["answer"] != "yes" {
		t.Errorf("answer = %q, want %q", result["answer"], "yes")
	}
	if calls != 2 {
		t.Errorf("calls = %d, want 2", calls)
	}
}

func TestAskAllTooManyAttempts(t *testing.T) {
	calls := 0
	ask := func(model.Question) (string, error) {
		calls++
		return "no", nil
	}

	_, err := askAll(model.QuestionnairePayload{
		Questions: []model.Question{
			{
				Message:  "?",
				Key:      "answer",
				Callback: func(string) bool { return false },
			},
		},
	}, ask)
	if !errors.Is(err, errorlist.ErrTooManyAttempts) {
		t.Fatalf("error = %v, want %v", err, errorlist.ErrTooManyAttempts)
	}
	if calls != maxAttempts {
		t.Errorf("calls = %d, want %d", calls, maxAttempts)
	}
}

func TestAskAllNilCallback(t *testing.T) {
	result, err := askAll(model.QuestionnairePayload{
		Questions: []model.Question{{Message: "?", Key: "answer"}},
	}, func(model.Question) (string, error) {
		return "anything", nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result["answer"] != "anything" {
		t.Errorf("answer = %q, want %q", result["answer"], "anything")
	}
}

func TestIsAcceptableYea(t *testing.T) {
	tests := map[string]bool{
		"yes":  true,
		"YES":  true,
		" y ":  true,
		"да":   true,
		"Да":   true,
		"no":   false,
		"":     false,
		"yeah": false,
	}
	for answer, want := range tests {
		if got := IsAcceptableYea(answer); got != want {
			t.Errorf("IsAcceptableYea(%q) = %v, want %v", answer, got, want)
		}
	}
}
