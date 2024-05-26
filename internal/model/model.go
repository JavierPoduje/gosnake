package model

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
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
	return m.msg
}
