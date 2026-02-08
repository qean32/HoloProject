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
	print("Enter command: ")
	command, _ := READER.ReadString('\n')
	if len(command) > 1 {
		event := PARSE_EVENT(command)
		fmt.Printf("%+v \n", event)

		depth.LOG()
	}
	ENTER_COMMAND()
}

func HOF_ACCESS_ACTION(f model.EventFunction, event model.Event) {
	if depth.ACCESS_ACTION() {
		f(event)
		return
	}

	fmt.Println("Command stop/")
}

func INIT() {
	var text, _ = depth.ReadFile("/log.txt")
	fmt.Println(depth.WriteFile(strings.Join(text, " \n"), "/log.txt"))
	for i := 0; i < len(PROJECT_INIT); i++ {
		fmt.Print(PROJECT_INIT[i])
	}
	ENTER_COMMAND()
}

func PUSH_CYCLE() {
}

func PARSE_EVENT(command string) model.Event {
	arr := strings.Split(command, " ")
	payload := getPaylaod(command)
	var event model.Event = model.Event{
		Time:     time.Now().Format("2006-01-02 15:04:05"),
		Key:      arr[0],
		Password: arr[1],
		Payload:  payload,
		Flags:    filterIsFlag(arr),
	}

	return event
}
