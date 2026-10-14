package test

import (
	"fmt"
	"strconv"
	"strings"

	"main/constants/sgr"
	"main/lib/utils"
	"main/terminal"
)

// status — результат подтеста.
type status int

const (
	statusPass status = iota
	statusFail
	statusSkip
)

// String возвращает подпись статуса для вывода.
func (s status) String() string {
	switch s {
	case statusFail:
		return utils.GetCustomMessage("FAIL", sgr.RED, sgr.BOLD)
	case statusSkip:
		return utils.GetCustomMessage("SKIP", sgr.YELLOW)
	default:
		return utils.GetCustomMessage("PASS", sgr.GREEN, sgr.BOLD)
	}
}

// subCase — один подтест: группа, имя и проверка.
type subCase struct {
	group string
	name  string
	run   func() error
}

// report — результат выполнения одного подтеста.
type report struct {
	group  string
	name   string
	status status
	detail string
}

// runAll выполняет подтесты всех групп, печатает результат каждого
// и возвращает отчёты для итоговой сводки.
func runAll(groups ...[]subCase) []report {
	var reports []report
	for _, group := range groups {
		for _, c := range group {
			reports = append(reports, run(c))
		}
	}
	return reports
}

// run выполняет один подтест и печатает его результат.
// Паника внутри подтеста не роняет программу, а проваливает подтест.
func run(c subCase) report {
	r := report{group: c.group, name: c.name, status: statusPass}

	func() {
		defer func() {
			if err := recover(); err != nil {
				r.status = statusFail
				r.detail = fmt.Sprintf("panic: %v", err)
			}
		}()

		if err := c.run(); err != nil {
			r.status = statusFail
			r.detail = err.Error()
		}
	}()

	printReport(r)
	return r
}

// printReport выводит результат одного подтеста.
func printReport(r report) {
	name := utils.GetCustomMessage(r.group+"/"+r.name, sgr.DIM)
	if r.detail == "" {
		terminal.Println(fmt.Sprintf(" [ %s ] %s", r.status, name))
		return
	}
	terminal.Println(fmt.Sprintf(" [ %s ] %s — %s", r.status, name, r.detail))
}

// summary выводит итоговую сводку по всем подтестам.
func summary(reports []report) {
	passed, failed, skipped := 0, 0, 0
	for _, r := range reports {
		switch r.status {
		case statusFail:
			failed++
		case statusSkip:
			skipped++
		default:
			passed++
		}
	}

	line := utils.GetCustomMessage(strings.Repeat("─", 52), sgr.DIM)
	terminal.Println(line)
	terminal.Println(fmt.Sprintf(" всего: %s, пройдено: %s, провалено: %s, пропущено: %s",
		strconv.Itoa(len(reports)),
		utils.GetCustomMessage(strconv.Itoa(passed), sgr.GREEN),
		utils.GetCustomMessage(strconv.Itoa(failed), sgr.RED),
		utils.GetCustomMessage(strconv.Itoa(skipped), sgr.YELLOW),
	))

	if failed == 0 {
		return
	}

	terminal.Println(utils.GetCustomMessage(" провалено:", sgr.RED, sgr.BOLD))
	for _, r := range reports {
		if r.status != statusFail {
			continue
		}
		terminal.Println(fmt.Sprintf("   %s — %s", r.group+"/"+r.name, r.detail))
	}
}

// failed возвращает ошибку, если среди отчётов есть проваленные подтесты.
func failed(reports []report) error {
	for _, r := range reports {
		if r.status == statusFail {
			return errTestsFailed
		}
	}
	return nil
}
