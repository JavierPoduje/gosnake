package ui

import (
	"gosnake/internal/game"
	"log"

	"github.com/charmbracelet/lipgloss"
)

func TitleStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(SecondaryTextColor()).
		Align(lipgloss.Center).
		Height(1).
		MarginBottom(1)
}

func CanvasLabelStyles(state int) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(ColorByStateStyles(state)).
		Height(1).
		MarginRight(2)
}

func ColorByStateStyles(state int) lipgloss.Color {
	var color lipgloss.Color
	switch state {
	case game.Running:
		color = GreenColor()
	case game.GameOver:
		color = RedColor()
	case game.Paused:
		color = OrangeColor()
	default:
		log.Panic("Unknown game state")
	}
	return color
}

func CanvasStyles(width, height int, state int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorByStateStyles(state)).
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
		Foreground(SecondaryTextColor()).
		Align(lipgloss.Left).
		PaddingRight(3).
		Height(1)
}

func StatValueStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(PrimaryTextColor()).
		Align(lipgloss.Right).
		Height(1)
}

func StatsStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		Foreground(PrimaryTextColor()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(GreyColor()).
		Align(lipgloss.Center, lipgloss.Center).
		Width(22).
		MarginTop(1).
		Height(6)
}

func HistoricScoresStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Bold(true).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(GreyColor()).
		Align(lipgloss.Center, lipgloss.Center).
		Width(22).
		Height(12)
}

func HistoricScoresPositionStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(SecondaryTextColor()).
		Height(1)
}

func HistoricScoresValueStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(PrimaryTextColor()).
		Height(1)
}
