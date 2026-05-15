package constants

import "os"

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

var PROJECT_INIT = `
                                                          ⠀⠀⣀⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
                                                          ⠀⠘⡏⠉⠳⢦⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⠶⠶⠤⣄⠀⠀⠀⠀⠀⠀
 _           _                       _           _        ⠀⠀⢹⣀⣀⠀⠈⠳⣤⡀⠀⠀⠀⠀⠀⠀⠀⢠⡞⠁⠐⣶⠄⢼⣧⠀⠀⠀⠀⠀
| |__   ___ | | ___  _ __  _ __ ___ (_) ___  ___| |_      ⠀⠀⠀⣇⠀⠧⠀⠀⠸⣿⣷⡦⣄⠀⠀⠀⢠⣿⡤⠀⡇⠘⡆⠀⢻⡄⠀⠀⠀⠀
| '_ \ / _ \| |/ _ \| '_ \| '__/ _ \| |/ _ \/ __| __|     ⠀⠀⠀⠸⡀⠐⣦⣤⠀⠙⠹⣕⢼⣷⣄⣠⡿⠛⠉⠉⠒⣄⣿⠀⠘⣧⠀⠀⠀⠀
| | | | (_) | | (_) | |_) | | | (_) | |  __/ (__| |_      ⠀⠀⠀⠀⢧⠚⢻⣿⡷⢦⠀⢿⣼⣏⣺⠟⡇⠀⠀⢀⣶⣶⣾⠀⠀⣿⣧⡀⠀⠀
|_| |_|\___/|_|\___/| .__/|_|  \___// |\___|\___|\__|     ⠀⠀⠀⠀⢀⠧⡀⠸⣗⠈⠷⠘⠏⢻⠇⠀⠀⣿⠆⣤⣾⣿⡿⠃⢸⣿⣷⢷⠀⠀
                    |_|           |__/                    ⠀⠀⠀⠀⠀⠹⡟⢲⠎⠃⠀⠀⠀⠀⠀⠀⠺⠁⢰⢿⡿⠋⠃⠠⡽⣷⡈⢻⠀⠀
                                                          ⠀⠀⠀⠀⠀⢠⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠐⠋⠀⠀⠀⠀⢀⣷⡟⠃⠸⡆⠀
                                                          ⠀⠀⠀⠀⢀⡾⠁⠀⠀⠀⠀⠀⢀⠀⠀⠀⠀⠀⠀⠘⢶⣷⣲⡛⠉⠰⠦⢄⣹⡄
                                                          ⠀⠀⠀⠀⠈⠙⠳⣆⠀⠀⠀⠀⣿⣷⡦⢤⣀⣰⣦⡀⠀⠓⣤⡘⣕⢤⡀⠀⢹⡇
                                                          ⠀⠀⠀⠀⠀⠀⣟⡏⠀⠀⠀⣼⠿⠛⠛⠉⠀⣠⡙⠫⣤⣀⠀⢱⣌⢧⠙⢦⠸⡇
                                                          ⠀⠀⠀⠀⠀⢀⡟⠀⠀⠰⡟⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠉⠻⣄⢿⣞⢇⠀⠀⣷
                                                          ⠀⠀⠀⠀⣠⠏⠀⠀⠀⢀⣷⣿⣦⣠⠤⠈⠡⠬⠤⠀⠀⢄⡀⠈⢳⡿⣯⢳⡄⣿
                                                          ⠀⠀⢀⣾⣥⣄⣀⣠⣤⣿⣿⡿⠙⢃⡤⠶⠀⠀⠀⠀⠀⢠⣧⠀⠀⠁⣿⠀⢹⡟
                                                          ⠀⠀⠈⢻⣿⣿⠏⠀⠀⠀⣈⣤⠾⠗⠒⠒⠉⠉⠉⠑⠒⠠⢿⡀⠀⠀⢻⢢⡿⠁
                                                          ⠀⠀⠀⠀⠉⠙⠓⠲⠾⠿⠋⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠑⠲⣼⠟⠁⠀
`

// ASCII
// https://www.asciiart.eu/text-to-ascii-art respect

var HelpMessage = `
  _           _                       _           _     _          _        
 | |__   ___ | | ___  _ __  _ __ ___ (_) ___  ___| |_  | |__   ___| |_ __   
 | '_ \ / _ \| |/ _ \| '_ \| '__/ _ \| |/ _ \/ __| __| | '_ \ / _ \ | '_ \  
 | | | | (_) | | (_) | |_) | | | (_) | |  __/ (__| |_  | | | |  __/ | |_) | 
 |_| |_|\___/|_|\___/| .__/|_|  \___// |\___|\___|\__| |_| |_|\___|_| .__/  
                     |_|           |__/                             |_|     

|  -cripto     |  Шифрование строки               |  cripto key_word key_password { payload }  |
|  -ecripto    |  Дешифровка строки               |  ecripto key_word key_password             |
|              |                                  |                                            |
|  -gmaster    |  Генерация мастер-ключа          |                                            |
|  -gkey       |  Генерация ключа для шифра       |                                            |
|              |                                  |                                            |
|  -clog       |  Очистка логов                   |                                            |
|  -drop       |  Удалить данные приложения       |                                            |
|  -stop       |  Остановить приложение           |                                            |
|  -help       |  Список команд                   |                                            |
|              |                                  |                                            |
|  -note       |  Открыть запись                  |                                            |
|  -dnote      |  Создание записи                 |  dnote key_word key_password               |
|  -notes      |  Получить список записей         |                                            |
|              |                                  |                                            |
|  -declare    |  Добавить команду                |  declare key_word { payload }              |
|  -commands   |  Просмотреть список команд       |                                            |
|  -run        |  Запустить команду               |  run key_word                              |
|  -runm       |  Запустить множественую команду  |  runm key_word                             |
|  -rmc        |  Удалить команду                 |  rmc key_word                              |
`

var ACCESS_VARIANTS = []string{"yes", "y", "yea"}

const NEXT_LINE = "\n"
const OUTPUT_MESSAGE = "~ OUTPUT ~ "
