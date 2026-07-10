package main

import (
	"main/constants"
	"main/lib"
	"main/terminal/list"
)

func main() {
	lib.INIT()
	list.List(constants.Menu)
}
