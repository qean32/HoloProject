package test
// package parse

// import (
// 	"errors"
// 	"testing"

// 	errorlist "main/constants/error-list"
// 	eventlist "main/constants/event-list"
// 	"main/model"
// )

// var (
// 	batOptions    = model.ParseOptions{KeyWord: true, Payload: true}
// 	rbatOptions   = model.ParseOptions{KeyWord: true, GetSubEvent: true, SubEvent: eventlist.PUREREMOVEBAT}
// 	cryptoOptions = model.ParseOptions{KeyWord: true, Payload: true, Password: true}
// )

// func TestParseBatPayload(t *testing.T) {
// 	tests := []struct {
// 		name    string
// 		args    []string
// 		payload string
// 	}{
// 		{"slash prefix", []string{"bat", "x", "p/echo hi"}, "echo hi"},
// 		{"flag prefix", []string{"bat", "x", "-p:echo hi"}, "echo hi"},
// 		{"several words", []string{"bat", "x", "p/echo", "hello", "world"}, "echo hello world"},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			event, err := Parse(tt.args, batOptions)
// 			if err != nil {
// 				t.Fatalf("unexpected error: %v", err)
// 			}
// 			if event.KeyWord != "x" {
// 				t.Errorf("KeyWord = %q, want %q", event.KeyWord, "x")
// 			}
// 			if event.Payload != tt.payload {
// 				t.Errorf("Payload = %q, want %q", event.Payload, tt.payload)
// 			}
// 		})
// 	}
// }

// func TestParseErrors(t *testing.T) {
// 	if _, err := Parse([]string{"bat"}, batOptions); !errors.Is(err, errorlist.ErrTooFewArgs) {
// 		t.Errorf("missing keyword: got %v, want %v", err, errorlist.ErrTooFewArgs)
// 	}
// 	if _, err := Parse([]string{"bat", "x"}, batOptions); !errors.Is(err, errorlist.ErrNoPayload) {
// 		t.Errorf("missing payload: got %v, want %v", err, errorlist.ErrNoPayload)
// 	}
// 	if _, err := Parse(nil, batOptions); !errors.Is(err, errorlist.ErrEmptyCommand) {
// 		t.Errorf("empty args: got %v, want %v", err, errorlist.ErrEmptyCommand)
// 	}
// }

// func TestParseRemoveSubEvent(t *testing.T) {
// 	event, err := Parse([]string{"rbat", "x"}, rbatOptions)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if event.SubEvent == nil {
// 		t.Fatal("SubEvent is nil")
// 	}
// 	if event.SubEvent.Key != eventlist.PUREREMOVEBAT {
// 		t.Errorf("SubEvent.Key = %q, want %q", event.SubEvent.Key, eventlist.PUREREMOVEBAT)
// 	}
// 	if event.SubEvent.KeyWord != "x" {
// 		t.Errorf("SubEvent.KeyWord = %q, want %q", event.SubEvent.KeyWord, "x")
// 	}
// }

// func TestParsePassword(t *testing.T) {
// 	event, err := Parse([]string{"crypto", "key", "p/data", "pass/secret"}, cryptoOptions)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if event.Password != "secret" {
// 		t.Errorf("Password = %q, want %q", event.Password, "secret")
// 	}
// 	if event.Payload != "data" {
// 		t.Errorf("Payload = %q, want %q", event.Payload, "data")
// 	}
// }

// func TestParseFlagsIgnorePayload(t *testing.T) {
// 	event, err := Parse([]string{"bat", "x", "p/echo", "-nl"}, batOptions)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if len(event.Flags) != 0 {
// 		t.Errorf("Flags = %v, want none", event.Flags)
// 	}
// }

// func TestParseFlagsBeforePayload(t *testing.T) {
// 	event, err := Parse([]string{"bat", "x", "-nl", "p/echo"}, batOptions)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if len(event.Flags) != 1 || event.Flags[0] != "-nl" {
// 		t.Errorf("Flags = %v, want [-nl]", event.Flags)
// 	}
// }

// func TestParseKey(t *testing.T) {
// 	event, err := Parse([]string{"list"}, model.ParseOptions{})
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// 	if event.Key != "list" {
// 		t.Errorf("Key = %q, want %q", event.Key, "list")
// 	}
// }
