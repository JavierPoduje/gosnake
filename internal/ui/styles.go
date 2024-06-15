package ui

import (
	"gosnake/internal/game"
	"log"

	"github.com/charmbracelet/lipgloss"
)

func CanvasStyles(width, height int) lipgloss.Style {
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

func StatHeaderStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(PurpleColor()).
		Align(lipgloss.Left).
		PaddingRight(3).
		Height(1)
}

func StatValueStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(WhiteColor()).
		Align(lipgloss.Right).
		Height(1)
}

func StatsStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(WhiteColor()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(PurpleColor()).
		Align(lipgloss.Center, lipgloss.Center).
		Width(22).
		Height(6)
}

func ActionButtonStyles(state int) lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(WhiteColor()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(colorByState(state)).
		Align(lipgloss.Center, lipgloss.Center).
		Width(22).
		Height(3)
}

// HELPERS

func colorByState(state int) lipgloss.Color {
	switch state {
	case game.Running:
		return GreenColor()
	case game.GameOver:
		return RedColor()
	case game.Paused:
		return OrangeColor()
	default:
		log.Panic("Unknown game state")
	}

	return lipgloss.Color("#000")
}
