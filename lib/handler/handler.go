package handler

import (
	"fmt"
	"os"

	"atomicgo.dev/cursor"

	"main/constants"
	"main/constants/literals"
	"main/lib/array"
	"main/lib/death"
	"main/model"
	"main/terminal"
	"main/terminal/list"
	"main/terminal/questionnaire"
)

func _response(e model.Event) {
	terminal.PrintResponse(e.Payload)
}

func inwork(e model.Event) {
	terminal.Println("в разработке!")
}

func clearLog(e model.Event) {
	death.ClearFile(literals.Path.PathLog)
}

func drop(e model.Event) {
	os.RemoveAll(constants.Root)
	os.Mkdir(constants.Root, 0755)
}

func help(e model.Event) {
	terminal.DownAndStart()
	cursor.StartOfLine()
	terminal.PrintASCIICenter(constants.HelpMessage, "")
}

func declare(e model.Event) {
	death.CreateFile(constants.Cmd+e.KeyWord+".bat", e.Payload)
}

func runCmd(e model.Event) {
	death.RunCmd(constants.Root + constants.Cmd + e.KeyWord)
}

func removeCmd(e model.Event) {
	err := os.Remove("file.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Файл удалён")
}

func runMenu(e model.Event) {
	list.List(Menu, literals.Titles.Menu)
}

func cmdList() {
	commands, _ := death.ListFilesByExt(constants.Root+constants.Cmd, literals.Extension.Bat)

	list.List(
		array.Map(commands, func(value string) model.Option {
			return model.Option{
				Message: value[:len(value)-4],
				Event:   model.Event{Key: literals.EventList.RUNCMD, KeyWord: value},
			}
		}),
		literals.Titles.ListCmd,
	)
}

func questionnaireAddCommand(e model.Event) {
	q := questionnaire.Questionnaire([]model.Question{
		{
			Message: "Ключ",
			Key:     "KeyWord",
			Callback: func(res string) bool {
				return true
			},
		},
		{
			Message: "Команда",
			Key:     "Payload",
			Callback: func(res string) bool {
				return true
			},
		},
	}, literals.Titles.EnterCmd)

	declare(model.Event{
		Key:      literals.EventList.DECLARE,
		KeyWord:  q["KeyWord"],
		Payload:  q["Payload"],
		DateTime: death.CurrentTime(),
	})
}

func encrypt(e model.Event)           {}
func decrypt(e model.Event)           {}
func generateKey(e model.Event)       {}
func generateMasterKey(e model.Event) {}

func openLog() {
	death.RunCmd(constants.Root + literals.Path.PathLog)
}
