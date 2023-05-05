package tui

import (
	"fmt"

	"github.com/NDOY3M4N/gopass/internal/password"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	listLength         []itemLength
	cursorListLength   int
	selectedItemLength int

	listChars         []itemChars
	cursorListChars   int
	selectedItemChars map[string]bool

	quitting          bool
	nextView          bool
	selectionComplete bool
}

// NOTE: should I transform this into an interface?
type itemLength struct {
	label string
	value int
}

func newListLength() []itemLength {
	return []itemLength{
		{label: "8 characters", value: 8},
		{label: "12 characters", value: 12},
		{label: "16 characters", value: 16},
		{label: "20 characters", value: 20},
	}
}

type itemChars struct {
	label string
	value bool
}

func newListChars() []itemChars {
	return []itemChars{
		{label: "Include lowercases", value: true},
		{label: "Include uppercases", value: false},
		{label: "Include numbers", value: false},
		{label: "Include symbols", value: false},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()

		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.quitting = true
			return m, tea.Quit
		}
	}

	if !m.nextView {
		return updateLength(msg, m)
	}
	return updateChars(msg, m)
}

func updateLength(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.cursorListLength < len(m.listLength)-1 {
				m.cursorListLength++
			}
		case "k", "up":
			if m.cursorListLength > 0 {
				m.cursorListLength--
			}
		case " ":
			m.selectedItemLength = m.listLength[m.cursorListLength].value
		case "enter":
			m.nextView = true
			return m, nil
		}
	}

	return m, nil
}

func updateChars(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.cursorListChars < len(m.listChars)-1 {
				m.cursorListChars++
			}
		case "k", "up":
			if m.cursorListChars > 0 {
				m.cursorListChars--
			}
		case " ":
			m.listChars[m.cursorListChars].value = !m.listChars[m.cursorListChars].value
		case "enter":
			var numSelectedItem int
			for _, item := range m.listChars {
				if item.value {
					numSelectedItem++
				}
			}

			// NOTE: If we have just one selected item, we can't deselect it
			if numSelectedItem > 0 {
				m.quitting = true
				m.selectionComplete = true
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "\n\nGoPass v1.0.0\n"
	s += "=======================\n\n"

	if m.quitting {
		var msg string
		if m.selectionComplete {
			o := password.Option{
				Length: m.selectedItemLength,
				// PERF: refactor this monstrosity
				HasLowercase: m.listChars[0].value,
				HasUppercase: m.listChars[1].value,
				HasNumber:    m.listChars[2].value,
				HasSymbol:    m.listChars[3].value,
			}
			pwd, score := password.Generate(o)

			msg = fmt.Sprintf("The password is: %s and the score is: %s\n\n", pwd, score)
		} else {
			msg = "No selection?\n"
		}

		return s + msg + "See you, space cowboy...\n\n"
	}

	if !m.nextView {
		s += lengthView(m)
	} else {
		s += charsView(m)
	}

	// NOTE: help text here
	s += "down/j: up - down/j: down - space: select item - enter: confirm selection"

	return s
}

func lengthView(m model) string {
	s := "Choose the character length\n\n"

	for index, i := range m.listLength {
		var cursor string
		if index == m.cursorListLength {
			cursor = ">"
		} else {
			cursor = " "
		}

		var check string
		if m.selectedItemLength == i.value {
			check = "[x]"
		} else {
			check = "[ ]"
		}

		s += fmt.Sprintf("%s %s %s\n", cursor, check, i.label)
	}

	return s + "\n\n"
}

func charsView(m model) string {
	s := "Choose the characters to include\n\n"
	for index, i := range m.listChars {
		var cursor string
		if index == m.cursorListChars {
			cursor = ">"
		} else {
			cursor = " "
		}

		var check string
		if i.value {
			check = "[x]"
		} else {
			check = "[ ]"
		}

		s += fmt.Sprintf("%s %s %s\n", cursor, check, i.label)
	}

	return s + "\n\n"
}

func NewModel() model {
	return model{
		selectedItemLength: 8,
		listChars:          newListChars(),
		listLength:         newListLength(),
	}
}
