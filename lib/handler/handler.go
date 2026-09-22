package handler

import (
	"fmt"
	"main/constants"
	"main/constants/literals"
	"main/lib/array"
	"main/lib/filter"
	"main/lib/low"
	"main/model"
	"main/terminal"

	"main/terminal/list"
	"main/terminal/questionnaire"
	"os"
	"slices"
	"strings"

	"atomicgo.dev/cursor"
)

func Inwork(e model.Event) {
	terminal.Outputln("в разработке!")
}

func Encrypt(e model.Event) {
	low.PushToFile(constants.PATH_DATA, fmt.Sprintf("%s %s", e.KeyWord, e.Payload))
	low.TMP_DATA = append(low.TMP_DATA, []string{e.KeyWord, e.Payload})
}

func Decrypt(e model.Event) {
	index := slices.IndexFunc(low.TMP_DATA, func(item []string) bool {
		return item[0] == e.KeyWord
	})

	if index != -1 {
		terminal.Output(low.TMP_DATA[index][1])
	} else {
		terminal.Output(constants.UNDEFINED_KEYWORD)
	}
}

func GenerateKey(e model.Event) {}

func ClearLog(e model.Event) {
	low.ClearFile(constants.PATH_LOG)
}

func GenerateMasterKey(e model.Event) {
	low.CreateFile(constants.PATH_LOG)
	low.CreateFile(constants.PATH_COMMAND)
	low.CreateFile(constants.PATH_DATA)
}

func Drop(e model.Event) {
	os.RemoveAll(constants.Root)
	os.Mkdir(constants.Root, 0755)
}

func Stop(e model.Event) {
	low.StopProcess()
}

func Help(e model.Event) {
	terminal.DownAndStart()
	cursor.StartOfLine()
	terminal.OutputASCII_CENTER(constants.HelpMessage, "")
}

func Declare(e model.Event) {
	low.PushToFile(constants.PATH_COMMAND, fmt.Sprintf("%s %s", e.KeyWord, e.Payload))
	low.TMP_COMMANDS = append(low.TMP_COMMANDS, []string{e.KeyWord, e.Payload})
}

func RunCommand(e model.Event) {
	index := slices.IndexFunc(low.TMP_COMMANDS, func(item []string) bool {
		return item[0] == strings.TrimSpace(e.KeyWord)
	})

	if index != -1 {
		low.RUN_CMD(low.TMP_COMMANDS[index][1])
	} else {
		terminal.Output(constants.UNDEFINED_KEYWORD)
	}
}

func RunMultipleCommands(e model.Event) {
	index := slices.IndexFunc(low.TMP_COMMANDS, func(item []string) bool {
		return item[0] == strings.TrimSpace(e.KeyWord)
	})

	if index != -1 {
		commands := strings.Split(low.TMP_COMMANDS[index][1], ";")

		for i := 0; i < len(commands); i++ {
			low.RUN_CMD(commands[i])
		}
	} else {
		terminal.Output(constants.UNDEFINED_KEYWORD)
	}
}

func ListCommands(e model.Event) {
	terminal.Output(strings.Join(low.ReadFile(constants.PATH_COMMAND), "\n~ "))
}

func RemoveCommand(e model.Event) {
	if slices.IndexFunc(low.TMP_COMMANDS, func(item []string) bool {
		return item[0] == e.KeyWord
	}) != -1 {
		filtered := filter.FILTER(low.TMP_COMMANDS, func(item []string) bool { return item[0] != e.KeyWord })
		low.TMP_COMMANDS = filtered
		low.WriteFile(strings.Join(array.MatrixToArrayString(filtered), "\n"), constants.PATH_COMMAND)
	} else {
		terminal.Output(constants.UNDEFINED_KEYWORD)
	}
}

func Notes(e model.Event) {}

func RunMenu(e model.Event) {
	list.List(Menu)
}

func Note(e model.Event) {}

func DeleteNote(e model.Event) {}

func Menu_runCommandsList() {
	list.List(
		array.Map(low.TMP_COMMANDS, func(value []string) model.Option {
			return model.Option{
				Message: value[0],
				Event:   model.Event{Key: literals.COMMANDS_LIST.RUN_COMMAND, KeyWord: value[0]},
			}
		}),
	)
}

func Menu_runQuestionnaireAddCommand(e model.Event) {
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
	})

	Declare(model.Event{
		Key:      literals.COMMANDS_LIST.DECLARE,
		KeyWord:  q["KeyWord"],
		Payload:  q["Payload"],
		DateTime: low.CurrentTime(),
	})
}
