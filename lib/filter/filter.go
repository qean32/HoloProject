package filter

func FilterIsFlag[T string](arr []T) []T {
	return FILTER(arr, isFlag)
}

func FILTER[T any](arr []T, f func(item T) bool) []T {
	filtered := make([]T, 0, len(arr))

	for _, e := range arr {
		if f(e) {
			filtered = append(filtered, e)
		}
	}

	return filtered
}

func isFlag[T string](item T) bool {
	return item[0:1] == "-"
}
