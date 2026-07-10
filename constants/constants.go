package constants

import (
	"main/model"
	"os"
)

const MODE = "dev"

const UNDEFINED_COMMAND = "Undefined command"
const SYNTAX_ERROR = "Syntax error"
const STOP_COMMAND = "Command stop"
const UNDEFINED_KEYWORD = "Undefined word key"

const PROJECT_NAME = "holoproject"

const PATH_COMMAND = "/command.asc"
const PATH_LOG = "/log.asc"
const PATH_DATA = "/data.asc"
const PATH_SETTINGS = "/settings.asc"
const PATH_NOTES = "/notes.asc"

var Root = "./private"

func INIT_ROOT() {
	if MODE == "dev" {
		Root = "./private"
	} else {
		os.Mkdir(os.TempDir()+`\`+PROJECT_NAME, 0755)
		Root = os.TempDir() + `\` + PROJECT_NAME
	}
}

var BinaryPROJECT_INIT = `
%v 01101000 01101111 01101100 01101111 01110000 01110010 01101111 01101010 01100101 01100011 01110100
`

var PROJECT_INIT = `
%v  __              ___                                                    __            ___     
%v /\ \            /\_ \                                 __               /\ \__       /'___'\   
%v \ \ \___     ___\//\ \     ___   _____   _ __   ___  /\_\     __    ___\ \ ,_\     /\_\ /\ \  
%v  \ \  _ '\  / __'\\ \ \   / __'\/\ '__'\/\''__\/ __'\\/\ \  /'__'\ /'___\ \ \/     \/_/// /__ 
%v   \ \ \ \ \/\ \L\ \\_\ \_/\ \L\ \ \ \L\ \ \ \//\ \L\ \\ \ \/\  __//\ \__/\ \ \_       // /_\ \
%v    \ \_\ \_\ \____//\____\ \____/\ \ ,__/\ \_\\ \____/_\ \ \ \____\ \____\\ \__\     /\______/
%v     \/_/\/_/\/___/ \/____/\/___/  \ \ \/  \/_/ \/___//\ \_\ \/____/\/____/ \/__/     \/_____/ 
%v                                    \ \_\             \ \____/                                 
%v                                     \/_/              \/___/                                  
`

// ASCII
// "\033[31m" red color text
// https://translated.turbopages.org/proxy_u/en-ru.ru.6c604ba4-6a11ef9b-b4e5d92b-74722d776562/https/student.cs.uwaterloo.ca/~cs452/terminal.html
// https://www.asciiart.eu/text-to-ascii-art respect

var HelpMessage = `
  _           _                       _           _     _          _        
 | |__   ___ | | ___  _ __  _ __ ___ (_) ___  ___| |_  | |__   ___| |_ __   
 | '_ \ / _ \| |/ _ \| '_ \| '__/ _ \| |/ _ \/ __| __| | '_ \ / _ \ | '_ \  
 | | | | (_) | | (_) | |_) | | | (_) | |  __/ (__| |_  | | | |  __/ | |_) | 
 |_| |_|\___/|_|\___/| .__/|_|  \___// |\___|\___|\__| |_| |_|\___|_| .__/  
                     |_|           |__/                             |_|     

| cripto       |  Шифрование строки               | cripto key_word key_password { payload }  |
| ecripto      |  Дешифровка строки               | ecripto key_word key_password             |
|              |                                  |                                           |
| gmaster      |  Генерация мастер-ключа          |                                           |
| gkey         |  Генерация ключа для шифра       |                                           |
|              |                                  |                                           |
| clog         |  Очистка логов                   |                                           |
| drop         |  Удалить данные приложения       |                                           |
| stop         |  Остановить приложение           |                                           |
| help         |  Список команд                   |                                           |
|              |                                  |                                           |
| note         |  Открыть запись                  |                                           |
| dnote        |  Создание записи                 | dnote key_word key_password               |
| notes        |  Получить список записей         |                                           |
|              |                                  |                                           |
| declare      |  Добавить команду                | declare key_word { payload }              |
| commands     |  Просмотреть список команд       |                                           |
| run          |  Запустить команду               | run key_word                              |
| runm         |  Запустить множественую команду  | runm key_word                             |
| rmc          |  Удалить команду                 | rmc key_word                              |
`

var ACCESS_VARIANTS = []string{"yes", "y", "yea"}

const NEXT_LINE = "\n"
const OUTPUT_MESSAGE = "output "

const CHAR_SELECTED_ITEM = "● "
const CHAR_UN_SELECTED_ITEM = "○ "

var (
	StyleError   = []int{31, 1, 4}
	StyleSuccess = []int{32, 1}
	StyleWarning = []int{33}
	StyleInfo    = []int{44, 37}
)

var AcceptableYeaOrNot []string = []string{"yes", "no", "y", "n", "yea", "not"}

var Menu = []model.Option{
	{
		Message: "Запустить команду",
		Command: func() {},
	},
	{
		Message: "Добавить команду",
		Command: func() {},
	},
	{
		Message: "Очистка логов",
		Command: func() {},
	},
	{
		Message: "Команды",
		Command: func() {},
	},
	{
		Message: "Выход",
		Command: func() {},
	},
}
