package ui

import (
	"github.com/charmbracelet/lipgloss"
)

func CanvasStyles(width, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(PurpleColor()).
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
		Foreground(PurpleColor()).
		Align(lipgloss.Left).
		PaddingRight(3).
		Height(1)
}

func StatValueStyles() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(RedColor()).
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
		Foreground(RedColor()).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(PurpleColor()).
		Align(lipgloss.Center, lipgloss.Center).
		Width(22).
		Height(3)
}
