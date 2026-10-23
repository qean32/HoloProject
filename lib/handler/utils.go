package handler

import (
	"path/filepath"

	"main/constants"
	"main/constants/literal"
	names "main/lib/names"
)

func validateName(value string) error {
	return names.Validate(value)
}

func cmdPath(keyword string) string {
	return filepath.Join(constants.Cmd, keyword+literal.Extension.Bat)
}
