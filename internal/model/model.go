package model

import (
	"gosnake/internal/game"
	"gosnake/internal/styles"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	TerminalWidth  = 97
	TerminalHeight = 34
	CanvasWidth    = 10
	CanvasHeight   = 10
)

const (
	SnakeChar   = "S"
	AppleChar   = "A"
	NeutralChar = "."
)

type TickMsg time.Time

type Model struct {
	msg  string
	game *game.Game
}

func (m Model) tick() tea.Cmd {
	return tea.Tick(time.Second/5, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func NewModel() Model {
	return Model{
		msg:  "Initializing...",
		game: game.NewGame(CanvasWidth, CanvasHeight),
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
	canvas := m.BuildCanvas()
	return canvas
}

func (m Model) BuildCanvas() string {
	strCanvas := strings.Builder{}

	width := m.game.Canvas.Width
	height := m.game.Canvas.Height

	snake := m.game.Snake
	apple := m.game.Apple

	for Y := 0; Y < width; Y++ {
		for X := 0; X < height; X++ {
			if snake.Contains(game.Coord{X: X, Y: Y}) {
				strCanvas.WriteString(SnakeChar)
			} else if apple.X == X && apple.Y == Y {
				strCanvas.WriteString(AppleChar)
			} else {
				strCanvas.WriteString(styles.NeutralChar().Render(NeutralChar))
			}
		}
	}

	canvas := lipgloss.Place(
		TerminalWidth, TerminalHeight,
		lipgloss.Center, lipgloss.Center,
		styles.Canvas(CanvasWidth, CanvasHeight).Render(strCanvas.String()),
	)

	return canvas
}
