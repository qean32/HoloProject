package constants

import (
	"main/constants/literals"
	"os"
)

const Mode = "prod"
const projectname = "holoproject"

var (
	Root = `.\private\`
	Cmd  = `cmd\`
)

func INIT_ROOT() {
	if Mode == "dev" {
	} else {
		Root = os.TempDir() + `\` + projectname + `\`
		os.Mkdir(Root, 0755)
		os.Mkdir(Root+Cmd, 0755)
	}
}

var BinaryPROJECT_INIT = `
%v01101000 01101111 01101100 01101111 01110000 01110010 01101111 01101010 01100101 01100011 01110100
`

var PROJECT_INIT = `
%v __  __    ______    __        ______     ______  ______    ______      __    ______    ______    ______     
%v/\ \_\ \  /\  __ \  /\ \      /\  __ \   /\  == \/\  == \  /\  __ \    /\ \  /\  ___\  /\  ___\  /\__  _\    
%v\ \  __ \ \ \ \/\ \ \ \ \____ \ \ \/\ \  \ \  _-/\ \  __<  \ \ \/\ \  _\_\ \ \ \  __\  \ \ \____ \/_/\ \/    
%v \ \_\ \_\ \ \_____\ \ \_____\ \ \_____\  \ \_\   \ \_\ \_\ \ \_____\/\_____\ \ \_____\ \ \_____\   \ \_\    
%v  \/_/\/_/  \/_____/  \/_____/  \/_____/   \/_/    \/_/ /_/  \/_____/\/_____/  \/_____/  \/_____/    \/_/    
%v
`

// ASCII
// https://translated.turbopages.org/proxy_u/en-ru.ru.6c604ba4-6a11ef9b-b4e5d92b-74722d776562/https/student.cs.uwaterloo.ca/~cs452/terminal.html
// https://www.asciiart.eu/text-to-ascii-art respect
// https://patorjk.com/software/taag/#p=testall&f=Broadway&t=holoproject+2&x=none&v=4&h=4&w=80&we=false

var HelpMessage = `
%v  __  __    ______    __        ______      ______  ______    ______      __    ______    ______    ______     __  __    ______    __        ______  
%v /\ \_\ \  /\  __ \  /\ \      /\  __ \    /\  == \/\  == \  /\  __ \    /\ \  /\  ___\  /\  ___\  /\__  _\   /\ \_\ \  /\  ___\  /\ \      /\  == \ 
%v \ \  __ \ \ \ \/\ \ \ \ \____ \ \ \/\ \   \ \  _-/\ \  __<  \ \ \/\ \  _\_\ \ \ \  __\  \ \ \____ \/_/\ \/   \ \  __ \ \ \  __\  \ \ \____ \ \  _-/ 
%v  \ \_\ \_\ \ \_____\ \ \_____\ \ \_____\   \ \_\   \ \_\ \_\ \ \_____\/\_____\ \ \_____\ \ \_____\   \ \_\    \ \_\ \_\ \ \_____\ \ \_____\ \ \_\   
%v   \/_/\/_/  \/_____/  \/_____/  \/_____/    \/_/    \/_/ /_/  \/_____/\/_____/  \/_____/  \/_____/    \/_/     \/_/\/_/  \/_____/  \/_____/  \/_/   
%v
%v  crypto         *ифрование
%v  ecrypto        *ешифрование
%v
%v  bat            *обавить .bat                     bat key_word { payload }
%v  list           *писок .bat
%v  run            *апустить .bat                    run key_word
%v  rbat           *далить .bat                      rbat key_word
%v
%v  gmaster        *енерация мастер-ключа
%v  gkey           *енерация ключа
%v
%v  menu
%v  clog
%v  log
%v  drop
%v  stop
%v  help
%v
`

const CHAR_SELECTED_ITEM = "+ "
const CHAR_UN_SELECTED_ITEM = "│  "

const FIELDPREFIX = "-> "
const RESPONSEPREFIX = " <- вывод"

var (
	StyleError    = []int{literals.SGR.RED, literals.SGR.BOLD}
	StyleDim      = []int{literals.SGR.DIM}
	StyleSelected = []int{literals.SGR.RED, literals.SGR.BOLD, literals.SGR.UNDERLINE}
)

var AcceptableYeaOrNot []string = []string{"yes", "y", "yea"}
