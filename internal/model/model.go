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

const defaultSnakeDir = game.Right

type Model struct {
	msg           string
	game          *game.Game
	logger        *logger.Logger
	nextSnakeMove game.Direction
}

func (m Model) tick() tea.Cmd {
	return tea.Tick(time.Second/3, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func NewModel() Model {
	return Model{
		msg:           "Initializing...",
		game:          game.NewGame(CanvasWidth, CanvasHeight),
		logger:        logger.NewLogger("debug.log"),
		nextSnakeMove: defaultSnakeDir,
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
			if !m.game.Snake.IsOppositeDir(game.Up) {
				m.nextSnakeMove = game.Up
			}
			return m, nil
		case "right", "l":
			if !m.game.Snake.IsOppositeDir(game.Right) {
				m.nextSnakeMove = game.Right
			}
			return m, nil
		case "down", "j":
			if !m.game.Snake.IsOppositeDir(game.Down) {
				m.nextSnakeMove = game.Down
			}
			return m, nil
		case "left", "h":
			if !m.game.Snake.IsOppositeDir(game.Left) {
				m.nextSnakeMove = game.Left
			}
			return m, nil
		}

	case TickMsg:
		// update message displayed in the up-button
		m.msg = m.game.StateAsString()

		// update the game state
		m.game.Tick(m.nextSnakeMove, m.logger)

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
