package questionnaire

import (
	"main/constants"
	"slices"
)

func isAcceptableYeaOrNot(answer string) bool {
	return slices.Index(constants.AcceptableYeaOrNot, answer) != -1
}
