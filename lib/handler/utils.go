package handler

import (
	"path/filepath"
	"strings"

	"main/constants"
	errorlist "main/constants/error-list"
	"main/constants/literal"
	"main/terminal"

	"atomicgo.dev/cursor"
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

func Help() {
	cursor.StartOfLine()
	terminal.PrintASCIICenter(constants.HelpMessage, "")
}
