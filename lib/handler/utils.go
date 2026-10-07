package handler

import (
	"errors"
	"strings"

	"main/constants"
	"main/constants/literal"
	"path/filepath"
)

func validateName(name string) error {
	if name == "" || name == "." || name == ".." {
		return errors.New("Недопустимое имя")
	}
	if strings.ContainsAny(name, `/\:*?"<>|`) {
		return errors.New("Недопустимые символы")
	}
	return nil
}

func cmdPath(keyword string) string {
	return filepath.Join(constants.Root, constants.Cmd, keyword+literal.Extension.Bat)
}

func trimExtention(str string) string {
	return strings.TrimSuffix(str, literal.Extension.Bat)
}
