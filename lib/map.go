package lib

import (
	"fmt"
	"main/constants"
	"main/deep"
	"main/model"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var KEY_FUNCTION = map[string]func(e model.Event){
	"hash": func(e model.Event) {
		deep.TMP_DATA = append(deep.TMP_DATA, []string{e.KeyWord, e.Payload})
		deep.PushToFile(constants.DATA_PATH, fmt.Sprintf("%s %s", e.KeyWord, e.Payload))
	},
	"dihash": func(e model.Event) {
		if len(deep.TMP_DATA) == 0 {
			strs := (deep.ReadFile(constants.DATA_PATH))
			var tmpArr [][]string

			for i := 0; i < len(strs); i++ {
				tmp := strings.Split(strs[i], " ")
				tmpArr = append(tmpArr, []string{tmp[0], strings.Join(tmp[1:], " ")})
			}
			deep.TMP_DATA = tmpArr
		}

		index := slices.IndexFunc(deep.TMP_DATA, func(item []string) bool {
			return item[0] == e.KeyWord
		})

		if index != -1 {
			fmt.Println(deep.TMP_DATA[index][1])
		} else {
			fmt.Println(constants.UNDEFINED_WORD_KEY)
		}
	},
	"g:key": func(e model.Event) {},
	"c:log": func(e model.Event) {
		deep.ClearFile(constants.LOG_PATH)
	},
	"master": func(e model.Event) {
		deep.CreateFile(constants.LOG_PATH)
		deep.CreateFile(constants.COMMAND_PATH)
		deep.CreateFile(constants.DATA_PATH)
	},
	"drop": func(e model.Event) {
		os.RemoveAll(constants.Root)
		os.Mkdir(constants.Root, 0755)
	},
	"stop": func(e model.Event) {
		os.Exit(0)
	},
	"help": func(e model.Event) {
		fmt.Println(constants.HelpMessage)
	},
	"place": func(e model.Event) {
		deep.PushToFile(constants.COMMAND_PATH, fmt.Sprintf("%s %s", e.KeyWord, e.Payload))
	},
	"run": func(e model.Event) {
		if len(deep.TMP_COMMANDS) == 0 {
			strs := (deep.ReadFile(constants.COMMAND_PATH))
			var tmpArr [][]string

			for i := 0; i < len(strs); i++ {
				tmp := strings.Split(strs[i], " ")
				tmpArr = append(tmpArr, []string{tmp[0], strings.Join(tmp[1:], " ")})
			}
			deep.TMP_COMMANDS = tmpArr
		}

		index := slices.IndexFunc(deep.TMP_COMMANDS, func(item []string) bool {
			return item[0] == strings.TrimSpace(e.KeyWord)
		})

		if index != -1 {
			cmd := exec.Command("CMD.exe", "/C", deep.TMP_COMMANDS[index][1])
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка при запуске команды:", err)
				return
			}
		} else {
			fmt.Println(constants.UNDEFINED_WORD_KEY)
		}
	},
	"commands": func(e model.Event) {
		fmt.Println(strings.Join(deep.ReadFile(constants.COMMAND_PATH), "\n"))
	},
}

var KEY_PARSE = map[string]func(arr []string) (event model.Event, _error bool){
	"hash": func(arr []string) (event model.Event, _error bool) {
		payload := getPaylaod(arr)

		if len(arr) < 3 || payload == "" {
			_error = true
			return
		}

		event = model.Event{
			Time:     deep.NewTime(),
			Key:      arr[0],
			KeyWord:  arr[1],
			Password: arr[2],
			Payload:  payload,
			Flags:    filterIsFlag(arr),
		}
		return
	},
	"g:key": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		return
	},
	"g:log": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		return
	},
	"c:log": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		return
	},
	"dihash": func(arr []string) (event model.Event, _error bool) {
		if len(arr) < 3 {
			_error = true
			return
		}

		event = model.Event{
			Time:     deep.NewTime(),
			Key:      arr[0],
			KeyWord:  arr[1],
			Password: arr[2],
			Flags:    filterIsFlag(arr),
		}
		return
	},
	"master": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		return
	},
	"drop": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		return
	},
	"stop": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		return
	},
	"help": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		return
	},
	"place": func(arr []string) (event model.Event, _error bool) {
		payload := getPaylaod(arr)

		if len(arr) < 3 || payload == "" {
			_error = true
			return
		}

		event = model.Event{
			Time:    deep.NewTime(),
			Key:     arr[0],
			KeyWord: arr[1],
			Payload: payload,
		}
		return
	},
	"run": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time:    deep.NewTime(),
			Key:     arr[0],
			KeyWord: arr[1],
		}
		return
	},
	"commands": func(arr []string) (event model.Event, _error bool) {
		event = model.Event{
			Time: deep.NewTime(),
			Key:  arr[0],
		}
		return
	},
}

var FLAG = map[string]string{
	"f":     "",
	"force": "",
	"nl":    "",
}
