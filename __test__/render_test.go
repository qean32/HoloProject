package test

// import (
// 	"fmt"
// 	"strings"
// 	"testing"
// )

// func TestFrameLayout(t *testing.T) {
// 	s := &state{title: "T", all: options("alpha", "beta", "gamma")}
// 	s.refilter()

// 	lines := s.frame(40, 10)
// 	if len(lines) != 10 {
// 		t.Fatalf("len(lines) = %d, want 10", len(lines))
// 	}
// 	if !strings.Contains(lines[0], "T") || !strings.Contains(lines[0], "[3/3]") {
// 		t.Errorf("header = %q", lines[0])
// 	}
// 	if !strings.Contains(lines[2], "> ") {
// 		t.Errorf("selected row = %q, want marker", lines[2])
// 	}
// 	if !strings.Contains(lines[len(lines)-1], "Enter") {
// 		t.Errorf("footer = %q, want Enter hint", lines[len(lines)-1])
// 	}
// }

// func TestFrameViewport(t *testing.T) {
// 	names := make([]string, 100)
// 	for i := range names {
// 		names[i] = fmt.Sprintf("cmd-%02d", i)
// 	}

// 	s := &state{title: "T", all: options(names...)}
// 	s.refilter()
// 	s.position = 50

// 	joined := strings.Join(s.frame(40, 10), "\n")
// 	if !strings.Contains(joined, "cmd-50") {
// 		t.Error("selected item is not visible")
// 	}
// 	if strings.Contains(joined, "cmd-00") {
// 		t.Error("first item should be scrolled out of the viewport")
// 	}
// }

// func TestFrameNoMatches(t *testing.T) {
// 	s := &state{title: "T", all: options("alpha", "beta")}
// 	s.query = "zzz"
// 	s.refilter()

// 	joined := strings.Join(s.frame(40, 10), "\n")
// 	if !strings.Contains(joined, "ничего не найдено") {
// 		t.Error("missing empty-state message")
// 	}
// }
