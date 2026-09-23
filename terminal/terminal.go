package terminal

import (
	"fmt"
	"strings"
	"sync"
)

var mu sync.Mutex

func Print(messages ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Print(" ", concatMessages(messages...))
}

func Println(messages ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(" ", concatMessages(messages...))
}

func PrintTechInfo(messages ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("\033[31mТехнический вывод: %s\033[0m\n", concatMessages(messages...))
}

func PrintResponse(messages ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(append([]any{" <- "}, messages...)...)
}

func concatMessages(messages ...any) string {
	var sb strings.Builder
	for _, msg := range messages {
		sb.WriteString(fmt.Sprint(msg))
	}
	return sb.String()
}
