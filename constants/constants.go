package constants

import "os"

var MODE = "prod"

const UNDEFINED_COMMAND = "Undefined command"
const ERROR_COMMAND = "Syntax error"
const STOP_COMMAND = "Command stop"
const UNDEFINED_WORD_KEY = "Undefined word key"
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

var HelpMessage = `
+------------------------------------------------------------+
|                   Holo Project                             |
+------------------------------------------------------------+
|                                                            |
|  - - cripto     : Шифрование строки                        |
|  - - ecripto    : Дешифровка строки                        |
|                                                            |
|  - - stop       : Остановить приложение                    |
|  - - drop       : Удалить данные приложения                |
|  - - gmasterkey : Генерация мастер-ключа                   |
|                                                            |
|  - - gkey       : Генерация ключа для шифра                |
|  - - clog       : Очистка логов                            |
|                                                            |
|  - - run        : Запустить команду                        |
|  - - run-m      : Запустить мультиязычную команду          |
|  - - place      : Добавить команду                         |
|  - - commands   : Просмотреть список команд                |
|  - - rmc        : Удалить команду                          |
|                                                            |
+------------------------------------------------------------+
|               Используйте команды для                      |
|        управления вашим приложением и системой.            |
+------------------------------------------------------------+
`
