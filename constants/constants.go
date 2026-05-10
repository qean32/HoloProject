package constants

import "os"

var MODE = "prod"

const UNDEFINED_COMMAND = "$ Undefined command"
const SYNTAX_ERROR = "$ Syntax error"
const STOP_COMMAND = "$ Command stop"
const UNDEFINED_WORD_KEY = "$ Undefined word key"
const PROJECT_NAME = "holoproject"

const COMMAND_PATH = "/command.asc"
const LOG_PATH = "/log.asc"
const DATA_PATH = "/data.asc"

var Root = "./private"

func INIT_ROOT() {
	if MODE == "dev" {
		Root = "./private"
	} else {
		os.Mkdir(os.TempDir()+`\`+PROJECT_NAME, 0755)
		Root = os.TempDir() + `\` + PROJECT_NAME
	}
}

var PROJECT_INIT = `
    __  ______  __    ____      ____  ____  ____      ____________________
   / / / / __ \/ /   / __ \    / __ \/ __ \/ __ \    / / ____/ ____/_  __/
  / /_/ / / / / /   / / / /   / /_/ / /_/ / / / /_  / / __/ / /     / /   
 / __  / /_/ / /___/ /_/ /   / ____/ _, _/ /_/ / /_/ / /___/ /___  / /    
/_/ /_/\____/_____/\____/   /_/   /_/ |_|\____/\____/_____/\____/ /_/             
`

// ASCII

var HelpMessage = `
    __  ______  __    ____      ____  ____  ____      ____________________   __  __________    ____ 
   / / / / __ \/ /   / __ \    / __ \/ __ \/ __ \    / / ____/ ____/_  __/  / / / / ____/ /   / __ \
  / /_/ / / / / /   / / / /   / /_/ / /_/ / / / /_  / / __/ / /     / /    / /_/ / __/ / /   / /_/ /
 / __  / /_/ / /___/ /_/ /   / ____/ _, _/ /_/ / /_/ / /___/ /___  / /    / __  / /___/ /___/ ____/ 
/_/ /_/\____/_____/\____/   /_/   /_/ |_|\____/\____/_____/\____/ /_/    /_/ /_/_____/_____/_/      

  -cripto     | Шифрование строки               |  "cripto key_word key_password {payload}"
  -ecripto    | Дешифровка строки               |  "ecripto key_word key_password"

  -run        | Запустить команду               |  "run key_word"
  -runm       | Запустить мультиязычную команду |  "runm key_word"
  -rmc        | Удалить команду                 |  "rmc key_word"
  -place      | Добавить команду                |  "place key_word {payload}"
  -commands   | Просмотреть список команд       |  "commands"

  -stop       | Остановить приложение           |  "stop"
  -drop       | Удалить данные приложения       |  "drop"
  -gmasterkey | Генерация мастер-ключа          |  "gmasterkey"

  -gkey       | Генерация ключа для шифра       |  "gkey"
  -clog       | Очистка логов                   |  "clog"
`

type CommandsType struct {
	CRIPTO     string
	ECRIPTO    string
	GKEY       string
	CLOG       string
	GMASTERKEY string
	DROP       string
	STOP       string
	HELP       string
	PLACE      string
	RUN        string
	RUNM       string
	COMM       string
	RMC        string
}

var COMMANDS = CommandsType{
	CRIPTO:     "cripto",
	ECRIPTO:    "ecripto",
	GKEY:       "gkey",
	CLOG:       "clog",
	GMASTERKEY: "gmasterkey",
	DROP:       "drop",
	STOP:       "stop",
	HELP:       "help",
	PLACE:      "place",
	RUN:        "run",
	RUNM:       "runm",
	COMM:       "comm",
	RMC:        "rmc",
}
