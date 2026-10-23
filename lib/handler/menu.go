package handler

import (
	"fmt"

	"atomicgo.dev/cursor"

	"main/constants"
	"main/terminal"
)

func help() {
	cursor.StartOfLine()
	terminal.PrintASCIICenter(constants.HelpMessage, "")
	terminal.Println()

	for _, command := range visible {
		terminal.Println(fmt.Sprintf("%-10s %-34s %s", command.Name, command.Title, command.Usage))
	}
	terminal.Println()
}
