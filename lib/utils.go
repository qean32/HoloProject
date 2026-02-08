package lib

import (
	"strings"
)

func getPaylaod(command string) string {
	var payload string
	start := strings.IndexAny(command, "{")
	if start != -1 {
		end := strings.IndexAny(command, "}")
		payload = command[start+1 : end]
	}

	return payload
}
