package terminal

import (
	"fmt"
	"sync"

	"main/constants"
	"main/lib/utils"
)

var mu sync.Mutex

func Print(messages ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Print(" " + utils.ConcatMessage(messages...))
}

func Println(messages ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(" " + utils.ConcatMessage(messages...))
}

func PrintResponse(messages ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(append([]any{constants.RESPONSEPREFIX}, messages...)...)
	fmt.Println("")
}

func PrintTopLine() {
	// Println(utils.GetCustomMessage("┌──", sgr.DIM))
	Println("")
}
