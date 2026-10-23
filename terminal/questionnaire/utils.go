package questionnaire

import (
	"slices"
	"strings"

	"main/constants"
)

func IsAcceptableYea(answer string) bool {
	answer = strings.TrimSpace(answer)
	return slices.ContainsFunc(constants.AcceptableYeaOrNot, func(v string) bool {
		return strings.EqualFold(answer, v)
	})
}
