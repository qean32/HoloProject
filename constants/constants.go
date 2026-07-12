package constants

import (
	"main/constants/literals"
	"os"
)

const MODE = "dev"
const PROJECT_NAME = "holo"

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
%v01101000 01101111 01101100 01101111 01110000 01110010 01101111 01101010 01100101 01100011 01110100
`

var PROJECT_INIT = `
%v __  __     ______     __         ______     ______   ______     ______       __     ______     ______     ______     
%v/\ \_\ \   /\  __ \   /\ \       /\  __ \   /\  == \ /\  == \   /\  __ \     /\ \   /\  ___\   /\  ___\   /\__  _\    
%v\ \  __ \  \ \ \/\ \  \ \ \____  \ \ \/\ \  \ \  _-/ \ \  __<   \ \ \/\ \   _\_\ \  \ \  __\   \ \ \____  \/_/\ \/    
%v \ \_\ \_\  \ \_____\  \ \_____\  \ \_____\  \ \_\    \ \_\ \_\  \ \_____\ /\_____\  \ \_____\  \ \_____\    \ \_\    
%v  \/_/\/_/   \/_____/   \/_____/   \/_____/   \/_/     \/_/ /_/   \/_____/ \/_____/   \/_____/   \/_____/     \/_/    
%v                                                                                                                      
`

// ASCII
// "\033[31m" red color text
// https://translated.turbopages.org/proxy_u/en-ru.ru.6c604ba4-6a11ef9b-b4e5d92b-74722d776562/https/student.cs.uwaterloo.ca/~cs452/terminal.html
// https://www.asciiart.eu/text-to-ascii-art respect
// https://patorjk.com/software/taag/#p=testall&f=Broadway&t=holoproject+2&x=none&v=4&h=4&w=80&we=false

var HelpMessage = `
%v  __  __     ______     __         ______     ______   ______     ______       __     ______     ______     ______      __  __     ______     __         ______  
%v /\ \_\ \   /\  __ \   /\ \       /\  __ \   /\  == \ /\  == \   /\  __ \     /\ \   /\  ___\   /\  ___\   /\__  _\    /\ \_\ \   /\  ___\   /\ \       /\  == \ 
%v \ \  __ \  \ \ \/\ \  \ \ \____  \ \ \/\ \  \ \  _-/ \ \  __<   \ \ \/\ \   _\_\ \  \ \  __\   \ \ \____  \/_/\ \/    \ \  __ \  \ \  __\   \ \ \____  \ \  _-/ 
%v  \ \_\ \_\  \ \_____\  \ \_____\  \ \_____\  \ \_\    \ \_\ \_\  \ \_____\ /\_____\  \ \_____\  \ \_____\    \ \_\     \ \_\ \_\  \ \_____\  \ \_____\  \ \_\   
%v   \/_/\/_/   \/_____/   \/_____/   \/_____/   \/_/     \/_/ /_/   \/_____/ \/_____/   \/_____/   \/_____/     \/_/      \/_/\/_/   \/_____/   \/_____/   \/_/   
%v
%v┌──────────────┬────────────────────────────────┬────────────────────────────────────────────┐
%v│ Команда      │ Описание                       │ Пример                                     │
%v├──────────────┼────────────────────────────────┼────────────────────────────────────────────┤
%v│ cripto       │ Шифрование строки              │ cripto key_word key_password { payload }   │
%v│ ecripto      │ Дешифровка строки              │ ecripto key_word key_password              │
%v│              │                                │                                            │
%v│ gmaster      │ Генерация мастер-ключа         │                                            │
%v│ gkey         │ Генерация ключа для шифра      │                                            │
%v│              │                                │                                            │
%v│ clog         │ Очистка логов                  │                                            │
%v│ drop         │ Удалить данные приложения      │                                            │
%v│ stop         │ Остановить приложение          │                                            │
%v│ help         │ Список команд                  │                                            │
%v│              │                                │                                            │
%v│ declare      │ Добавить команду               │ declare key_word { payload }               │
%v│ commands     │ Просмотреть список команд      │                                            │
%v│ run          │ Запустить команду              │ run key_word                               │
%v│ runm         │ Запустить множественную        │ runm key_word                              │
%v│ rmc          │ Удалить команду                │ rmc key_word                               │
%v│              │                                │                                            │
%v│              │ in progress                    │                                            │
%v│              │                                │                                            │
%v│ note         │ Открыть запись                 │                                            │
%v│ dnote        │ Создание записи                │ dnote key_word key_password                │
%v│ notes        │ Получить список записей        │                                            │
%v└──────────────┴────────────────────────────────┴────────────────────────────────────────────┘
`

// const CHAR_SELECTED_ITEM = "● "
// const CHAR_UN_SELECTED_ITEM = "○ "
const CHAR_SELECTED_ITEM = "> "
const CHAR_UN_SELECTED_ITEM = ""

var (
	StyleError    = []int{literals.SGR.RED, literals.SGR.BOLD}
	StyleSelected = []int{literals.SGR.RED, literals.SGR.BOLD, literals.SGR.UNDERLINE}
)

var AcceptableYeaOrNot []string = []string{"yes", "no", "y", "n", "yea", "not"}
