package eventlist

var (
	// КЛИЕНТСКИЕ

	CRYPTO   = "crypto"
	DECRYPTO = "ecrypto"

	GENERATEMASTER = "gmaster"
	GENERATEKEY    = "gkey"

	CLEARLOG = "clog"
	OPENLOG  = "log"

	DROP = "drop"
	STOP = "stop"
	HELP = "help"

	BAT       = "bat"
	RUNBAT    = "run"
	REMOVEBAT = "rbat"

	BATLIST       = "list"
	BATLISTREMOVE = "listremove"

	MENU  = "menu"
	QNEXT = "qcallback"
	INIT  = "init"

	// ТЕХНИЧЕСКИЕ

	INWORK          = "inwork"
	RESPONSE        = "response"
	MANUCREATEBAT   = "MANUCREATEBAT"
	MENUDECRYPTO    = "MENUDECRYPTO"
	MENUCRYPTO      = "MENUCRYPTO"
	PURE_REMOVE_BAT = "67REMOVEBAT67"
	RESET           = "reset"
	TEST            = "test"
	MANUAL          = "manual"
)

// All — все события проекта, клиентские и технические.
// По этому списку проходят встроенные тесты (main/__test__),
// поэтому новое событие нужно добавлять и сюда.
var All = []string{
	// клиентские
	CRYPTO,
	DECRYPTO,
	GENERATEMASTER,
	GENERATEKEY,
	CLEARLOG,
	OPENLOG,
	DROP,
	STOP,
	HELP,
	BAT,
	RUNBAT,
	REMOVEBAT,
	BATLIST,
	BATLISTREMOVE,
	MENU,
	QNEXT,
	INIT,

	// технические
	INWORK,
	RESPONSE,
	MANUCREATEBAT,
	MENUDECRYPTO,
	MENUCRYPTO,
	PURE_REMOVE_BAT,
	RESET,
	TEST,
	MANUAL,
}
