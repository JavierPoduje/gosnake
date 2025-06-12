package model

import (
	"gosnake/internal/game"
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) HandleKeyPressed(msg tea.KeyMsg) (Model, tea.Cmd) {
	switch {
	// Direction keys
	case key.Matches(msg, m.keys.Up):
		return m.handleKeyUp()
	case key.Matches(msg, m.keys.Right):
		return m.handleKeyRight()
	case key.Matches(msg, m.keys.Down):
		return m.handleKeyDown()
	case key.Matches(msg, m.keys.Left):
		return m.handleKeyLeft()

	// Game actions
	case key.Matches(msg, m.keys.Quit):
		return m.handleQuit()
	case key.Matches(msg, m.keys.Pause):
		return m.handlePause()
	case key.Matches(msg, m.keys.Restart):
		return m.handleRestart()

	// Help Action
	case key.Matches(msg, m.keys.Help):
		return m.handleHelp()
	default:
		return *m, nil
	}
}

func (m *Model) handleKeyUp() (Model, tea.Cmd) {
	if !m.game.Snake.Dir.IsOpposite(game.Up) {
		m.nextSnakeMove = game.Up
	}
	return *m, nil
}

func (m *Model) handleKeyRight() (Model, tea.Cmd) {
	if !m.game.Snake.Dir.IsOpposite(game.Right) {
		m.nextSnakeMove = game.Right
	}
	return *m, nil
}

func (m *Model) handleKeyDown() (Model, tea.Cmd) {
	if !m.game.Snake.Dir.IsOpposite(game.Down) {
		m.nextSnakeMove = game.Down
	}
	return *m, nil
}

func (m *Model) handleKeyLeft() (Model, tea.Cmd) {
	if !m.game.Snake.Dir.IsOpposite(game.Left) {
		m.nextSnakeMove = game.Left
	}
	return *m, nil
}

func (m *Model) handleQuit() (Model, tea.Cmd) {
	return *m, tea.Quit
}

func (m *Model) handlePause() (Model, tea.Cmd) {
	if m.game.State == game.Running {
		m.game.State = game.Paused
		m.msg = m.getActionButtonLabel()
	}
	return *m, nil
}

func (m *Model) handleRestart() (Model, tea.Cmd) {
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
	return *m, nil
}

func (m *Model) handleHelp() (Model, tea.Cmd) {
	m.help.ShowAll = !m.help.ShowAll
	return *m, nil
}
