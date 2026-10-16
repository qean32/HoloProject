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

	MENU = "menu"
	INIT = "init"

	// ТЕХНИЧЕСКИЕ

	QNEXT         = "qcallback"
	INWORK        = "inwork"
	RESPONSE      = "response"
	MANUCREATEBAT = "MANUCREATEBAT"
	MENUDECRYPTO  = "MENUDECRYPTO"
	MENUCRYPTO    = "MENUCRYPTO"
	PUREREMOVEBAT = "REMOVEBAT"
	RESET         = "r"
	MANUAL        = "manual"
	TEST          = "test"
	TESTVIEW      = "testview"
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
	PUREREMOVEBAT,
	RESET,
	TEST,
	MANUAL,
}
