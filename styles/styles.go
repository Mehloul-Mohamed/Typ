package styles

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

var termWidth, termHeight, _ = term.GetSize(int(os.Stdout.Fd()))

var Upcoming = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#4C566A"))

var Current = lipgloss.NewStyle().
	Underline(true).
	Foreground(lipgloss.Color("#4C566A"))

var Incorrect = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#BF616A"))

var Correct = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#ECEFF4"))

var Display = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Width(termWidth).
	Height(termHeight)
