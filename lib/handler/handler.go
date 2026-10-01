package handler

import (
	"os"

	"atomicgo.dev/cursor"

	"main/callstack"
	"main/callstack/event"
	"main/callstack/manual"
	"main/constants"
	"main/constants/literals"
	"main/lib/array"
	"main/lib/death"
	"main/model"
	"main/terminal"
	"main/terminal/list"
	"main/terminal/qNext"
	"main/terminal/questionnaire"
)

func reset() {
	death.ClearTerminal()
	terminal.RenderBaner()
	manual.Manual()
}

func _qnext() {
	qnext.QNext()
}

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

func bat(e model.Event) {
	death.CreateFile(constants.Cmd+e.KeyWord+".bat", e.Payload)
}

func runCmd(e model.Event) {
	death.RunCmd(constants.Root + constants.Cmd + e.KeyWord)
}

func removeCmd(e model.Event) {

	err := os.Remove(constants.Root + constants.Cmd + e.KeyWord)
	if err != nil {
		event.Response("Ошибка: ", err)
		return
	}
}

func questionnaireAccess(e model.Event) {
	q := questionnaire.Questionnaire(model.QuestionnairePayload{
		Questions: []model.Question{
			{
				Message: "",
				Key:     "Response",
				Callback: func(res string) bool {
					return true
				},
			},
		},
		Title: literals.Titles.Access,
	})

	if questionnaire.IsAcceptableYea(q["Response"]) {
		callstack.PushCallStack(*e.SubEvent)
	} else {
		event.Response("Процесс остановлен")
	}
}

func runMenu(e model.Event) {
	list.List(model.ListPayload{
		Options: Menu,
		Title:   literals.Titles.Menu,
	})
}

func batList(eventKey string) {
	commands, _ := death.ListFilesByExt(constants.Root+constants.Cmd, literals.Extension.Bat)

	list.List(
		model.ListPayload{
			Options: array.Map(commands, func(value string) model.Option {
				return model.Option{
					Message: value[:len(value)-4],
					Event:   model.Event{Key: eventKey, KeyWord: value},
				}
			}),
			Title: literals.Titles.ListCmd,
		},
	)
}

func removeBatList() {
	commands, _ := death.ListFilesByExt(constants.Root+constants.Cmd, literals.Extension.Bat)

	list.List(
		model.ListPayload{
			Options: array.Map(commands, func(value string) model.Option {
				return model.Option{
					Message: value[:len(value)-4],
					Event: model.Event{
						Key:     literals.EventList.REMOVEBAT,
						KeyWord: value,
						SubEvent: &model.Event{
							Key:     literals.EventList.RRBAT,
							KeyWord: value,
						},
					},
				}
			}),
			Title: literals.Titles.ListCmd,
		},
	)
}

func questionnaireAddCommand(e model.Event) {
	q := questionnaire.Questionnaire(model.QuestionnairePayload{
		Questions: []model.Question{
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
		},
		Title: literals.Titles.EnterCmd,
	})

	bat(model.Event{
		Key:      literals.EventList.BAT,
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
