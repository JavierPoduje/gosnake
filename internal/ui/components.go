package ui

import "github.com/charmbracelet/lipgloss"

func Layout(width, height int, content string) string {
	return lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		content,
	)
}

func Apple(char string) string {
	return lipgloss.NewStyle().
		Foreground(Red()).
		Render(char)
}

func Snake(char string) string {
	return lipgloss.NewStyle().
		Foreground(Yellow()).
		Render(char)
}

func NeutralChar(char string) string {
	return lipgloss.NewStyle().
		Width(1).
		Height(1).
		Foreground(Neutral()).
		Render(char)
}
