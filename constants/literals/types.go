package literals

type eventType struct {
	CRYPTO   string
	DECRYPTO string

	GENERATEMASTER string
	GENERATEKEY    string

	CLEARLOG string
	LOGS     string

	DROP   string
	STOP   string
	HELP   string
	INWORK string

	DECLARE   string
	EventList string
	RUNCMD    string
	REMOVECMD string

	MENU   string
	MANUAL string

	RESPONSE string

	MANUADDCMD   string
	MENUCRYPTO   string
	MENUDECRYPTO string
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
