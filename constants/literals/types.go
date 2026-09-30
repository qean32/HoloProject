package literals

type eventType struct {
	CRYPTO   string
	DECRYPTO string

	GENERATEMASTER string
	GENERATEKEY    string

	CLEARLOG string
	LOG      string

	DROP string
	STOP string
	HELP string

	BAT       string
	RUNBAT    string
	REMOVEBAT string

	BATLIST       string
	BATLISTREMOVE string

	MENU   string
	MANUAL string

	// ТЕХНИЧЕСКИЕ

	RESPONSE     string
	INWORK       string
	MANUADDCMD   string
	MENUCRYPTO   string
	MENUDECRYPTO string
	RRBAT        string
}

type extensionType struct {
	Bat string
}

type responseType struct {
	Success          string
	UndefinedCommand string
	SyntaxError      string
}

type pathType struct {
	PathLog     string
	PathSetting string
}

type titleType struct {
	Menu     string
	ListCmd  string
	EnterCmd string
	Access   string
}

type flagsType struct {
	NOLOG string
}

type SGRtype struct {
	BOLD   int
	ITALIC int

	DIM int

	UNDERLINE int

	RED     int
	GREEN   int
	YELLOW  int
	BLUE    int
	MAGENTA int
	CYAN    int
	WHITE   int

	RESET int
}
