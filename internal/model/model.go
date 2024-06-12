package model

import (
	"gosnake/internal/game"
	"gosnake/internal/logger"
	"gosnake/internal/styles"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	TerminalWidth  = 40
	TerminalHeight = 24
	CanvasWidth    = 20
	CanvasHeight   = 20
)

const (
	SnakeChar   = "S"
	AppleChar   = "A"
	NeutralChar = "."
)

type TickMsg time.Time

type Model struct {
	msg    string
	game   *game.Game
	logger *logger.Logger
}

func (m Model) tick() tea.Cmd {
	return tea.Tick(time.Second/5, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func NewModel() Model {
	return Model{
		msg:    "Initializing...",
		game:   game.NewGame(CanvasWidth, CanvasHeight),
		logger: logger.NewLogger("debug.log"),
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
		case "up", "k":
			m.game.Tick(game.Up, m.logger)
			return m, nil
		case "right", "l":
			m.game.Tick(game.Right, m.logger)
			return m, nil
		case "down", "j":
			m.game.Tick(game.Down, m.logger)
			return m, nil
		case "left", "h":
			m.game.Tick(game.Left, m.logger)
			return m, nil
		}

	case TickMsg:
		//m.msg = "timer: " + time.Now().Format("15:04:05")
		//m.msg = m.game.Snake.Dir.String()
		//m.msg = m.game.Snake.Body[0].String()
		m.msg = m.game.StateAsString()

		if m.game.State == game.GameOver {
			return m, nil
		}

		return m, m.tick()
	}

	return m, nil
}

func (m Model) View() string {
	canvas := m.BuildNextCanvasFrame()

	return lipgloss.JoinVertical(
		lipgloss.Center,
		styles.Button(m.msg),
		canvas,
	)
}

func (m Model) BuildNextCanvasFrame() string {
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
