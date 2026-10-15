package callstack

import (
	"fmt"
	"sync"

	"main/model"
)

var (
	errorsMu sync.Mutex
	Errors   []model.StackError
)

func PushError(event model.Event, err error) {
	if err == nil {
		return
	}
	errorsMu.Lock()
	defer errorsMu.Unlock()
	Errors = append(Errors, model.StackError{Event: event, Err: err})
}

func GetErrorsString() string {
	errorsMu.Lock()
	defer errorsMu.Unlock()

	if len(Errors) == 0 {
		return "no error"
	}

	res := ""

	for i, stackErr := range Errors {
		res = res + StackErrorToString(stackErr, i)
	}
	fmt.Println(res)
	return res
}

func StackErrorToString(stackErr model.StackError, i int) string {
	return fmt.Sprintf("[%d] Event: %+v, Error: %v\n", i+1, stackErr.Event, stackErr.Err)
}
