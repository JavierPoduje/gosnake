package model

import (
	"gosnake/internal/game"
	"gosnake/internal/logger"
	"gosnake/internal/ui"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
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

const defaultSnakeDir = game.Right
const defaultSnakeSpeed = float64(3)

type TickMsg time.Time

type Model struct {
	game   *game.Game
	logger *logger.Logger
	height int
	width  int

	msg           string
	nextSnakeMove game.Direction

	help help.Model
	keys keyMap
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
		help:          help.New(),
		keys:          keys,
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
		m.HandleWindowResize(msg)
	case tea.KeyMsg:
		return m.HandleKeyPressed(msg)
	case TickMsg:
		return m.HandleTick()
	}

	return m, nil
}

func (m Model) View() string {
	canvasContent := m.BuildNextCanvasContent()

	stats := [][]string{
		{"Eaten apples:", strconv.Itoa(m.game.Stats.EatenApples)},
		{"Score:", m.game.Stats.RoundedScoreAsString()},
	}

	keysAsString := m.help.View(m.keys)

	return ui.Layout(
		m.width, m.height,
		lipgloss.JoinVertical(
			lipgloss.Right,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				ui.CanvasStyles(CanvasWidth, CanvasHeight).Render(canvasContent),
				lipgloss.JoinVertical(
					lipgloss.Center,
					ui.ActionButtonStyles(m.game.State).Render(m.msg),
					ui.StatsCard(stats),
				),
			),
			ui.HelpContainer(keysAsString),
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

func (m *Model) HandleWindowResize(msg tea.WindowSizeMsg) {
	m.width, m.height = msg.Width, msg.Height
}

func (m *Model) HandleKeyPressed(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.Up):
		if m.game.Snake.IsOppositeDir(game.Up) {
			return *m, nil
		}
		m.nextSnakeMove = game.Up
	case key.Matches(msg, m.keys.Right):
		if m.game.Snake.IsOppositeDir(game.Right) {
			return *m, nil
		}
		m.nextSnakeMove = game.Right
	case key.Matches(msg, m.keys.Down):
		if m.game.Snake.IsOppositeDir(game.Down) {
			return *m, nil
		}
		m.nextSnakeMove = game.Down
	case key.Matches(msg, m.keys.Left):
		if m.game.Snake.IsOppositeDir(game.Left) {
			return *m, nil
		}
		m.nextSnakeMove = game.Left
	// GAME ACTIONS
	case key.Matches(msg, m.keys.Quit):
		return *m, tea.Quit
	case key.Matches(msg, m.keys.Pause):
		m.game.State = game.Paused
		m.msg = m.getActionButtonLabel()
		return *m, nil
	case key.Matches(msg, m.keys.Restart):
		switch m.game.State {
		case game.Paused:
			m.game.State = game.Running
			m.msg = m.getActionButtonLabel()
			return *m, m.tick(m.game.Snake.Speed)
		case game.GameOver:
			m.game.State = game.Running
			m.msg = m.getActionButtonLabel()
			return RestartModel(m.width, m.height), m.tick(m.game.Snake.Speed)
		default:
			log.Panic("Unreachable")
		}
	// Help Action
	case key.Matches(msg, m.keys.Help):
		m.help.ShowAll = !m.help.ShowAll
	default:
		return *m, nil
	}

	return *m, nil
}

func (m *Model) HandleTick() (Model, tea.Cmd) {
	m.game.Tick(m.nextSnakeMove)
	m.msg = m.getActionButtonLabel()

	if m.game.State == game.Running {
		return *m, m.tick(m.game.Snake.Speed)
	}

	return *m, nil
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
