package ui

import (
	"gosnake/internal/game"
	"log"

	"github.com/charmbracelet/lipgloss"
)

func TitleStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(PinkColor()).
		Align(lipgloss.Center).
		Height(1).
		MarginBottom(1)
}

func CanvasStyles(width, height int, state int) lipgloss.Style {
	var borderColor lipgloss.Color
	switch state {
	case game.Running:
		borderColor = GreenColor()
	case game.GameOver:
		borderColor = RedColor()
	case game.Paused:
		borderColor = OrangeColor()
	default:
		log.Panic("Unknown game state")
	}

	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(borderColor).
		MarginRight(2).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true)
}

func HelpContainerStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Height(4).
		Width(35).
		MarginTop(1)
}

func StatHeaderStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(PinkColor()).
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

func HistoricScoresStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(PurpleColor()).
		Align(lipgloss.Center, lipgloss.Center).
		Width(22).
		Height(12)
}

func HistoricScoresPositionStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(PinkColor()).
		Height(1)
}

func HistoricScoresValueStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(WhiteColor()).
		Height(1)
}
