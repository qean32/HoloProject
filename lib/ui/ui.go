package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func callback(key tcell.Key) {
	RENDER_RESPONSE("zxc")
}

func Get_UNPUT() *tview.InputField {
	return tview.NewInputField().
		SetLabel("> ").
		SetFieldBackgroundColor(tcell.ColorBlack).
		SetFieldTextColor(tcell.ColorWhite)
}

var Display = tview.NewFlex().SetDirection(tview.FlexRow)
