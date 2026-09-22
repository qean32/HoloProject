package terminal

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
)

func Output(messages ...string) {
	var result string
	for _, msg := range messages {
		result += fmt.Sprint(msg)
	}
	mu.Lock()
	defer mu.Unlock()

	fmt.Print(" ", result)
}

func Outputln(messages ...string) {
	var result string
	for _, msg := range messages {
		result += fmt.Sprint(msg)
	}
	mu.Lock()
	defer mu.Unlock()

	fmt.Println(" ", result)
}

func OutputTechInfo(messages ...any) {
	var result string
	for _, msg := range messages {
		result += fmt.Sprint(msg)
	}
	// Outputln("\033[31m" + "Технический вывод: " + result + "\033[0m")
}
