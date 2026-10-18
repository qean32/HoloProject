package handler

import (
	"path/filepath"
	"strings"

	"main/constants"
	errorlist "main/constants/error-list"
	"main/constants/literal"
)

func validateName(name string) error {
	if name == "" || name == "." || name == ".." {
		return errorlist.ErrForbiddenName
	}
	if strings.ContainsAny(name, `/\:*?"<>|`) {
		return errorlist.ErrForbiddenChars
	}
	return nil
}

func cmdPath(keyword string) string {
	return filepath.Join(constants.Cmd, keyword+literal.Extension.Bat)
}

func trimExtension(str string) string {
	return strings.TrimSuffix(str, literal.Extension.Bat)
}
