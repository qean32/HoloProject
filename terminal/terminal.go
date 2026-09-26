package terminal

import (
	"fmt"
	"sync"

	"main/constants"
	"main/constants/literals"
	"main/lib/utils"
)

var mu sync.Mutex

func Print(messages ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Print(utils.ConcatMessage(messages...))
}

func Println(messages ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(utils.ConcatMessage(messages...))
}

func PrintResponse(messages ...any) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println(append([]any{constants.RESPONSEPREFIX}, messages...)...)
}

func PrintTopLine() {
	Println(utils.GetCustomMessage("┌──", literals.SGR.DIM))
}

func PrintTechInfo(messages ...any) {
	fmt.Printf(utils.GetCustomMessage("Технический вывод: %s", literals.SGR.GREEN), utils.ConcatMessage(messages...))
}
