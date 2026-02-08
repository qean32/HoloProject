package lib

import (
	"main/model"
)

var KEY_FUNCTION = map[string]func(e model.Event){
	"hash":         func(e model.Event) {}, // hash "ключ" "сообщество" "хеш" (-nl)
	"dihash":       func(e model.Event) {}, // dihash "ключ" (-nl)
	"master":       func(e model.Event) {}, // master (-nl)
	"generate:key": func(e model.Event) {}, // generate:key
	"drop":         func(e model.Event) {}, // drop (-f)
	"clear:log":    func(e model.Event) {}, // clear:log (-nl -f)
	"stop":         func(e model.Event) {}, // stop
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
