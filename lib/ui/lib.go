package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func RENDER_CENTER(output string) {
	Display.AddItem(tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText(output), 1, 1, false)
}

func RENDER_RESPONSE(response string) {
	Display.AddItem(tview.NewTextView().SetText("> "+response), 1, 1, false)
	RENDER_INPUT()
}

func RENDER_INPUT() {
	input := Get_UNPUT()
	input.SetDoneFunc(func(key tcell.Key) {
		Display.RemoveItem(input)
	})
	Display.AddItem(input, 1, 1, true)
}

func RENDER_UI() {
	app := tview.NewApplication()
	RENDER_CENTER("zxc")
	RENDER_INPUT()

	if err := app.SetRoot(Display, true).Run(); err != nil {
		panic(err)
	}
}
