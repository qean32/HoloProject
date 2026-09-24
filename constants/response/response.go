package response

import (
	"main/constants"
	"main/terminal"
)

func Undefinedcommand() {
	terminal.PrintResponse(constants.Undefinedcommand)
}

func UndefinedKeyword() {
	terminal.PrintResponse(constants.UndefinedKeyword)
}

func SyntaxError() {
	terminal.PrintResponse(constants.SyntaxError)
}

func Success() {
	terminal.PrintResponse(constants.SyntaxError)
}
