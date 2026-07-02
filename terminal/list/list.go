package list

import (
	"fmt"
	"main/lib"
	"main/model"
	"main/terminal"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func List(list []model.Option) {
	_setOptions(list)
	UnwrapList(list)
	terminal.JumpToUpCount(len(list))
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

func UnwrapList(options []model.Option) {
	for i, item := range options {
		fmt.Println("\r" + getStartChar(i == list.Position) + item.Message + "\r")
	}
}

func getStartChar(isSelected bool) string {
	if isSelected {
		return "● "
	}

	return "○ "
}

func SelectItem() {
}

func MoveUpList() {
	if DecrimentPosition() {
		RefreshList()
	}
}

func MoveDownList() {
	if IncrementPosition() {
		RefreshList()
	}
}

func RefreshList() {
	terminal.JumpDownCount(list.Length - list.Position)
	terminal.ClearLineByCount(list.Length)
	UnwrapList(list.Options)
	terminal.JumpToUpCount(list.Length - list.Position)
}
