package model

import (
	"gosnake/internal/game"
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) HandleKeyPressed(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.Up):
		if m.game.Snake.Dir.IsOpposite(game.Up) {
			return *m, nil
		}
		m.nextSnakeMove = game.Up
	case key.Matches(msg, m.keys.Right):
		if m.game.Snake.Dir.IsOpposite(game.Right) {
			return *m, nil
		}
		m.nextSnakeMove = game.Right
	case key.Matches(msg, m.keys.Down):
		if m.game.Snake.Dir.IsOpposite(game.Down) {
			return *m, nil
		}
		m.nextSnakeMove = game.Down
	case key.Matches(msg, m.keys.Left):
		if m.game.Snake.Dir.IsOpposite(game.Left) {
			return *m, nil
		}
		m.nextSnakeMove = game.Left
	// GAME ACTIONS
	case key.Matches(msg, m.keys.Quit):
		return *m, tea.Quit
	case key.Matches(msg, m.keys.Pause):
		if m.game.State == game.Running {
			m.game.State = game.Paused
			m.msg = m.getActionButtonLabel()
		}
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
			return RestartModel(m.terminalWidth, m.terminalHeight), m.tick(m.game.Snake.Speed)
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
