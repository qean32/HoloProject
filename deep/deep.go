package deep

import (
	"bufio"
	"main/constants"
	"main/model"
	"os"
	"strings"
)

func LOG(event model.Event) {
}

func ITERATION_CYCLE() {
}

var CALLSTACK = []string{}

func ReadFile(path string) []string {
	file, err := os.Open(constants.Root + path)
	scanner := bufio.NewScanner(file)

	if err != nil {
		return nil
	}
	defer file.Close()

	var data []string
	for scanner.Scan() {
		data =
			append(data, scanner.Text())
	}
	return data
}

func WriteFile(data string, path string) bool {
	file, err := os.Create(constants.Root + path)

	if err != nil {
		return false
	}
	defer file.Close()

	file.WriteString(data)
	return true
}

func PushToFile(path string, newText string) bool {
	return WriteFile(strings.Join(
		append(ReadFile(path), newText+"\n"), " \n"), path)
}
