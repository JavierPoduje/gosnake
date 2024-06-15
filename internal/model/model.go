package model

import (
	"fmt"
	"gosnake/internal/game"
	"gosnake/internal/logger"
	"gosnake/internal/ui"
	"log"
	"strconv"
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
const defaultSnakeSpeed = float64(3)

type Model struct {
	msg           string
	game          *game.Game
	logger        *logger.Logger
	nextSnakeMove game.Direction
	width         int
	height        int
}

func (m Model) tick(snakeSpeed float64) tea.Cmd {
	interval := time.Second / time.Duration(snakeSpeed)
	return tea.Tick(interval, func(t time.Time) tea.Msg {
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

func RestartModel(width, height int) Model {
	m := NewModel()
	m.width = width
	m.height = height
	return m
}

func (m Model) Init() tea.Cmd {
	return m.tick(defaultSnakeSpeed)
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

		// SNAKE MOVEMENTS
		case "up", "k":
			if m.game.Snake.IsOppositeDir(game.Up) {
				return m, nil
			}
			m.nextSnakeMove = game.Up
		case "right", "l":
			if m.game.Snake.IsOppositeDir(game.Right) {
				return m, nil
			}
			m.nextSnakeMove = game.Right
		case "down", "j":
			if m.game.Snake.IsOppositeDir(game.Down) {
				return m, nil
			}
			m.nextSnakeMove = game.Down
		case "left", "h":
			if m.game.Snake.IsOppositeDir(game.Left) {
				return m, nil
			}
			m.nextSnakeMove = game.Left
		// GAME ACTIONS
		case "p", "P":
			m.logger.Log("Pausing the game")
			m.game.State = game.Paused
			m.msg = m.getActionButtonLabel()
			return m, nil
		case "r", "R":
			m.logger.Log("Restarting the game")
			switch m.game.State {
			case game.Paused:
				m.game.State = game.Running
				m.msg = m.getActionButtonLabel()
				return m, m.tick(m.game.Snake.Speed)
			case game.GameOver:
				m.game.State = game.Running
				m.msg = m.getActionButtonLabel()
				return RestartModel(m.width, m.height), m.tick(m.game.Snake.Speed)
			default:
				log.Panic("Unreachable")
			}
		}

	case TickMsg:
		// update the game state
		m.game.Tick(m.nextSnakeMove)
		// update the game state
		m.msg = m.getActionButtonLabel()

		m.logger.Log(fmt.Sprintf("Score: %s", m.game.Stats.ScoreAsString()))

		if m.game.State == game.Running {
			return m, m.tick(m.game.Snake.Speed)
		}

		return m, nil
	}

	return m, nil
}

func (m Model) View() string {
	canvasContent := m.BuildNextCanvasContent()

	stats := [][]string{
		{"Eaten apples:", strconv.Itoa(m.game.Stats.EatenApples)},
		{"Score:", m.game.Stats.RoundedScoreAsString()},
	}

	return ui.Layout(
		m.width, m.height,
		lipgloss.JoinHorizontal(
			lipgloss.Top,
			ui.CanvasStyles(CanvasWidth, CanvasHeight).Render(canvasContent),
			lipgloss.JoinVertical(
				lipgloss.Center,
				ui.ActionButtonStyles(m.game.State).Render(m.msg),
				ui.StatsCard(stats),
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
