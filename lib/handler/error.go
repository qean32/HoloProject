package handler

import "errors"

var (
	ErrAccessDenied   = errors.New("Процесс остановлен")
	ErrNotExist       = errors.New("файл не существует")
	ErrIsDir          = errors.New("это директория")
	ErrInvalidName    = errors.New("некорректное имя")
	ErrNoSubEvent     = errors.New("SubEvent не задан")
	ErrorInwork       = errors.New("В разработке!")
	ErrUnknownCommand = errors.New("неизвестная команда")
	ErrEmptyCommand   = errors.New("пустая команда")
	ErrTooFewArgs     = errors.New("недостаточно аргументов")
	ErrNoPayload      = errors.New("payload не найден")
)
