package list

import "main/constants"

func getStartChar(isSelected bool) string {
	if isSelected {
		return constants.CHAR_SELECTED_ITEM
	}

	return constants.CHAR_UN_SELECTED_ITEM
}
