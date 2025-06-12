package model

import (
	"gosnake/internal/db"
	"gosnake/internal/game"
	"gosnake/internal/logger"
	"gosnake/internal/ui"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/help"
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
	game           *game.Game
	logger         *logger.Logger
	terminalHeight int
	terminalWidth  int

	db            db.DB
	msg           string
	nextSnakeMove game.Direction
	scores        []int

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
	db := db.NewDB()
	return Model{
		db:             db,
		game:           game.NewGame(CanvasWidth, CanvasHeight),
		help:           help.New(),
		keys:           keys,
		logger:         logger.NewLogger("debug.log"),
		msg:            "Initializing...",
		nextSnakeMove:  defaultSnakeDir,
		scores:         db.GetScores(),
		terminalHeight: DefaultTerminalHeight,
		terminalWidth:  DefaultTerminalWidth,
	}
}

func RestartModel(width, height int) Model {
	m := NewModel()
	m.terminalWidth = width
	m.terminalHeight = height
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

func (model Model) View() string {
	canvasContent := model.BuildNextCanvasContent()
	stats := model.BuildNextStatsContent()

	// components
	canvas := ui.Canvas(CanvasWidth, CanvasHeight, model.game.State, canvasContent)
	footer := ui.HelpContainer(model.help.View(model.keys))
	historicScoresCard := ui.HistoricScoresCard(model.scores)
	statsCard := ui.StatsCard(stats)

	// build the content
	infoCards := lipgloss.JoinVertical(lipgloss.Center, statsCard, historicScoresCard)
	contentSection := lipgloss.JoinHorizontal(lipgloss.Top, canvas, infoCards)
	content := lipgloss.JoinVertical(lipgloss.Right, contentSection, footer)

	return ui.Layout(model.terminalWidth, model.terminalHeight, content)
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

func (m *Model) HandleTick() (Model, tea.Cmd) {
	if m.game.State == game.Paused {
		return *m, nil
	}

	m.game.Tick(m.nextSnakeMove)
	m.msg = m.getActionButtonLabel()

	if m.game.State == game.Running {
		return *m, m.tick(m.game.Snake.Speed)
	}

	if m.game.State == game.GameOver && m.game.Stats.ScoreAsInt() > 0 {
		m.db.SaveScore(m.game.Stats.ScoreAsInt())
		m.scores = m.db.GetScores()
		return *m, nil
	}

	return *m, m.tick(m.game.Snake.Speed)
}

func (model Model) BuildNextStatsContent() [][]string {
	return [][]string{
		{"Eaten apples:", strconv.Itoa(model.game.Stats.EatenApples)},
		{"Score:", model.game.Stats.RoundedScoreAsString()},
	}
}

func (m Model) BuildNextCanvasContent() string {
	strCanvas := strings.Builder{}

	width := m.game.Canvas.Width
	height := m.game.Canvas.Height

	snake := m.game.Snake
	apple := m.game.Apple

	for Y := range width {
		for X := range height {
			coordinate := game.Coord{X: X, Y: Y}

			if snake.Contains(coordinate) {
				if snake.IsHead(coordinate) {
					strCanvas.WriteString(ui.SnakeHead(SnakeChar))
				} else {
					strCanvas.WriteString(ui.SnakeBody(SnakeChar))
				}
			} else if apple.Equals(coordinate) {
				strCanvas.WriteString(ui.Apple(AppleChar))
			} else {
				strCanvas.WriteString(ui.NeutralChar(NeutralChar))
			}
		}
	}

	return strCanvas.String()
}
