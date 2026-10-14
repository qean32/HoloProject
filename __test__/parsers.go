package test

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	errorlist "main/constants/error-list"
	eventlist "main/constants/event-list"
	"main/constants/literal"
	"main/lib/handler/parse"
	"main/model"
)

// parserCase — один подтест парсера: команда и ожидания к результату.
type parserCase struct {
	name    string
	command string
	key     string

	wantErr   error // ожидаемая ошибка (nil — ошибки быть не должно)
	want      model.Event
	noPayload bool
}

// parserCases формирует подтесты по всем парсерам из parse.MAP:
// по одному корректному и одному ошибочному сценарию на парсер.
func parserCases() []parserCase {
	return []parserCase{
		{
			name:    "crypto/корректная команда",
			command: "crypto my_key my_pass {start cmd /K echo 67}",
			key:     eventlist.CRYPTO,
			want: model.Event{
				Key:      eventlist.CRYPTO,
				KeyWord:  "my_key",
				Password: "my_pass",
				Payload:  "start cmd /K echo 67",
			},
		},
		{
			name:      "crypto/мало аргументов",
			command:   "crypto my_key",
			key:       eventlist.CRYPTO,
			wantErr:   errorlist.ErrTooFewArgs,
			noPayload: true,
		},
		{
			name:      "crypto/нет payload",
			command:   "crypto my_key my_pass",
			key:       eventlist.CRYPTO,
			wantErr:   errorlist.ErrNoPayload,
			noPayload: true,
		},
		{
			name:      "ecrypto/корректная команда",
			command:   "ecrypto my_key my_pass",
			key:       eventlist.DECRYPTO,
			noPayload: true,
			want: model.Event{
				Key:      eventlist.DECRYPTO,
				KeyWord:  "my_key",
				Password: "my_pass",
			},
		},
		{
			name:      "ecrypto/мало аргументов",
			command:   "ecrypto my_key",
			key:       eventlist.DECRYPTO,
			wantErr:   errorlist.ErrTooFewArgs,
			noPayload: true,
		},
		{
			name:    "bat/корректная команда",
			command: "bat my_key {start cmd /K echo 67}",
			key:     eventlist.BAT,
			want: model.Event{
				Key:     eventlist.BAT,
				KeyWord: "my_key",
				Payload: "start cmd /K echo 67",
			},
		},
		{
			name:      "bat/мало аргументов",
			command:   "bat my_key",
			key:       eventlist.BAT,
			wantErr:   errorlist.ErrTooFewArgs,
			noPayload: true,
		},
		{
			name:      "bat/нет payload",
			command:   "bat my_key my_pass",
			key:       eventlist.BAT,
			wantErr:   errorlist.ErrNoPayload,
			noPayload: true,
		},
		{
			name:      "run/корректная команда",
			command:   "run my_key",
			key:       eventlist.RUNBAT,
			noPayload: true,
			want: model.Event{
				Key:     eventlist.RUNBAT,
				KeyWord: "my_key",
			},
		},
		{
			name:      "run/мало аргументов",
			command:   "run",
			key:       eventlist.RUNBAT,
			wantErr:   errorlist.ErrTooFewArgs,
			noPayload: true,
		},
		{
			name:    "rbat/корректная команда",
			command: "rbat my_key",
			key:     eventlist.REMOVEBAT,
			want: model.Event{
				Key:     eventlist.REMOVEBAT,
				KeyWord: "my_key",
				SubEvent: &model.Event{
					Key:     eventlist.PURE_REMOVE_BAT,
					KeyWord: "my_key" + literal.Extension.Bat,
				},
			},
		},
		{
			name:      "rbat/мало аргументов",
			command:   "rbat",
			key:       eventlist.REMOVEBAT,
			wantErr:   errorlist.ErrTooFewArgs,
			noPayload: true,
		},
		{
			name:    "menu/событие без аргументов",
			command: "menu",
			key:     eventlist.MENU,
			want: model.Event{
				Key: eventlist.MENU,
			},
		},
		{
			name:    "menu/флаг",
			command: "menu -nl",
			key:     eventlist.MENU,
			want: model.Event{
				Key:   eventlist.MENU,
				Flags: []string{literal.Flags.NOLOG},
			},
		},
		{
			name:    "log/событие без аргументов",
			command: "log",
			key:     eventlist.OPENLOG,
			want: model.Event{
				Key: eventlist.OPENLOG,
			},
		},
		{
			// strings.Split("", " ") возвращает [""] (длина 1), поэтому ветка
			// len(arr) == 0 в parse.event через ParseEvent недостижима,
			// и errorlist.ErrEmptyCommand отсюда не вернётся.
			// Фиксируем фактическое поведение: событие с пустым Key.
			name:    "пустая команда",
			command: "",
			key:     eventlist.MENU,
			want: model.Event{
				Key: "",
			},
		},
	}
}

// runParserCase выполняет один подтест парсера.
func runParserCase(c parserCase) error {
	event, err := parse.ParseEvent(c.command, c.key)

	if c.wantErr != nil {
		if err == nil {
			return fmt.Errorf("ожидалась ошибка %q", c.wantErr)
		}
		if !errors.Is(err, c.wantErr) {
			return fmt.Errorf("ошибка %q, ожидалась %q", err, c.wantErr)
		}
		return nil
	}

	if err != nil {
		return err
	}
	if c.noPayload && event.Payload != "" {
		return fmt.Errorf("Payload = %q, ожидался пустой", event.Payload)
	}
	// DateTime заполняется текущим временем, поэтому сверяем только
	// факт заполнения, а само значение исключаем из сравнения.
	if event.DateTime == "" {
		return errors.New("DateTime не заполнен")
	}
	c.want.DateTime = event.DateTime

	if diff := compareEvents(c.want, event); diff != "" {
		return errors.New(diff)
	}
	return nil
}

// parserSubCases превращает таблицу парсеров в подтесты фреймворка.
func parserSubCases() []subCase {
	cases := parserCases()
	subs := make([]subCase, 0, len(cases)+2)

	subs = append(subs,
		subCase{group: "parse", name: "все парсеры на месте", run: allParsersCovered},
		subCase{group: "parse", name: "неизвестный ключ", run: unknownKeyFallsBack},
	)

	for _, c := range cases {
		subs = append(subs, subCase{
			group: "parse",
			name:  c.name,
			run:   func() error { return runParserCase(c) },
		})
	}
	return subs
}

// allParsersCovered проверяет, что для каждого парсера из parse.MAP
// есть хотя бы один подтест.
func allParsersCovered() error {
	covered := make(map[string]bool)
	for _, c := range parserCases() {
		covered[c.key] = true
	}

	var missing []string
	for key := range parse.MAP {
		if !covered[key] {
			missing = append(missing, key)
		}
	}
	if len(missing) > 0 {
		slices.Sort(missing)
		return fmt.Errorf("нет подтестов для парсеров: %s", strings.Join(missing, ", "))
	}
	return nil
}

// unknownKeyFallsBack проверяет, что для ключа без парсера
// используется общий разбор по первому слову.
func unknownKeyFallsBack() error {
	event, err := parse.ParseEvent("clog", eventlist.CLEARLOG)
	if err != nil {
		return err
	}
	if event.Key != eventlist.CLEARLOG {
		return fmt.Errorf("Key = %q, ожидался %q", event.Key, eventlist.CLEARLOG)
	}
	return nil
}
