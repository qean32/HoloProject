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

var KEY_PARSE = map[string]func(command string) model.Event{
	"hash": func(command string) model.Event {
		return model.Event{}
	},
	"g:key": func(command string) model.Event {
		return model.Event{}
	},
	"g:log": func(command string) model.Event {
		return model.Event{}
	},
	"c:log": func(command string) model.Event {
		return model.Event{}
	},
	"dihash": func(command string) model.Event {
		return model.Event{}
	},
	"master": func(command string) model.Event {
		return model.Event{}
	},
	"drop": func(command string) model.Event {
		return model.Event{}
	},
	"stop": func(command string) model.Event {
		return model.Event{}
	},
	"help": func(command string) model.Event {
		return model.Event{}
	},
}

var FLAG = map[string]string{
	"f":     "",
	"force": "",
	"nl":    "",
}

const PROJECT_NAME = "Holo"

var PROJECT_INIT = `
       )                  )    (             )                            
    ( /(  ( /(   )\ )  ( /(    )\ ) )\ )  ( /(                (        ) 
    )\( ) )\()) (()/(  )\())  (()/((()/(  )\())    (   (      )\   )  /( 
   ((_)\ ((_)\   /(_))((_)\    /(_))/(_))((_)\     )\  )\   (((_)  ( )(_))
    _((_)  ((_) (_))    ((_)  (_)) (_))    ((_)   ((_)((_)  )\___ (_(_()) 
   | || | / _ \ | |    / _ \  | _ \| _ \  / _ \  _ | || __|((/ __||_   _| 
   | __ || (_) || |__ | (_) | |  _/|   / | (_) || || || _|  | (__   | |   
   |_||_| \___/ |____| \___/  |_|  |_|_\  \___/  \__/ |___|  \___|  |_|   
`
