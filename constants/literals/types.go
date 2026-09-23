package literals

type commandsType struct {
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

	DECLARE            string
	COMMANDLIST        string
	RUNCOMMAND         string
	RUNMULTIPLECOMMAND string
	REMOVECOMMAND      string

	MENU   string
	MANUAL string

	MANUADDCOMMAND string
	MENUCRYPTO     string
	MENUDECRYPTO   string
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

type flagsType struct {
	NOLOG string
}

var SGR = SGRtype{
	BOLD:      1,
	ITALIC:    3,
	DIM:       2,
	UNDERLINE: 4,

	RED:     31,
	GREEN:   32,
	YELLOW:  33,
	BLUE:    34,
	MAGENTA: 35,
	CYAN:    36,
	WHITE:   37,

	RESET: 0,
}
