package depth

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const root = "./private"

func ACCESS_ACTION() bool {
	var response string
	fmt.Print("Need access (yes/no): ")
	fmt.Scan(&response)

	if response == "yes" {

		return true
	}

	return false
}

func LOG() {
}

func ITERATION_CYCLE() {
}

var CALLSTACK = []string{}

func ReadFile(path string) []string {
	file, err := os.Open(root + path)
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
	file, err := os.Create(root + path)

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
