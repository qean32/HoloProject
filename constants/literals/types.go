package literals

type commandsType struct {
	CRIPTO  string
	ECRIPTO string

	GMASTER string
	GKEY    string

	CLOG string
	DROP string
	STOP string
	HELP string

	NOTE  string
	DNOTE string
	NOTES string

	DECLARE  string
	COMMANDS string
	RUN      string
	RUNM     string
	RMC      string

	MENU string
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
