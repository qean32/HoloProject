package handler

import (
	test "main/__test__"
	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/lib/death"
	"main/lib/flow"
	"main/model"
)

var commands = []model.Command{
	{
		Name:    eventlist.BAT,
		Title:   "Добавить .bat",
		Usage:   "bat <имя> p/<команда>",
		Options: model.ParseOptions{KeyWord: true, Payload: true},
		Run:     bat,
		Success: true,
	},
	{
		Name:  eventlist.BATLIST,
		Title: "Список .bat",
		Usage: "list",
		Flow:  flow.RunBatList,
	},
	{
		Name:    eventlist.RUNBAT,
		Title:   "Выполнить .bat",
		Usage:   "run <имя>",
		Options: model.ParseOptions{KeyWord: true},
		Run:     runBat,
		Success: true,
	},
	{
		Name:    eventlist.REMOVEBAT,
		Title:   "Удалить .bat",
		Usage:   "rbat <имя>",
		Options: model.ParseOptions{KeyWord: true, GetSubEvent: true, SubEvent: eventlist.PUREREMOVEBAT},
		Flow:    flow.ConfirmRemove,
	},
	{
		Name:  eventlist.MENU,
		Title: "Меню",
		Usage: "menu",
		Flow:  flow.Menu,
	},
	{
		Name:  eventlist.HELP,
		Title: "Справка",
		Usage: "help",
		Run:   ignoreError(help),
	},
	{
		Name:  eventlist.CLEAR,
		Title: "Очистить экран",
		Usage: "clear",
		Run:   ignoreError(clearScreen),
	},
	{
		Name:    eventlist.CLEARLOG,
		Title:   "Очистить лог",
		Usage:   "clog",
		Run:     ignoreEvent(clearLog),
		Success: true,
	},
	{
		Name:  eventlist.OPENLOG,
		Title: "Открыть лог",
		Usage: "log",
		Run:   ignoreEvent(openLog),
	},
	{
		Name:  eventlist.LOGGER,
		Title: "Технический лог",
		Usage: "logger",
		Run:   ignoreError(toggleLogger),
	},
	{
		Name:    eventlist.DROP,
		Title:   "Удалить все данные",
		Usage:   "drop",
		Run:     ignoreEvent(drop),
		Success: true,
	},
	{
		Name:  eventlist.STOP,
		Title: "Выход",
		Usage: "stop",
		Run:   ignoreError(death.Exit),
	},
	// Криптокоманды пока не реализованы и не показываются в справке.
	{
		Name:    eventlist.CRYPTO,
		Title:   "Шифрование (в разработке)",
		Usage:   "crypto <имя> p/<данные> pass/<пароль>",
		Options: model.ParseOptions{KeyWord: true, Payload: true, Password: true},
		Run:     ignoreEvent(encrypt),
		Success: true,
		Hidden:  true,
	},
	{
		Name:    eventlist.DECRYPTO,
		Title:   "Расшифровка (в разработке)",
		Usage:   "ecrypto <имя> pass/<пароль>",
		Options: model.ParseOptions{KeyWord: true, Password: true},
		Run:     ignoreEvent(decrypt),
		Hidden:  true,
	},
	{
		Name:   eventlist.GENERATEKEY,
		Title:  "Сгенерировать ключ",
		Usage:  "gkey",
		Run:    ignoreEvent(generateKey),
		Hidden: true,
	},
	{
		Name:   eventlist.GENERATEMASTER,
		Title:  "Сгенерировать мастер-ключ",
		Usage:  "gmaster",
		Run:    ignoreEvent(generateMasterKey),
		Hidden: true,
	},
	{
		Name:    eventlist.TEST,
		Title:   "Самопроверка",
		Usage:   "test",
		Run:     ignoreError(test.Test),
		Success: true,
	},

	// Служебные команды и события меню — не показываем в help.
	{
		Name:   eventlist.MANUCREATEBAT,
		Title:  "Добавить .bat (диалог)",
		Flow:   flow.CreateBat,
		Hidden: true,
	},
	{
		Name:   eventlist.BATLISTREMOVE,
		Title:  "Удалить .bat (диалог)",
		Flow:   flow.RemoveBatList,
		Hidden: true,
	},
	{
		Name:    eventlist.PUREREMOVEBAT,
		Run:     removeCmd,
		Success: true,
		Hidden:  true,
	},
	{
		Name:    eventlist.INIT,
		Title:   "Инициализация",
		Usage:   "init",
		Run:     ignoreEvent(init_),
		Success: true,
		Hidden:  true,
	},
	{
		Name:   eventlist.INWORK,
		Run:    ignoreError(inwork),
		Hidden: true,
	},
	{
		Name:   eventlist.RESPONSE,
		Run:    _response,
		Hidden: true,
	},
	{
		Name:   eventlist.TESTVIEW,
		Run:    ignoreError(test.TestView),
		Hidden: true,
	},
	{
		Name:   eventlist.MENUCRYPTO,
		Run:    ignoreError(inwork),
		Hidden: true,
	},
}

// visible заполняется в init, чтобы help не создавал цикл инициализации
// с commands (commands содержит функцию help).
var visible []model.Command

func init() {
	for _, command := range commands {
		if !command.Hidden && command.Usage != "" {
			visible = append(visible, command)
		}
	}
}

var registry = buildRegistry(commands)

func buildRegistry(list []model.Command) map[string]model.Command {
	result := make(map[string]model.Command, len(list))
	for _, command := range list {
		if command.Name == "" || (command.Run == nil && command.Flow == nil) {
			panic("invalid command: " + command.Name)
		}
		if _, exists := result[command.Name]; exists {
			panic("duplicate command: " + command.Name)
		}
		result[command.Name] = command
	}
	return result
}

func ignoreError(fn func()) model.Handler {
	return func(_ model.Event) error {
		fn()
		return nil
	}
}

func ignoreEvent(fn func() error) model.Handler {
	return func(_ model.Event) error {
		return fn()
	}
}

func Dispatch(e model.Event) (model.Action, error) {
	command, ok := registry[e.Key]
	if !ok {
		return nil, errorlist.ErrUnknownCommand
	}
	if command.Flow != nil {
		return command.Flow(e)
	}
	return nil, command.Run(e)
}

func Success(name string) bool {
	command, ok := registry[name]
	return ok && command.Success
}

// VisibleCommands возвращает команды для справки и неинтерактивного списка.
func VisibleCommands() []model.Command {
	result := make([]model.Command, len(visible))
	copy(result, visible)
	return result
}
