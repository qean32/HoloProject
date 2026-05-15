package ui

import (
	"main/deprecated/constants"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RENDER_CENTER(output string) {
	Display.AddItem(tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText(output), len(strings.Split(output, "\n")), 1, false)
}

func RENDER_RESPONSE(response string) {
	Display.AddItem(tview.NewTextView().SetText("> "+response), 1, 1, false)
	RENDER_INPUT()
}

var app = tview.NewApplication()

func RENDER_INPUT() {
	input := Get_UNPUT()
	app.SetFocus(input)
	input.SetDoneFunc(func(key tcell.Key) {
		RENDER_RESPONSE(input.GetText())
	})
	Display.AddItem(input, 1, 1, true)
}

func RENDER_UI() {
	RENDER_CENTER(constants.PROJECT_INIT)
	RENDER_INPUT()

	if err := app.SetRoot(Display, true).Run(); err != nil {
		panic(err)
	}
}
