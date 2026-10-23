package test

// import (
// 	"testing"

// 	"main/model"
// )

// func options(names ...string) []model.Option {
// 	result := make([]model.Option, len(names))
// 	for i, name := range names {
// 		result[i] = model.Option{Message: name}
// 	}
// 	return result
// }

// func TestFilterIndices(t *testing.T) {
// 	all := options("alpha", "Beta", "gamma", "alphabet")

// 	tests := []struct {
// 		query string
// 		want  []int
// 	}{
// 		{"", []int{0, 1, 2, 3}},
// 		{"alp", []int{0, 3}},
// 		{"ALP", []int{0, 3}},
// 		{"beta", []int{1}},
// 		{"zzz", nil},
// 	}

// 	for _, tt := range tests {
// 		got := filterIndices(all, tt.query)
// 		if len(got) != len(tt.want) {
// 			t.Fatalf("filterIndices(%q) = %v, want %v", tt.query, got, tt.want)
// 		}
// 		for i := range got {
// 			if got[i] != tt.want[i] {
// 				t.Errorf("filterIndices(%q)[%d] = %d, want %d", tt.query, i, got[i], tt.want[i])
// 			}
// 		}
// 	}
// }

// func TestEnsureVisible(t *testing.T) {
// 	s := &state{filtered: make([]int, 100), position: 50}
// 	s.ensureVisible(10)
// 	if s.offset != 41 {
// 		t.Errorf("offset = %d, want 41", s.offset)
// 	}

// 	s.position = 0
// 	s.ensureVisible(10)
// 	if s.offset != 0 {
// 		t.Errorf("offset = %d, want 0", s.offset)
// 	}

// 	s.position = 99
// 	s.ensureVisible(10)
// 	if s.offset != 90 {
// 		t.Errorf("offset = %d, want 90", s.offset)
// 	}

// 	s.position = 5
// 	s.offset = 20
// 	s.ensureVisible(10)
// 	if s.offset != 5 {
// 		t.Errorf("offset = %d, want 5", s.offset)
// 	}
// }

// func TestTruncate(t *testing.T) {
// 	tests := []struct {
// 		value string
// 		width int
// 		want  string
// 	}{
// 		{"abcdef", 4, "abc…"},
// 		{"abc", 4, "abc"},
// 		{"abcdef", 1, "a"},
// 		{"привет", 4, "при…"},
// 		{"abc", 0, ""},
// 	}
// 	for _, tt := range tests {
// 		if got := truncate(tt.value, tt.width); got != tt.want {
// 			t.Errorf("truncate(%q, %d) = %q, want %q", tt.value, tt.width, got, tt.want)
// 		}
// 	}
// }
