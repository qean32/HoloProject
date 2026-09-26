package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func GetCustomMessage(message string, sgr ...int) string {
	if len(sgr) == 0 {
		return message
	}
	parts := make([]string, len(sgr))
	for i, v := range sgr {
		parts[i] = strconv.Itoa(v)
	}
	return fmt.Sprintf("\033[%sm%s\033[0m", strings.Join(parts, ";"), message)
}

func ConcatMessage(messages ...any) string {
	var sb strings.Builder
	for _, msg := range messages {
		sb.WriteString(fmt.Sprint(msg))
	}
	return sb.String()
}
