package literals

type commandsType struct {
	CRIPTO  string
	ECRIPTO string

	GENERATEMASTER string
	GENERATEKEY    string

	CLEARLOG string
	DROP     string
	STOP     string
	HELP     string

	NOTE       string
	DELETENOTE string
	NOTES      string

	DECLARE               string
	COMMANDS_LIST         string
	RUN_COMMAND           string
	RUN_MULTIPLE_COMMANDS string
	REMOVE_COMMAND        string

	MENU   string
	INWORK string
	MANUAL string

	MENU_COMMANDS_LIST string
	MANU_ADD_COMMAND   string
	MENU_HASH_STRING   string
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
}
