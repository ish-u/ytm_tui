package tui

import (
	"fmt"
	"strings"
)

// View Method
func (m Model) View() string {
	// The header
	s := ""
	if !m.queue.Loading {
		s += fmt.Sprintf("\n%s\t", (m.textInput.View()))
	}

	if m.queue.Loading {
		s += fmt.Sprintf("\nAdding - %s to the Queue  %d / %d\t", m.enteredText, m.queue.Start, m.queue.End)
		s += m.spinner.View()
	}

	s += "\n\n"
	if len(m.queue.Videos) > 0 {
		var idx = 0
		s += "Queue\n\n"
		for idx < 30 {
			if idx == 0 {
				if !m.queue.Paused {
					s += m.spinner.View()
				} else {
					s += "~ "
				}
				if len(m.queue.Videos[idx+m.queue.Current].Title) > 60 {
					s += fmt.Sprintf("%10s", m.queue.Videos[idx+m.queue.Current].Title[:60])
				} else {
					s += fmt.Sprintf("%10s", m.queue.Videos[idx+m.queue.Current].Title)
					s += strings.Repeat(" ", 60-len(m.queue.Videos[idx+m.queue.Current].Title))
				}
				s += fmt.Sprintf("\t\t%s - ", m.position)
				s += m.queue.Videos[m.queue.Current].Duration.String()
				s += "\n"
			} else if (idx + m.queue.Current) < len(m.queue.Videos) {
				s += fmt.Sprintf("%d ", idx+1+m.queue.Current)
				if len(m.queue.Videos[idx+m.queue.Current].Title) > 60 {
					s += fmt.Sprintf("%10s", m.queue.Videos[idx+m.queue.Current].Title[:60])
				} else {
					s += fmt.Sprintf("%10s", m.queue.Videos[idx+m.queue.Current].Title)
					s += strings.Repeat(" ", 60-len(m.queue.Videos[idx+m.queue.Current].Title))
				}
				s += "\t\t     "
				s += m.queue.Videos[idx+m.queue.Current].Duration.String()
				s += "\n"

			} else {
				s += "\n"
			}

			idx++
		}
		s += "\n"

	}

	helpView := m.help.FullHelpView(m.keys.FullHelp())

	height := 2 - strings.Count(helpView, "\n")

	s += strings.Repeat("\n", height) + helpView

	// Send the UI for rendering
	return s
}
