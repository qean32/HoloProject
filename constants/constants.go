package constants

import "os"

var MODE = "prod"

const UNDEFINED_COMMAND = "$ Undefined command"
const ERROR_COMMAND = "$ Syntax error"
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
      )                                                      
   ( /(       (                                           )  
   )\())      )\               (         (     (       ( /(  
  ((_)\   (  ((_) (     '  )   )(    (   )\   ))\  (   )\()) 
   _((_)  )\  _   )\    /(/(  (()\   )\ ((_) /((_) )\ (_))/  
  | || | ((_)| | ((_)  ((_)_\  ((_) ((_)  ! (_))  ((_)| |_   
  | __ |/ _ \| |/ _ \  | '_ \)| '_|/ _ \ | |/ -_)/ _| |  _|  
  |_||_|\___/|_|\___/  | .__/ |_|  \___/_/ |\___|\__|  \__|  
                       |_|             |__/                  
`

// ASCII

var HelpMessage = `
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
~                                  Holo Project                                       ~
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
~                                                                                     ~
~  - - cripto     : Шифрование строки               |  "cripto key_word {payload}"    ~
~  - - ecripto    : Дешифровка строки               |  "ecripto key_word"             ~
~                                                                                     ~
~  - - stop       : Остановить приложение           |  "stop"                         ~
~  - - drop       : Удалить данные приложения       |  "drop"                         ~
~  - - gmasterkey : Генерация мастер-ключа          |  "gmasterkey"                   ~
~                                                                                     ~
~  - - gkey       : Генерация ключа для шифра       |  "gkey"                         ~
~  - - clog       : Очистка логов                   |  "clog"                         ~
~                                                                                     ~
~  - - run        : Запустить команду               |  "run key_word"                 ~
~  - - runm       : Запустить мультиязычную команду |  "runm key_word"                ~
~  - - place      : Добавить команду                |  "place key_word {payload}"     ~
~  - - commands   : Просмотреть список команд       |  "commands"                     ~
~  - - rmc        : Удалить команду                 |  "rmc key_word"                 ~
~                                                                                     ~
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
~                           Используйте команды для                                   ~
~                  управления вашим приложением и системой.                           ~
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
`
