package errorlist

import "errors"

var (
	// ОБЩИЕ
	ErrAccessDenied   = errors.New("Процесс остановлен")
	ErrNotExist       = errors.New("файл не существует")
	ErrIsDir          = errors.New("это директория")
	ErrInvalidName    = errors.New("некорректное имя")
	ErrForbiddenName  = errors.New("Недопустимое имя")
	ErrForbiddenChars = errors.New("Недопустимые символы")

	// СОБЫТИЯ
	ErrNoSubEvent     = errors.New("SubEvent не задан")
	ErrUnknownCommand = errors.New("Неизвестная команда")

	// ПАРСИНГ
	ErrEmptyCommand = errors.New("Пустая команда")
	ErrTooFewArgs   = errors.New("Недостаточно аргументов")
	ErrNoPayload    = errors.New("Payload не найден")

	// В РАЗРАБОТКЕ
	ErrorInwork = errors.New("В разработке!")

	// ИНИЦИАЛИЗАЦИЯ
	ErrTempDirEmpty = errors.New("os.TempDir() вернул пустой путь")
	ErrInitRoot     = errors.New("Ошибка инициализации!")
)
