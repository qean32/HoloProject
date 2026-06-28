package main

import (
	"atomicgo.dev/cursor"
	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"fmt"
	"os"
)

type POINT struct {
	x int
	y int
}

type FIELD struct {
	position      int
	positionRange int
	field         string
}

func AddChar() {
}

func RemoveChar() {
}

func EnterCommand() {
}

func MoveCursorLeft() {
}

func MoveCursorRight() {
}

func StopProccess() {
}

func resetField() {
}

func main() {
	command := ""
	maxRight := 0
	currentPos := 0
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.Escape || key.Code == keys.CtrlC {
			os.Exit(0)
		}
		if key.Code == keys.Enter {
			fmt.Println("")
			cursor.StartOfLineDown(1)
			fmt.Println(command)
			maxRight = 0
			currentPos = 0
			command = ""
			return false, nil
		}
		if key.Code == keys.Right {
			if currentPos < maxRight {
				cursor.Right(1)
				currentPos++
			}
		}
		if key.Code == keys.Left {
			cursor.Left(1)
			if currentPos > 0 {
				currentPos--
			}
		}
		char := key.String()
		if key.Code != keys.Backspace || key.Code == keys.Space {
			if len(char) == 1 {
				if currentPos != len(command) {
					// runes := []rune(command)
					// runes[currentPos] = 
					// command = string(runes)
				}
				command = command + char
				maxRight++
				currentPos++
				fmt.Print(key.String())
			} else if char == "space" {
				command = command + " "
				maxRight++
				currentPos++
				fmt.Print(" ")
			}
		} else {
			if len(command) != 0 {
				command = 
				cursor.Left(1)
				fmt.Print("\033[K\r")
				maxRight--
				currentPos--
				fmt.Print(command)
			}
		}
		return false, nil
	})
}
