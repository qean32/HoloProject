package eventlist

// КЛИЕНТСКИЕ
const (
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

	LOGGER = "logger"
)

// ТЕХНИЧЕСКИЕ
const (
	INWORK        = "inwork"
	RESPONSE      = "response"
	QCREATEBAT    = "QCREATEBAT~"
	PUREREMOVEBAT = "PUREREMOVEBAT~"
	RESET         = "r"
	MANUAL        = "manual"
	TEST          = "test"
	TESTVIEW      = "testview"
)
