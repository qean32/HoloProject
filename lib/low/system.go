package low

import (
	"bufio"
	"os"
	"path/filepath"
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

func CreateFile(path string, content string) bool {
	err := os.WriteFile(constants.Root+path, []byte(content), 0644)
	return err == nil
}

func ClearFile(path string) {
	WriteFile("", path)
}

func ListFilesByExt(dir string, ext string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		if filepath.Ext(entry.Name()) == ext {
			files = append(files, entry.Name())
		}
	}
	return files, nil
}
