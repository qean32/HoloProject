package lib

import (
	"strings"
)

func AddAfterIndex[T any](array []T, item T, index int) []T {
	result := make([]T, 0, len(array)+1)
	result = append(result, array[:index]...)
	result = append(result, item)
	result = append(result, array[index:]...)
	return result
}

func RemoveByIndex[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

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
