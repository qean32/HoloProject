package handler

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"atomicgo.dev/cursor"

	"main/callstack"
	"main/constants"
	"main/constants/literals"
	"main/constants/response"
	"main/lib/array"
	"main/lib/filter"
	"main/lib/low"
	"main/model"
	"main/terminal"
	"main/terminal/list"
	"main/terminal/questionnaire"
)

func _response(e model.Event) {
	terminal.Println(e.Payload)
}

func inwork(e model.Event) {
	terminal.Println("в разработке!")
}

func encrypt(e model.Event) {
	low.PushToFile(constants.PATH_DATA, fmt.Sprintf("%s %s", e.KeyWord, e.Payload))
	low.TMP_DATA = append(low.TMP_DATA, []string{e.KeyWord, e.Payload})
}

func decrypt(e model.Event) {
	index := slices.IndexFunc(low.TMP_DATA, func(item []string) bool {
		return item[0] == e.KeyWord
	})
	if index == -1 {
		callstack.PushCallStack(response.UndefinedKeyword())
		return
	}
	terminal.Print(low.TMP_DATA[index][1])
}

func generateKey(e model.Event) {}

func clearLog(e model.Event) {
	low.ClearFile(constants.PATH_LOG)
}

func generateMasterKey(e model.Event) {}

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
	low.CreateFile(constants.Cmd+e.KeyWord+".bat", e.Payload)
}

func runCommand(e model.Event) {
	low.RUN_CMD(constants.Root + constants.Cmd + e.KeyWord)
}

func runMultipleCommand(e model.Event) {
	index := FindCommand(e.KeyWord)
	if index == -1 {
		callstack.PushCallStack(response.UndefinedKeyword())
		return
	}
	for _, cmd := range strings.Split(low.TMP_COMMANDS[index][1], ";") {
		low.RUN_CMD(cmd)
	}
}

func listCommand() {
	terminal.Print(strings.Join(low.ReadFile(constants.PATH_COMMAND), "\n~ "))
}

func removeCommand(e model.Event) {
	if FindCommand(e.KeyWord) == -1 {
		callstack.PushCallStack(response.UndefinedKeyword())
		return
	}
	filtered := filter.FILTER(low.TMP_COMMANDS, func(item []string) bool {
		return item[0] != e.KeyWord
	})
	low.TMP_COMMANDS = filtered
	low.WriteFile(strings.Join(array.MatrixToArrayString(filtered), "\n"), constants.PATH_COMMAND)
}

func openLog() {
	low.RUN_CMD(constants.Root + constants.PATH_LOG)
}

func runMenu(e model.Event) {
	list.List(Menu, literals.Titles.Menu)
}

func menuCommandList() {
	commands, _ := low.ListFilesByExt(constants.Root+constants.Cmd, constants.Bat)

	list.List(
		array.Map(commands, func(value string) model.Option {
			return model.Option{
				Message: value[:len(value)-4],
				Event:   model.Event{Key: literals.COMMANDLIST.RUNCOMMAND, KeyWord: value},
			}
		}),
		literals.Titles.ListCommand,
	)
}

func menuQuestionnaireAddCommand(e model.Event) {
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
	}, literals.Titles.EnterCommand)

	declare(model.Event{
		Key:      literals.COMMANDLIST.DECLARE,
		KeyWord:  q["KeyWord"],
		Payload:  q["Payload"],
		DateTime: low.CurrentTime(),
	})
}

func FindCommand(keyword string) int {
	kw := strings.TrimSpace(keyword)
	return slices.IndexFunc(low.TMP_COMMANDS, func(item []string) bool {
		return item[0] == kw
	})
}

func IgnoreEvent(fn func()) model.EventFunction {
	return func(model.Event) { fn() }
}
