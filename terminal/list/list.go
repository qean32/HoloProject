package list

import (
	"fmt"
	"main/lib"
	"main/model"
	"main/terminal"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func List(list model.Options) {
	UnwrapList(list)
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.Enter {
			SelectItem()
			return true, nil
		}
		if key.Code == keys.Up {
			MoveUpList()
		}
		if key.Code == keys.Down {
			MoveDownList()
		}
		if key.Code == keys.Escape || key.Code == keys.CtrlC {
			lib.StopProcess()
		}
		return false, nil
	})
}

func UnwrapList(_list model.Options) {
	i := 0
	for message := range _list {
		fmt.Println(getStartChar(i == list.Position) + message)
		i++
	}
}

func getStartChar(isSelected bool) string {
	if isSelected {
		return "> "
	}

	return "< "
}

func SelectItem() {
}

func MoveUpList() {
	IncrementPosition()
	RefreshList()
}

func MoveDownList() {
	DecrimentPosition()
	RefreshList()
}

func RefreshList() {
	terminal.ClearLineByCount(len(list.Options))
	UnwrapList(list.Options)
}
