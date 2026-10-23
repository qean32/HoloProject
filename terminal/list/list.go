package list

import (
	"fmt"
	"strings"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	"github.com/nathan-fiscaletti/consolesize-go"

	"main/callstack/event"
	"main/constants/sgr"
	"main/lib/utils"
	"main/model"
)

type state struct {
	title    string
	all      []model.Option
	filtered []int
	query    string
	position int
	offset   int
}

// List показывает список в отдельном экране и возвращает выбранный пункт.
// Второе значение — false, если пользователь вышел без выбора.
func List(payload model.ListPayload) (model.Event, bool) {
	if len(payload.Options) == 0 {
		event.Response("length == zero! ave caesar! ")
		return model.Event{}, false
	}

	s := &state{title: payload.Title, all: payload.Options}
	s.refilter()
	defer s.leave()
	s.enter()

	return s.run()
}

func (s *state) run() (model.Event, bool) {
	var selected model.Event
	chosen := false

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.Enter:
			option, ok := s.current()
			if !ok {
				return false, nil
			}
			selected = utils.GetShortEvent(option.Event)
			chosen = true
			return true, nil

		case keys.Escape, keys.CtrlC:
			if key.Code == keys.Escape && s.query != "" {
				s.query = ""
				s.refilter()
				s.draw()
				return false, nil
			}
			return true, nil

		case keys.Up:
			s.move(-1)
			s.draw()
		case keys.Down:
			s.move(1)
			s.draw()
		case keys.PgUp:
			s.move(-s.pageSize())
			s.draw()
		case keys.PgDown:
			s.move(s.pageSize())
			s.draw()
		case keys.Home:
			s.position = 0
			s.draw()
		case keys.End:
			if len(s.filtered) > 0 {
				s.position = len(s.filtered) - 1
			}
			s.draw()

		case keys.Backspace, keys.CtrlH:
			s.dropRune()
			s.draw()

		case keys.RuneKey:
			if len(key.Runes) > 0 && !key.AltPressed && key.Runes[0] >= ' ' {
				s.query += string(key.Runes)
				s.refilter()
				s.draw()
			}
		}

		return false, nil
	})

	return selected, chosen
}

func (s *state) current() (model.Option, bool) {
	if s.position < 0 || s.position >= len(s.filtered) {
		return model.Option{}, false
	}
	return s.all[s.filtered[s.position]], true
}

func (s *state) refilter() {
	s.filtered = filterIndices(s.all, s.query)
	s.position = 0
	s.offset = 0
}

func (s *state) move(delta int) {
	if len(s.filtered) == 0 {
		return
	}
	position := s.position + delta
	if position < 0 {
		position = 0
	}
	if position >= len(s.filtered) {
		position = len(s.filtered) - 1
	}
	s.position = position
}

func (s *state) dropRune() {
	runes := []rune(s.query)
	if len(runes) == 0 {
		return
	}
	s.query = string(runes[:len(runes)-1])
	s.refilter()
}

func (s *state) ensureVisible(page int) {
	if page < 1 {
		page = 1
	}
	if len(s.filtered) == 0 {
		s.position = 0
		s.offset = 0
		return
	}
	if s.position >= len(s.filtered) {
		s.position = len(s.filtered) - 1
	}
	if s.position < 0 {
		s.position = 0
	}

	if s.position < s.offset {
		s.offset = s.position
	}
	if s.position >= s.offset+page {
		s.offset = s.position - page + 1
	}

	maxOffset := len(s.filtered) - page
	if maxOffset < 0 {
		maxOffset = 0
	}
	if s.offset > maxOffset {
		s.offset = maxOffset
	}
	if s.offset < 0 {
		s.offset = 0
	}
}

func (s *state) enter() {
	fmt.Print("\x1b[?1049h\x1b[?25l")
	s.draw()
}

func (s *state) leave() {
	fmt.Print("\x1b[?25h\x1b[?1049l")
}

func (s *state) draw() {
	width, height := size()
	lines := s.frame(width, height)
	fmt.Print("\x1b[H" + strings.Join(lines, "\r\n") + "\x1b[J")
}

func (s *state) frame(width, height int) []string {
	if width < 20 {
		width = 20
	}
	if height < 4 {
		height = 4
	}

	page := height - 3
	if page < 1 {
		page = 1
	}
	s.ensureVisible(page)

	contentWidth := width
	if contentWidth > 1 {
		contentWidth--
	}

	lines := make([]string, 0, height)
	header := s.title
	if header == "" {
		header = "Выбор"
	}
	header = fmt.Sprintf("%s  [%d/%d]", header, len(s.filtered), len(s.all))
	if s.query != "" {
		header += "  фильтр: " + s.query
	}
	lines = append(lines, utils.GetCustomMessage(truncate(header, contentWidth), sgr.BOLD))
	lines = append(lines, utils.GetCustomMessage(strings.Repeat("─", contentWidth), sgr.DIM))

	if len(s.filtered) == 0 {
		lines = append(lines, utils.GetCustomMessage("  ничего не найдено", sgr.DIM))
	} else {
		end := s.offset + page
		if end > len(s.filtered) {
			end = len(s.filtered)
		}
		for i := s.offset; i < end; i++ {
			index := s.filtered[i]
			marker := "  "
			if i == s.position {
				marker = "> "
			}
			line := truncate(fmt.Sprintf("%s%3d. %s", marker, index+1, s.all[index].Message), contentWidth)
			if i == s.position {
				line = utils.GetCustomMessage(line, sgr.GREEN)
			}
			lines = append(lines, line)
		}
	}

	for len(lines) < height-1 {
		lines = append(lines, "")
	}
	footer := "↑/↓  PgUp/PgDn  Enter — выбрать  Esc — сброс/выход  ввод — фильтр"
	lines = append(lines, utils.GetCustomMessage(truncate(footer, contentWidth), sgr.DIM))

	return lines
}

func (s *state) pageSize() int {
	_, height := size()
	page := height - 3
	if page < 1 {
		page = 1
	}
	return page
}

func size() (int, int) {
	width, height := consolesize.GetConsoleSize()
	if width <= 0 {
		width = 80
	}
	if height <= 0 {
		height = 24
	}
	return width, height
}

func truncate(value string, width int) string {
	if width <= 0 {
		return ""
	}
	runes := []rune(value)
	if len(runes) <= width {
		return value
	}
	if width == 1 {
		return string(runes[:1])
	}
	return string(runes[:width-1]) + "…"
}
