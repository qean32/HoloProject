package questionnaire

import (
	"main/constants"
	"slices"
)

func IsAcceptableYea(answer string) bool {
	return slices.Contains(constants.AcceptableYeaOrNot, answer)
}
