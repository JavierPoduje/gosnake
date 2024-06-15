package model

import (
	"fmt"
	"gosnake/internal/game"
	"gosnake/internal/logger"
	"gosnake/internal/ui"
	"log"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	DefaultTerminalWidth  = 40
	DefaultTerminalHeight = 24
	CanvasWidth           = 20
	CanvasHeight          = 20
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
	width         int
	height        int
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
		width:         DefaultTerminalWidth,
		height:        DefaultTerminalHeight,
	}
}

func (m Model) Init() tea.Cmd {
	return m.tick()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.logger.Log(fmt.Sprintf("width: %d, height: %d", msg.Width, msg.Height))
		m.width, m.height = msg.Width, msg.Height

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
		// update the game state
		m.game.Tick(m.nextSnakeMove, m.logger)

		m.msg = m.getActionButtonLabel()

		if m.game.State == game.Running {
			return m, m.tick()
		}

		return m, nil
	}

	return m, nil
}

func (m Model) View() string {
	canvasContent := m.BuildNextCanvasContent()

	return ui.Layout(
		m.width, m.height,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			ui.Canvas(CanvasWidth, CanvasHeight).Render(canvasContent),
			lipgloss.JoinVertical(
				lipgloss.Center,
				ui.Button().Render(m.msg),
				ui.Button().Render("something else"),
			),
		),
	)
}

// Update action button message by game state.
func (m Model) getActionButtonLabel() string {
	switch m.game.State {
	case game.Running:
		return "[P]ause"
	case game.GameOver:
		return "[R]estart"
	case game.Paused:
		return "[R]estart"
	case game.Win:
		return "[R]estart"
	default:
		log.Panic("Unknown game state")
	}

	return ""
}

// TODO: should this be a component?
func (m Model) BuildNextCanvasContent() string {
	strCanvas := strings.Builder{}

	width := m.game.Canvas.Width
	height := m.game.Canvas.Height

	snake := m.game.Snake
	apple := m.game.Apple

	for Y := 0; Y < width; Y++ {
		for X := 0; X < height; X++ {
			if snake.Contains(game.Coord{X: X, Y: Y}) {
				strCanvas.WriteString(ui.Snake(SnakeChar))
			} else if apple.X == X && apple.Y == Y {
				strCanvas.WriteString(ui.Apple(AppleChar))
			} else {
				strCanvas.WriteString(ui.NeutralChar(NeutralChar))
			}
		}
	}

	return strCanvas.String()
}
