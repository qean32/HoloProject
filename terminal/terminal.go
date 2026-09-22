package terminal

import (
	"fmt"
	"sync"
)

var (
	mu sync.Mutex
)

func Output(message string) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Print(message)
}

func Outputln(message string) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Println(message)
}

func OutputTechInfo(messages ...any) {
	var result string
	for _, msg := range messages {
		result += fmt.Sprint(msg)
	}
	Outputln("\033[31mТехнический вывод: " + result + "\033[0m")
}
