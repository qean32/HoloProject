package test

// import (
// 	"strings"
// 	"testing"

// 	names "main/lib/names"
// )

// func TestValidateName(t *testing.T) {
// 	valid := []string{
// 		"x",
// 		"my-command_1",
// 		"Привет",
// 		"a.b",
// 		"test",
// 	}
// 	for _, name := range valid {
// 		if err := validateName(name); err != nil {
// 			t.Errorf("validateName(%q) = %v, want nil", name, err)
// 		}
// 	}

// 	invalid := []string{
// 		"",
// 		".",
// 		"..",
// 		"a/b",
// 		`a\b`,
// 		"a:b",
// 		"a*b",
// 		"a?b",
// 		`a"b`,
// 		"<a",
// 		"a>",
// 		"a|",
// 		"CON",
// 		"con",
// 		"CON.txt",
// 		"PRN",
// 		"COM1",
// 		"LPT9",
// 		"name.",
// 		" name",
// 		"name ",
// 		"a\x00b",
// 		strings.Repeat("a", names.MaxLength+1),
// 	}
// 	for _, name := range invalid {
// 		if err := validateName(name); err == nil {
// 			t.Errorf("validateName(%q) = nil, want error", name)
// 		}
// 	}
// }
