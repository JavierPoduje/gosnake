package ui

import (
	"fmt"
	"gosnake/internal/game"
	"log"

	"github.com/charmbracelet/lipgloss"
)

func Canvas(width, height int, state int, content string) string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		CanvasLabel(state),
		CanvasStyles(width, height, state).Render(content),
	)
}

func CanvasLabel(state int) string {
	var label string
	switch state {
	case game.Running:
		label = "Running"
	case game.GameOver:
		label = "Game Over"
	case game.Paused:
		label = "Paused"
	default:
		log.Panic("Unknown game state")
	}

	return CanvasLabelStyles(state).Render(label)
}

func Layout(width, height int, content ...string) string {
	return lipgloss.Place(
		width, height,
		lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center, content...),
	)
}

func Apple(char string) string {
	return lipgloss.NewStyle().
		Foreground(RedColor()).
		Render(char)
}

func HelpContainer(keysAsString string) string {
	return HelpContainerStyles().Render(keysAsString)
}

func SnakeHead(char string) string {
	return lipgloss.NewStyle().
		Foreground(GreenColor()).
		Render(char)
}

func SnakeBody(char string) string {
	return lipgloss.NewStyle().
		Foreground(PrimaryTextColor()).
		Render(char)
}

func NeutralChar(char string) string {
	return lipgloss.NewStyle().
		Width(1).
		Height(1).
		Foreground(NeutralColor()).
		Render(char)
}

func StatsCard(stats [][]string) string {
	title := "Stats"

	var headersColumn []string
	for _, stat := range stats {
		header := StatHeaderStyles().Render(stat[0])
		headersColumn = append(headersColumn, header)
	}
	styledHeader := lipgloss.JoinVertical(lipgloss.Left, headersColumn...)

	var valuesColumn []string
	for _, stat := range stats {
		value := StatValueStyles().Render(stat[1])
		valuesColumn = append(valuesColumn, value)
	}
	styledValues := lipgloss.JoinVertical(lipgloss.Right, valuesColumn...)

	return StatsStyles().Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			TitleStyles().Render(title),
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				styledHeader,
				styledValues,
			),
		),
	)
}

func HistoricScoresCard(scores []int) string {
	title := "History"

	numberOfScoresToDisplay := 10

	var scorePosition []string
	var scoreValues []string
	for i := 0; i < numberOfScoresToDisplay; i++ {
		posStr := fmt.Sprintf("%d. ", i+1)

		scoreStr := ""
		if i < len(scores) {
			scoreStr = fmt.Sprintf("%d", scores[i])
		}

		scorePosition = append(scorePosition, HistoricScoresPositionStyles().Render(posStr))
		scoreValues = append(scoreValues, HistoricScoresValueStyles().Render(scoreStr))
	}
	styledPositions := lipgloss.JoinVertical(lipgloss.Right, scorePosition...)
	stylesScores := lipgloss.JoinVertical(lipgloss.Right, scoreValues...)

	return HistoricScoresStyles().Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			TitleStyles().Render(title),
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				styledPositions,
				stylesScores,
			),
		),
	)
}
