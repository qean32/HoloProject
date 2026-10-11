package terminal

import (
	"fmt"
	"strings"

	"main/constants/sgr"
	"main/lib/utils"
	"main/model"
)

func PrintTechInfo(event model.Event) {
	Println(utils.GetCustomMessage("Технический вывод:", sgr.UNDERLINE, sgr.RED))
	printEvent(event, "  ")
}

func printEvent(event model.Event, indent string) {
	label := func(name, value string) {
		Println(fmt.Sprintf("%s%s %s", indent, utils.GetCustomMessage(fmt.Sprintf("%-8s =", name), sgr.DIM), value))
	}
	field := func(name, value string) {
		if value == "" {
			label(name, utils.GetCustomMessage("<empty>", sgr.DIM))
			return
		}
		label(name, utils.GetCustomMessage(value, sgr.RED))
	}

	field("Key", event.Key)
	field("KeyWord", event.KeyWord)
	field("Payload", event.Payload)
	field("Password", event.Password)
	field("DateTime", event.DateTime)
	label("Flags", formatFlags(event.Flags))

	if event.SubEvent != nil {
		label("SubEvent", "")
		printEvent(*event.SubEvent, indent+"  ")
	} else {
		label("SubEvent", utils.GetCustomMessage("<nil>", sgr.DIM))
	}
}

func formatFlags(flags []string) string {
	if len(flags) == 0 {
		return utils.GetCustomMessage("<empty>", sgr.DIM)
	}
	parts := make([]string, len(flags))
	for i, flag := range flags {
		parts[i] = utils.GetCustomMessage(flag, sgr.RED)
	}
	sep := utils.GetCustomMessage(", ", sgr.DIM)
	return utils.GetCustomMessage("[", sgr.DIM) + strings.Join(parts, sep) + utils.GetCustomMessage("]", sgr.DIM)
}
