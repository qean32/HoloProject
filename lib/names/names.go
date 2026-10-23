package names

import (
	"strings"
	"unicode/utf8"

	errorlist "main/constants/error-list"
)

const MaxLength = 100

var reserved = map[string]struct{}{
	"CON": {}, "PRN": {}, "AUX": {}, "NUL": {},
	"COM1": {}, "COM2": {}, "COM3": {}, "COM4": {}, "COM5": {},
	"COM6": {}, "COM7": {}, "COM8": {}, "COM9": {},
	"LPT1": {}, "LPT2": {}, "LPT3": {}, "LPT4": {}, "LPT5": {},
	"LPT6": {}, "LPT7": {}, "LPT8": {}, "LPT9": {},
}

// Validate проверяет имя файла команды: без разделителей пути,
// управляющих символов, хвостовых точек/пробелов и зарезервированных имён Windows.
func Validate(value string) error {
	if value == "" || value == "." || value == ".." {
		return errorlist.ErrForbiddenName
	}
	if strings.TrimSpace(value) != value || strings.HasSuffix(value, ".") {
		return errorlist.ErrForbiddenName
	}
	if utf8.RuneCountInString(value) > MaxLength {
		return errorlist.ErrForbiddenName
	}
	for _, r := range value {
		if r < 0x20 || r == 0x7f {
			return errorlist.ErrForbiddenChars
		}
	}
	if strings.ContainsAny(value, `/\:*?"<>|`) {
		return errorlist.ErrForbiddenChars
	}
	base, _, _ := strings.Cut(value, ".")
	if _, ok := reserved[strings.ToUpper(base)]; ok {
		return errorlist.ErrForbiddenName
	}
	return nil
}
