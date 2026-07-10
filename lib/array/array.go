package array

import "strings"

func AddAfterIndex[T any](array []T, item T, index int) []T {
	result := make([]T, 0, len(array)+1)
	result = append(result, array[:index]...)
	result = append(result, item)
	result = append(result, array[index:]...)
	return result
}

func Map[T any, U any](slice []T, fn func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func RemoveByIndex[T any](slice []T, index int) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func MatrixToArrayString(matrix [][]string) []string {
	var tmpArr []string

	for i := 0; i < len(matrix); i++ {
		tmpArr = append(tmpArr, strings.Join(matrix[i], " "))
	}

	return tmpArr
}
