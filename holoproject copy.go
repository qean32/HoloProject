package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"

// 	"atomicgo.dev/cursor"
// 	"atomicgo.dev/keyboard"
// 	"atomicgo.dev/keyboard/keys"
// )

// type POINT struct {
// 	x int
// 	y int
// }

// type FIELDTYPE struct {
// 	position      int
// 	positionRange int
// 	field         string
// }

// var field = FIELDTYPE{
// 	position:      0,
// 	positionRange: 0,
// 	field:         "",
// }

// func AddChar(char string) {
// 	if field.position != len(field.field) {
// 		chars := strings.Split(field.field, "")
// 		chars[field.position] = char
// 		field.field = strings.Join(chars, "")
// 	}
// 	field.field = field.field + char
// 	ChangePositionCursor(1)
// 	fmt.Print(char)
// }

// func ChangePositionCursor(operation int) {
// 	if operation != -1 && operation != 1 {
// 		return
// 	}

// 	field.position++
// 	field.positionRange++
// }

// func AddSpace() {
// 	field.field = field.field + " "
// 	ChangePositionCursor(1)
// 	fmt.Print(" ")
// }

// func RemoveChar() {
// 	if len(field.field) != 0 {
// 		field.field = field.field[:len(field.field)-1]
// 		cursor.Left(1)
// 		fmt.Print("\033[K\r")
// 		ChangePositionCursor(-1)
// 		fmt.Print(field.field)
// 	}
// }

// func EnterCommand() {
// 	fmt.Println("")
// 	cursor.StartOfLineDown(1)
// 	fmt.Println("обработка")
// 	cursor.StartOfLineDown(1)
// 	ResetField()
// }

// func MoveCursorLeft() {
// 	cursor.Left(1)
// 	if field.position > 0 {
// 		ChangePositionCursor(-1)
// 	}
// }

// func MoveCursorRight() {
// 	if field.position < field.positionRange {
// 		cursor.Right(1)
// 		ChangePositionCursor(1)
// 	}
// }

// func StopProccess() {
// 	os.Exit(0)
// }

// func ResetField() {
// 	field.field = ""
// 	field.position = 0
// 	field.positionRange = 0
// }

// func main() {
// 	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
// 		char := key.String()
// 		if key.Code != keys.Backspace {
// 			if len(char) == 1 {
// 				AddChar(char)
// 			} else if char == "space" {
// 				AddSpace()
// 			}
// 		} else {
// 			RemoveChar()
// 		}
// 		if key.Code == keys.Enter {
// 			EnterCommand()
// 			return false, nil
// 		}
// 		if key.Code == keys.Right {
// 			MoveCursorRight()
// 		}
// 		if key.Code == keys.Left {
// 			MoveCursorLeft()
// 		}
// 		if key.Code == keys.Escape || key.Code == keys.CtrlC {
// 			StopProccess()
// 		}
// 		return false, nil
// 	})
// }
