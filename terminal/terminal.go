package terminal

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
)

func Output(messages ...any) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Print(" ", calsMessages(messages...))
}

func Outputln(messages ...any) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println(" ", calsMessages(messages...))
}

func OutputTechInfo(messages ...any) {
	// Outputln("\033[31m" + "Технический вывод: " + calsMessages(messages...) + "\033[0m")
}

func OutputResponse(messages ...any) {
	fmt.Println(append([]any{" <- "}, messages...)...)
}

func calsMessages(messages ...any) string {
	var result string
	for _, msg := range messages {
		result += fmt.Sprint(msg)
	}

	return result
}
