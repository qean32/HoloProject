package response

import (
	"main/constants"
	"main/terminal"
)

func UNDEFINED_COMMAND() {
	terminal.OutputResponse(constants.UNDEFINED_COMMAND)
}

func UNDEFINED_KEYWORD() {
	terminal.OutputResponse(constants.UNDEFINED_KEYWORD)
}

func SYNTAX_ERROR() {
	terminal.OutputResponse(constants.SYNTAX_ERROR)
}
