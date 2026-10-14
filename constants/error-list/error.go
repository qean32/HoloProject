package errorlist

import (
	"errors"
	"fmt"
)

var (
	// ОБЩИЕ
	ErrAccessDenied   = errors.New("Процесс остановлен")
	ErrNotExist       = errors.New("файл не существует")
	ErrIsDir          = errors.New("Это директория")
	ErrInvalidName    = errors.New("Некорректное имя")
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

	// ТЕСТЫ
	ErrTestNotConnected = errors.New("Набор тестов не подключён")
)

// InvalidName формирует ошибку некорректного имени,
// добавляя проблемное имя name и оборачивая причину err.
func InvalidName(name string, err error) error {
	return fmt.Errorf("Некорректное имя %q: %w", name, err)
}

// Create — ошибка создания объекта с именем name.
func Create(name string, err error) error {
	return fmt.Errorf("Создание %s: %w", name, err)
}

// Run — ошибка запуска объекта с именем name.
func Run(name string, err error) error {
	return fmt.Errorf("Запуск %s: %w", name, err)
}

// Remove — ошибка удаления объекта с именем name.
func Remove(name string, err error) error {
	return fmt.Errorf("Удаление %s: %w", name, err)
}

// Read — ошибка чтения объекта с именем name.
func Read(name string, err error) error {
	return fmt.Errorf("Чтение %s: %w", name, err)
}

// Open — ошибка открытия объекта с именем name.
func Open(name string, err error) error {
	return fmt.Errorf("Сткрытие %s: %w", name, err)
}

// ClearLog — ошибка очистки лога.
func ClearLog(err error) error {
	return fmt.Errorf("Счистка лога: %w", err)
}

// RunCommand — ошибка запуска команды command.
func RunCommand(command string, err error) error {
	return fmt.Errorf("Запуск команды %q: %w", command, err)
}

// ExecCommand — ошибка выполнения команды command.
func ExecCommand(command string, err error) error {
	return fmt.Errorf("Выполнение команды %q: %w", command, err)
}

// ClearTerminal — ошибка очистки терминала.
func ClearTerminal(err error) error {
	return fmt.Errorf("Очистка терминала: %w", err)
}

// OpenFile — ошибка открытия файла path.
func OpenFile(path string, err error) error {
	return fmt.Errorf("Открытие %s: %w", path, err)
}

// ReadFile — ошибка чтения файла/каталога path.
func ReadFile(path string, err error) error {
	return fmt.Errorf("Чтение %s: %w", path, err)
}

// CreateFile — ошибка создания файла path.
func CreateFile(path string, err error) error {
	return fmt.Errorf("Создание %s: %w", path, err)
}

// WriteFile — ошибка записи в файл path.
func WriteFile(path string, err error) error {
	return fmt.Errorf("Запись %s: %w", path, err)
}

// RemoveFile — ошибка удаления файла path.
func RemoveFile(path string, err error) error {
	return fmt.Errorf("Удаление %s: %w", path, err)
}

// WithContext оборачивает ошибку err, добавляя контекст context.
func WithContext(context string, err error) error {
	return fmt.Errorf("%w: %s", err, context)
}
