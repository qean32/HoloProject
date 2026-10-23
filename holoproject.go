package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"

	"atomicgo.dev/cursor"

	"main/callstack"
	"main/callstack/manual"
	eventlist "main/constants/event-list"
	"main/lib"
	"main/lib/eventlog"
	"main/lib/handler"
	"main/terminal"
)

// version задаётся при сборке: -ldflags "-X main.version=0.1.0".
var version = "dev"

var errInteractive = errors.New("команда требует интерактивного режима")

func main() {
	// root := flag.String("root", "", "каталог данных (по умолчанию %LOCALAPPDATA%\\holoproject)")
	showVersion := flag.Bool("version", false, "показать версию и выйти")
	listMode := flag.Bool("list", false, "показать команды и выйти")
	runLine := flag.String("run", "", "выполнить одну команду без TUI")
	yes := flag.Bool("yes", false, "подтверждать удаление в --run")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "holoproject %s\n\n", version)
		fmt.Fprintln(flag.CommandLine.Output(), `Использование: holoproject [--root <каталог>] [--version] [--list] [--run "<команда>"] [--yes]`)
		fmt.Fprintln(flag.CommandLine.Output())
		flag.PrintDefaults()
	}
	flag.Parse()

	if *showVersion {
		fmt.Printf("holoproject %s (%s/%s, %s)\n", version, runtime.GOOS, runtime.GOARCH, runtime.Version())
		return
	}
	if flag.NArg() > 0 {
		fmt.Fprintf(os.Stderr, "неизвестные аргументы: %v\n", flag.Args())
		flag.Usage()
		os.Exit(2)
	}
	if *listMode && *runLine != "" {
		fmt.Fprintln(os.Stderr, "--list и --run нельзя использовать вместе")
		os.Exit(2)
	}

	// constants.RootOverride = *root

	if runtime.GOOS != "windows" {
		fmt.Fprintln(os.Stderr, "Эта программа поддерживается только на Windows.")
		os.Exit(1)
	}

	if *listMode {
		listCommands()
		return
	}
	if *runLine != "" {
		os.Exit(runHeadless(*runLine, *yes))
	}
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
		os.Exit(1)
	}
}

func listCommands() {
	for _, command := range handler.VisibleCommands() {
		fmt.Printf("%-10s %-34s %s\n", command.Name, command.Title, command.Usage)
	}
}

func runHeadless(line string, yes bool) int {
	err := executeHeadless(line, yes)
	if err == nil {
		return 0
	}

	fmt.Fprintln(os.Stderr, "Ошибка:", err)
	if errors.Is(err, errInteractive) {
		return 2
	}
	return 1
}

func executeHeadless(line string, yes bool) error {
	if err := lib.InitData(); err != nil {
		return err
	}

	args := strings.Fields(line)
	if len(args) == 0 {
		return errors.New("пустая команда")
	}

	event, err := handler.Parse(args)
	if err != nil {
		return err
	}

	// rbat в headless-режиме требует явного --yes: без него это
	// интерактивная команда с подтверждением.
	if yes && event.Key == eventlist.REMOVEBAT {
		if event.SubEvent == nil {
			return errors.New("нечего удалять: нет вложенного события")
		}
		event = *event.SubEvent
	}

	action, err := handler.Dispatch(event)
	if err != nil {
		return err
	}
	if action != nil {
		return errInteractive
	}
	return nil
}

func run() (err error) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	defer signal.Stop(sigCh)

	defer cursor.Show()
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	go func() {
		<-sigCh
		cursor.Show()
		os.Exit(130)
	}()

	if err := lib.Init(); err != nil {
		return err
	}

	runLoop()
	return nil
}

func runLoop() {
	for {
		drain()
		manual.Manual()
	}
}

func drain() {
	for {
		ev, ok := callstack.Pop()
		if !ok {
			return
		}
		if eventlog.Enabled() && ev.Key != eventlist.RESPONSE {
			terminal.PrintTechInfo(ev)
		}
		if err := lib.HandleEvent(ev); err != nil {
			callstack.PushError(ev, err)
		}
	}
}
