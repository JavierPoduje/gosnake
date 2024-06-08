package model

import (
	"gosnake/internal/snake"
	"gosnake/internal/ui"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	TerminalWidth  = 96
	TerminalHeight = 33
	CanvasWidth    = 10
	CanvasHeight   = 10
)

const (
	SnakeChar = "S"
	AppleChar = "A"
)

type TickMsg time.Time

type Model struct {
	msg string
}

func (m Model) tick() tea.Cmd {
	return tea.Tick(time.Second/5, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func NewModel() Model {
	return Model{
		msg: "Initializing...",
	}
}

func (m Model) Init() tea.Cmd {
	return m.tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case TickMsg:
		m.msg = "timer: " + time.Now().Format("15:04:05")
		return m, m.tick()
	}

	return m, nil
}

func (m Model) View() string {
	strCanvas := strings.Builder{}
	s := snake.NewSnake()
	appleCoord := snake.Coord{X: 2, Y: 2}

	for y := 0; y < CanvasWidth; y++ {
		for x := 0; x < CanvasHeight; x++ {
			if s.Contains(snake.Coord{X: x, Y: y}) {
				strCanvas.WriteString(SnakeChar)
			} else if appleCoord.X == x && appleCoord.Y == y {
				strCanvas.WriteString(AppleChar)
			} else {
				strCanvas.WriteString(ui.NeutralChar().Render("."))
			}
		}
	}

	filledCanvas := lipgloss.Place(
		TerminalWidth, TerminalHeight,
		lipgloss.Center, lipgloss.Center,
		ui.Canvas(CanvasWidth, CanvasHeight).Render(strCanvas.String()),
		lipgloss.WithWhitespaceChars("\uef0f"),
		lipgloss.WithWhitespaceForeground(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}),
	)

	return filledCanvas
}
