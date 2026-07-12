package main

import (
	"main/lib"
	"main/lib/low"
)

func main() {
	lib.INIT()
}

func LOOP() {
	for value := range low.Callstack_channel {
		lib.ITERATION_CYCLE(value)
	}
}
