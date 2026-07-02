package main

import (
	"fmt"
	"main/model"
	"main/terminal/list"
	// "main/terminal"
)

func main() {
	fmt.Println("----------------- start -----------------")
	// terminal.Field()
	list.List([]model.Option{
		{
			Message: "1",
			Command: "zxc1",
		},
		{
			Message: "2",
			Command: "zxc1",
		},
		{
			Message: "3",
			Command: "zxc1",
		},
		{
			Message: "4",
			Command: "zxc1",
		},
	})
}
