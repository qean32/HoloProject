package lib

import (
	"strings"
)

func getPaylaod(arr []string) string {
	var payload string
	command := strings.Join(arr, " ")
	start := strings.IndexAny(command, "{")
	if start != -1 {
		end := strings.IndexAny(command, "}")
		payload = command[start+1 : end]
	}

	return payload
}

func matrixToArrayString(matrix [][]string) []string {
	var tmpArr []string

	for i := 0; i < len(matrix); i++ {
		tmpArr = append(tmpArr, strings.Join(matrix[i], " "))
	}

	return tmpArr
}
