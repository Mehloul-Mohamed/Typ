package main

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/Mehloul-Mohamed/typ/styles"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var startTime = time.Now()

var termWidth, termHeight, _ = term.GetSize(int(os.Stdout.Fd()))

type model struct {
	message         string
	correct_message string
	error_count     int
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "backspace":
			if len(m.message) > 0 {
				return model{message: m.message[:len(m.message)-1], correct_message: m.correct_message, error_count: countErrors(m.correct_message, m.message[:len(m.message)-1])}, nil
			}
		case "up":
			return m, nil
		case "down":
			return m, nil
		case "left":
			return m, nil
		case "right":
			return m, nil
		case "enter":
			return m, nil
		case "esc":
			return m, nil
		default:
			if len(m.message) >= len(m.correct_message) {
				return m, tea.Quit
			}
			return model{message: m.message + msg.String(), correct_message: m.correct_message, error_count: countErrors(m.correct_message, m.message+msg.String())}, nil
		}
	}
	return m, nil
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) View() string {
	out := ""
	for i, c := range m.correct_message {
		if i < len(m.message) {
			if m.message[i] == m.correct_message[i] {
				out += styles.Correct.Render(string(c))
			} else {
				out += styles.Incorrect.Render(string(c))
			}
		} else if i == len(m.message) {
			out += styles.Current.Render(string(c))
		} else {
			out += styles.Upcoming.Render(string(c))
		}
	}
	out += fmt.Sprintf("\n\n\n%d WPM", m.error_count)
	return styles.Display.Render(lipgloss.Place(termWidth, termHeight, lipgloss.Center, lipgloss.Center, out))
}

func countErrors(c string, m string) int {
	e := 0
	if len(c) == 0 || len(m) == 0 {
		return 0
	}
	for i := 0; i <= min(len(c), len(m))-1; i++ {
		if c[i] != m[i] {
			e += 1
		}
	}
	dir := (time.Now().Sub(startTime)).Minutes()
	if dir == 0 {
		return 0
	}
	return int(math.Ceil((float64((len(m) / 5) - e)) / dir))
}

func main() {
	txt := "For example, this code always computes a positive elapsed time of approximately 20 millisecondsOperating systems provide both a “wall clock,” which is subject to changes for clock synchronization, and a “monotonic clock,” which is not. The general rule is that the wall clock is for telling time and the monotonic clock is for measuring time. Rather than split the API, in this package the Time returned by time.Now contains both a wall clock reading and a monotonic clock reading; later time-telling operations use the wall clock reading, but later time-measuring operations, specifically comparisons and subtractions, use the monotonic clock reading."
	p := tea.NewProgram(model{message: "", correct_message: txt}, tea.WithAltScreen())
	p.Run()
}
