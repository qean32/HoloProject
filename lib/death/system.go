package death

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"main/constants"
	errorlist "main/constants/error-list"
)

func ReadFile(path string) ([]string, error) {
	file, err := os.Open(filepath.Join(constants.Root, path))
	if err != nil {
		return nil, errorlist.OpenFile(path, err)
	}
	defer file.Close()

	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, errorlist.ReadFile(path, err)
	}
	return data, nil
}

func WriteFile(data, path string) error {
	file, err := os.Create(filepath.Join(constants.Root, path))
	if err != nil {
		return errorlist.CreateFile(path, err)
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		return errorlist.WriteFile(path, err)
	}
	return nil
}

func PushToFile(path, newText string) error {
	data, err := ReadFile(path)
	if err != nil {
		return err
	}
	return WriteFile(strings.Join(append(data, newText+"\n"), " \n"), path)
}

func CreateFile(path string, content string) error {
	if err := os.WriteFile(filepath.Join(constants.Root, path), []byte(content), 0644); err != nil {
		return errorlist.CreateFile(path, err)
	}
	return nil
}

func ClearFile(path string) error {
	return WriteFile("", path)
}

func RemoveFile(path string) error {
	if err := os.Remove(filepath.Join(constants.Root, path)); err != nil {
		return errorlist.RemoveFile(path, err)
	}
	return nil
}

func ListFilesByExt(dir string, ext string) ([]string, error) {
	var files []string

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, errorlist.ReadFile(dir, err)
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

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}
