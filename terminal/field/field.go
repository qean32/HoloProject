package field

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"

	"main/lib/utils"
	"main/model"
	"main/terminal"
)

var reader = bufio.NewReader(os.Stdin)

func Field(payload model.FieldPayload) (string, error) {
	terminal.Print(utils.GetCustomMessage(payload.Prefix, payload.PrefixSRG...))
	text, err := reader.ReadString('\n')
	terminal.DownAndStart()

	if err != nil && !(errors.Is(err, io.EOF) && text != "") {
		return "", err
	}
	return strings.TrimSpace(text), nil
}
