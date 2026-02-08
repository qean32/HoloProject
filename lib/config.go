package lib

import (
	"main/model"
)

var KEY_FUNCTION = map[string]func(e model.Event){
	"hash":   func(e model.Event) {}, // hash ключ пароль {сообщение} (-nl)
	"g:key":  func(e model.Event) {}, // g:key
	"g:log":  func(e model.Event) {}, // g:log
	"c:log":  func(e model.Event) {}, // c:log (-nl -f)
	"dihash": func(e model.Event) {}, // dihash ключ пароль (-nl)
	"master": func(e model.Event) {}, // master (-nl)
	"drop":   func(e model.Event) {}, // drop (-f)
	"stop":   func(e model.Event) {}, // stop
	"help":   func(e model.Event) {}, // help
}

var FLAG = map[string]string{
	"f":     "",
	"force": "",
	"nl":    "",
}

const PROJECT_NAME = "Holo"

var PROJECT_INIT = []string{
	"\n",
	"                           )      )    (         )       (      (         )                                                         \n ",
	"                        ( /(   ( /(    )\\ )   ( /(       )\\ )   )\\ )   ( /(                   (      *   )                          \n ",
	"                        )\\())  )\\())  (()/(   )\\())     (()/(  (()/(   )\\())     (    (       )\\     )  /(                          \n ",
	"                       ((_)\\  ((_)\\    /(_)) ((_)\\       /(_))  /(_)) ((_)\\      )\\   )\\     ((_)   ( )(_))                         \n ",
	"                        _((_)   ((_)  (_))     ((_)     (_))   (_))     ((_)    ((_) ((_)   )\\___  (_(_())                          \n ",
	"                       | || |  / _ \\  | |     / _ \\     | _ \\  | _ \\   / _ \\   _ | | | __| ((/ __| |_   _|                          \n ",
	"                       | __ | | (_) | | |__  | (_) |    |  _/  |   /  | (_) | | || | | _|   | (__    | |                            \n ",
	"                       |_||_|  \\___/  |____|  \\___/     |_|    |_|_\\   \\___/   \\__/  |___|   \\___|   |_|                            \n ",
	"\n",
}
