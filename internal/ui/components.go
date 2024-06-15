package ui

import (
	"github.com/charmbracelet/lipgloss"
)

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

func StatCard(stats [][]string) string {
	var headersColumn []string
	for _, stat := range stats {
		header := StatHeader().Render(stat[0])
		headersColumn = append(headersColumn, header)
	}
	styledHeader := lipgloss.JoinVertical(lipgloss.Left, headersColumn...)

	var valuesColumn []string
	for _, stat := range stats {
		value := StatValue().Render(stat[1])
		valuesColumn = append(valuesColumn, value)
	}
	styledValues := lipgloss.JoinVertical(lipgloss.Right, valuesColumn...)

	return Stats().Render(
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			styledHeader,
			styledValues,
		),
	)
}
