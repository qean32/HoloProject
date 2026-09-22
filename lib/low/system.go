package low

import (
	"bufio"
	"main/constants"
	"os"
	"strings"
)

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

func CreateFile(path string) {
	file, _ := os.Create(constants.Root + path)
	defer file.Close()
}

func ClearFile(path string) {
	WriteFile("", path)
}
