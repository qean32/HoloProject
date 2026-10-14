package field

import (
	"bufio"
	"main/lib/utils"
	"main/model"
	"main/terminal"
	"os"
	"strings"
)

var r = bufio.NewReader(os.Stdin)

func Field(payload model.FieldPayload) string {
	set(payload)
	terminal.Print(utils.GetCustomMessage(field.Prefix, field.PrefixSRG...))
	text, _ := r.ReadString('\n')

	return strings.TrimSpace(text)
}
