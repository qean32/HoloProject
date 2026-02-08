package lib

func filterIsFlag[T string](arr []T) []T {
	return FILTER(arr, isFlag)
}

func isFlag[T string](item T) bool {
	return item[0:1] == "-"
}

func FILTER[T any](arr []T, f func(item T) bool) []T {
	fltd := make([]T, 0, len(arr))

	for _, e := range arr {
		if f(e) {
			fltd = append(fltd, e)
		}
	}

	return fltd
}
