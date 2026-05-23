package main

import (
	"fmt"
	"os"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func main() {
	// keyboard.Listen(func(key keys.Key) (stop bool, err error) {
	// 	switch key.Code {
	// 	case keys.CtrlC, keys.Escape:
	// 		return true, nil // Return true to stop listener
	// 	case keys.RuneKey: // Check if key is a rune key (a, b, c, 1, 2, 3, ...)
	// 		if key.String() == "q" { // Check if key is "q"
	// 			fmt.Println("\rQuitting application")
	// 			os.Exit(0) // Exit application
	// 		}
	// 		fmt.Printf("\rYou pressed the rune key: %s\n", key)
	// 	default:
	// 		fmt.Printf("\rYou pressed: %s\n", key)
	// 	}

	// 	return false, nil // Return false to continue listening
	// })

	go func() {
		// keyboard.SimulateKeyPress(keys.Enter)          // Simulate key press for Enter
		// keyboard.SimulateKeyPress(keys.CtrlShiftRight) // Simulate key press for Ctrl+Shift+Right
		// keyboard.SimulateKeyPress('x')                 // Simulate key press for a single rune
		// keyboard.SimulateKeyPress('x', keys.Down, 'a') // Simulate key presses for multiple inputs

		// keyboard.SimulateKeyPress(keys.Escape) // Simulate key press for Escape, which quits the program
	}()

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.Escape || key.Code == keys.CtrlC {
			os.Exit(0) // Exit program on Escape
		}
		if key.Code == keys.Backspace {
			fmt.Print("\r")
			return false, nil
		}
		if key.Code == keys.Enter {
			fmt.Println("\033[K\r")
			return false, nil
		}

		fmt.Print("\033[33m" + key.String()) // Print every key press
		return false, nil                    // Return false to continue listening
	})
}
