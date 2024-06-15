package ui

import (
	"gosnake/internal/game"
	"log"

	"github.com/charmbracelet/lipgloss"
)

func Canvas(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#fff")).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
}

func StatHeader() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(Purple()).
		Align(lipgloss.Left).
		PaddingRight(3).
		Height(1)
}

func StatValue() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(White()).
		Align(lipgloss.Right).
		Height(1)
}

func Stats() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(White()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(Purple()).
		Align(lipgloss.Center, lipgloss.Center).
		Width(22).
		Height(6)
}

func ActionButton(state int) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(White()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(colorByState(state)).
		Align(lipgloss.Center, lipgloss.Center).
		Width(22).
		Height(3)
}

func colorByState(state int) lipgloss.Color {
	switch state {
	case game.Running:
		return Green()
	case game.GameOver:
		return Red()
	case game.Paused:
		return Orange()
	default:
		log.Panic("Unknown game state")
	}

	return lipgloss.Color("#000")
}
