// Package test содержит встроенные проверки проекта.
// Запускается по событию eventlist.TEST (команда "test").
package test

import (
	"main/constants/sgr"
	"main/lib/utils"
	"main/terminal"
)

// Test выполняет подтесты по всем Event из eventlist и по всем парсерам,
// печатает результат каждого подтеста и итоговую сводку.
//
// Обработчики событий (main/lib/handler) не вызываются намеренно:
// пакет handler импортирует этот пакет, поэтому вызов обработчиков
// создал бы циклическую зависимость. Проверяется контракт события:
// разбор команды ручного ввода и прохождение через callstack.
func Test() error {
	defer drain()

	terminal.Println(utils.GetCustomMessage(" unit-тесты", sgr.BOLD, sgr.UNDERLINE))

	reports := runAll(eventCases(), parserSubCases())
	summary(reports)

	return failed(reports)
}
