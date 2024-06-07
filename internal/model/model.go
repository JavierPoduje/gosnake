package model

import (
	"gosnake/internal/ui"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	doc := strings.Builder{}

	//buttons := lipgloss.JoinHorizontal(lipgloss.Top,
	//    ui.Text(m.msg),
	//    ui.Button("Another"),
	//)

	//doc.WriteString(buttons)
	for i := 0; i < 100; i++ {
		if i == 13 {
			doc.WriteString("\uee98")
		} else {
			doc.WriteString(" ")
		}
	}
	//doc.WriteString(ui.Grid())

	//return doc.String()

	filledCanvas := lipgloss.Place(
		80, 40,
		lipgloss.Center, lipgloss.Center,
		ui.Canvas().Render(doc.String()),
		lipgloss.WithWhitespaceChars("[]"),
		lipgloss.WithWhitespaceForeground(lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}),
	)

	return filledCanvas
}
