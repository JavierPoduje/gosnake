package model

import tea "github.com/charmbracelet/bubbletea"

func (m *Model) HandleWindowResize(msg tea.WindowSizeMsg) {
	m.terminalWidth, m.terminalHeight = msg.Width, msg.Height
}
