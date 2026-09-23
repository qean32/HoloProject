package low

import (
	"bufio"
	"os"
	"strings"

	"main/constants"
)

func ReadFile(path string) []string {
	file, err := os.Open(constants.Root + path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	return data
}

func WriteFile(data, path string) bool {
	file, err := os.Create(constants.Root + path)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err == nil
}

func PushToFile(path, newText string) bool {
	data := ReadFile(path)
	if data == nil {
		return false
	}
	return WriteFile(strings.Join(append(data, newText+"\n"), " \n"), path)
}

func CreateFile(path string) bool {
	file, err := os.Create(constants.Root + path)
	if err != nil {
		return false
	}
	file.Close()
	return true
}

func ClearFile(path string) {
	WriteFile("", path)
}
