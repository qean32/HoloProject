package main

import (
	"fmt"
	"main/model"
	"main/terminal/list"
)

func main() {
	fmt.Println("----------------- start -----------------")
	// terminal.Field()
	list.List([]model.Option{
		{
			Message: "1",
			Command: func() { fmt.Println("Command start") },
		},
		{
			Message: "2",
			Command: func() { fmt.Println("Command start") },
		},
		{
			Message: "3",
			Command: func() { fmt.Println("Command start") },
		},
		{
			Message: "4",
			Command: func() { fmt.Println("Command start") },
		},
	})
}
