package main

import (
	"main/constants/menu"
	"main/lib"
	"main/terminal/list"
)

func main() {
	lib.INIT()
	list.List(menu.Menu)
}
