package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	appNameStyle    = lipgloss.NewStyle().Background(lipgloss.Color("99")).Padding(0, 1)
	faintStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Faint(true)
	enumeratorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("99")).MarginRight(1)
)

func (m model) View() string {
	s := appNameStyle.Render(("Notes App")) + "\n\n"

	if m.state == titleView {
		s += "Note title:\n\n"
		s += m.textinput.View() + "\n\n"
		s += faintStyle.Render("enter - save, esc - discard")
	}

	if m.state == bodyView {
		s += "Notes:\n\n"
		s += m.textarea.View() + "\n\n"
		s += faintStyle.Render("ctrl+s - save, esc - discard")
	}

	if m.state == listView {
		for index, note := range m.notes {

			prefix := " "
			if index == m.listIndex {
				prefix = ">"
			}

			shortBody := strings.ReplaceAll(note.Body, "\n", " ")
			if len(shortBody) > 30 {
				shortBody = shortBody[:30]
			}
			s += enumeratorStyle.Render(prefix) + note.Title + " | " + faintStyle.Render(shortBody) + "\n\n"
		}
		s += faintStyle.Render("n - new note, q -  quit")
	}

	return s
}
