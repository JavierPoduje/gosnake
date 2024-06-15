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

type TickMsg time.Time

const defaultSnakeDir = game.Right
const defaultSnakeSpeed = float64(3)

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

type keyMap struct {
	Down    key.Binding
	Help    key.Binding
	Left    key.Binding
	Pause   key.Binding
	Quit    key.Binding
	Restart key.Binding
	Right   key.Binding
	Up      key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right}, // first column
		{k.Quit},                        // second column
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
	),
	Pause: key.NewBinding(
		key.WithKeys("p"),
		key.WithHelp("p", "pause the game"),
	),
	Restart: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "restart the game"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
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
		m.logger.Log(fmt.Sprintf("width: %d, height: %d", msg.Width, msg.Height))
		m.width, m.height = msg.Width, msg.Height

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		// SNAKE MOVEMENTS
		case key.Matches(msg, m.keys.Up):
			if m.game.Snake.IsOppositeDir(game.Up) {
				return m, nil
			}
			m.nextSnakeMove = game.Up
		case key.Matches(msg, m.keys.Right):
			if m.game.Snake.IsOppositeDir(game.Right) {
				return m, nil
			}
			m.nextSnakeMove = game.Right
		case key.Matches(msg, m.keys.Down):
			if m.game.Snake.IsOppositeDir(game.Down) {
				return m, nil
			}
			m.nextSnakeMove = game.Down
		case key.Matches(msg, m.keys.Left):
			if m.game.Snake.IsOppositeDir(game.Left) {
				return m, nil
			}
			m.nextSnakeMove = game.Left
		// GAME ACTIONS
		case key.Matches(msg, m.keys.Pause):
			m.logger.Log("Pausing the game")
			m.game.State = game.Paused
			m.msg = m.getActionButtonLabel()
			return m, nil
		case key.Matches(msg, m.keys.Restart):
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
		// Help Action
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
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

	helpView := m.help.View(m.keys)

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
		helpView,
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
