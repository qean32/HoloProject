package list

import (
	"strings"

	"main/model"
)

func filterIndices(options []model.Option, query string) []int {
	needle := strings.ToLower(strings.TrimSpace(query))

	indices := make([]int, 0, len(options))
	for i, option := range options {
		if needle == "" || strings.Contains(strings.ToLower(option.Message), needle) {
			indices = append(indices, i)
		}
	}
	return indices
}
