package test

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"main/callstack"
	eventlist "main/constants/event-list"
	"main/lib/handler/parse"
	"main/model"
)

// commands — команды ручного ввода для событий, которым нужны аргументы.
// Для остальных событий из eventlist.All используется имя события.
var commands = map[string]string{
	eventlist.CRYPTO:    "crypto my_key my_pass {start cmd /K echo 67}",
	eventlist.DECRYPTO:  "ecrypto my_key my_pass",
	eventlist.BAT:       "bat my_key {start cmd /K echo 67}",
	eventlist.RUNBAT:    "run my_key",
	eventlist.REMOVEBAT: "rbat my_key",
}

// eventCases формирует подтесты по всем Event из eventlist:
// разбор команды ручного ввода и прохождение события через callstack.
func eventCases() []subCase {
	cases := make([]subCase, 0, len(eventlist.All)*2+3)

	cases = append(cases,
		subCase{group: "eventlist", name: "непустые ключи", run: keysNotEmpty},
		subCase{group: "eventlist", name: "уникальность ключей", run: keysUnique},
		subCase{group: "eventlist", name: "парсеры ссылаются на события", run: parsersKnown},
	)

	for _, key := range eventlist.All {
		cases = append(cases,
			subCase{group: "event", name: key, run: func() error { return parseEvent(key) }},
			subCase{group: "event", name: key + " (callstack)", run: func() error { return roundTrip(key) }},
		)
	}

	return cases
}

// commandOf возвращает команду ручного ввода для события key.
func commandOf(key string) string {
	if command, ok := commands[key]; ok {
		return command
	}
	return key
}

// parseEvent разбирает команду ручного ввода события key и проверяет,
// что полученное событие соответствует ключу.
func parseEvent(key string) error {
	command := commandOf(key)

	event, err := parse.ParseEvent(command, key)
	if err != nil {
		return fmt.Errorf("разбор %q: %w", command, err)
	}
	if event.Key != key {
		return fmt.Errorf("Key = %q, ожидался %q", event.Key, key)
	}
	if event.DateTime == "" {
		return errors.New("DateTime не заполнен")
	}
	return nil
}

// roundTrip проверяет, что событие проходит через callstack без изменений.
func roundTrip(key string) error {
	event, err := parse.ParseEvent(commandOf(key), key)
	if err != nil {
		return err
	}

	defer drain()
	callstack.Push(event)

	got, ok := callstack.Pop()
	if !ok {
		return errors.New("событие не вернулось из очереди")
	}
	if diff := compareEvents(event, got); diff != "" {
		return errors.New(diff)
	}
	return nil
}

// compareEvents возвращает описание различий двух событий
// или пустую строку, если события совпадают.
func compareEvents(want, got model.Event) string {
	var problems []string

	field := func(name, want, got string) {
		if want != got {
			problems = append(problems, fmt.Sprintf("%s = %q, ожидалось %q", name, got, want))
		}
	}
	field("Key", want.Key, got.Key)
	field("KeyWord", want.KeyWord, got.KeyWord)
	field("Payload", want.Payload, got.Payload)
	field("Password", want.Password, got.Password)
	field("DateTime", want.DateTime, got.DateTime)

	if !slices.Equal(want.Flags, got.Flags) {
		problems = append(problems, fmt.Sprintf("Flags = %v, ожидалось %v", got.Flags, want.Flags))
	}

	switch {
	case want.SubEvent == nil && got.SubEvent != nil:
		problems = append(problems, "SubEvent не ожидался")
	case want.SubEvent != nil && got.SubEvent == nil:
		problems = append(problems, "SubEvent отсутствует")
	case want.SubEvent != nil:
		field("SubEvent.Key", want.SubEvent.Key, got.SubEvent.Key)
		field("SubEvent.KeyWord", want.SubEvent.KeyWord, got.SubEvent.KeyWord)
	}

	return strings.Join(problems, "; ")
}

// keysNotEmpty проверяет, что у всех событий заполнен ключ.
func keysNotEmpty() error {
	for _, key := range eventlist.All {
		if key == "" {
			return errors.New("в eventlist.All есть пустой ключ")
		}
	}
	return nil
}

// keysUnique проверяет, что ключи событий не повторяются.
func keysUnique() error {
	seen := make(map[string]bool, len(eventlist.All))
	for _, key := range eventlist.All {
		if seen[key] {
			return fmt.Errorf("ключ %q встречается несколько раз", key)
		}
		seen[key] = true
	}
	return nil
}

// parsersKnown проверяет, что каждый парсер привязан к событию из eventlist.
func parsersKnown() error {
	for key := range parse.MAP {
		if !slices.Contains(eventlist.All, key) {
			return fmt.Errorf("парсер %q отсутствует в eventlist.All", key)
		}
	}
	return nil
}

// drain забирает из очереди все оставшиеся события,
// чтобы тесты не влияли на работу программы.
func drain() {
	for {
		if _, ok := callstack.Pop(); !ok {
			return
		}
	}
}
