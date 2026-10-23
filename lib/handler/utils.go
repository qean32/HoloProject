package handler

import (
	"fmt"
	"path/filepath"
	"strings"

	"main/constants"
	"main/constants/literal"
	names "main/lib/names"
	"main/terminal"

	"atomicgo.dev/cursor"
)

func validateName(value string) error {
	return names.Validate(value)
}

func cmdPath(keyword string) string {
	return filepath.Join(constants.Cmd, keyword+literal.Extension.Bat)
}

func trimExtension(str string) string {
	return strings.TrimSuffix(str, literal.Extension.Bat)
}

func help() {
	cursor.StartOfLine()
	terminal.PrintASCIICenter(constants.HelpMessage, "")
	terminal.Println()

	for _, command := range visible {
		terminal.Println(fmt.Sprintf("%-10s %-34s %s", command.Name, command.Title, command.Usage))
	}
	terminal.Println()
}
