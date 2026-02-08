package lib

import (
	"bufio"
	"fmt"
	"main/depth"
	"main/model"
	"os"
	"strings"
	"time"
)

var READER = bufio.NewReader(os.Stdin)

func ENTER_COMMAND() {
	command, _ := READER.ReadString('\n')
	parsed := PARSE_EVENT(command)
	fmt.Printf("%+v \n", parsed)

	depth.LOG()
	ENTER_COMMAND()
}

func HOF_ACCESS_ACTION(funcion func()) {
	if depth.ACCESS_ACTION() {
	}
}

func INIT() {
	for i := 0; i < len(PROJECT_INIT); i++ {
		fmt.Print(PROJECT_INIT[i])
	}
	ENTER_COMMAND()
}

func PUSH_CYCLE() {
}

func PARSE_EVENT(command string) model.Event {
	arr := strings.Split(command, " ")
	var payload string
	start := strings.Index(command, "{")
	if start {
		end := strings.Index(command, "}")
		payload = command[start:end]
	}
	var event model.Event = model.Event{
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Key:     arr[0],
		Payload: payload,
		Hash:    arr[2],
		Flags:   filterIsFlag(arr),
	}

	return event
}
