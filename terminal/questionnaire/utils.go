package questionnaire

import (
	"main/constants"
	"slices"
)

func isAcceptableYeaOrNot(answer string) bool {
	return slices.Contains(constants.AcceptableYeaOrNot, answer)
}
