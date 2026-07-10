package list

import (
	"main/constants"

	"atomicgo.dev/cursor"
)

func getStartChar(isSelected bool) string {
	if isSelected {
		return constants.CHAR_SELECTED_ITEM
	}

	return constants.CHAR_UN_SELECTED_ITEM
}

func jumpToEndList() {
	cursor.Down(list.Length - list.Position)
}

func jumpToStartList() {
	cursor.Up(list.Length - list.Position)
}
