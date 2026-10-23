package model

type ParseOptions struct {
	KeyWord     bool
	Payload     bool
	Password    bool
	GetSubEvent bool
	SubEvent    string
}

type Command struct {
	Name    string
	Title   string
	Usage   string
	Options ParseOptions
	Run     Handler
	Flow    func(Event) (Action, error)
	Success bool
	Hidden  bool
}
