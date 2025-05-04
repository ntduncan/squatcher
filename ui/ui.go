package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	asciiwindow "ntduncan.com/squatcher/ui/components/ascii-window"
	filemanager "ntduncan.com/squatcher/ui/components/file-manager"
	fileviewer "ntduncan.com/squatcher/ui/components/file-viewer"
	"ntduncan.com/squatcher/ui/context"
	"ntduncan.com/squatcher/ui/utils"
)

type Model struct {
	filemanager tea.Model
	asciiwindow tea.Model
	fileviewer  tea.Model
	ctx         *context.ProgramContext
}

func NewModel() tea.Model {
	ctx := context.NewProgramContext()

	m := Model{
		ctx:         ctx,
		filemanager: filemanager.NewModel(ctx),
		asciiwindow: asciiwindow.NewModel(ctx),
		fileviewer:  fileviewer.NewModel(ctx),
	}

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var activeWin tea.Model

	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {
		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}

		switch m.ctx.ActiveWindow {
		case utils.FileManager:
			activeWin, cmd = m.filemanager.Update(msg)
			m.filemanager = activeWin
		case utils.FileViewer:
			activeWin, cmd = m.fileviewer.Update(msg)
			m.fileviewer = activeWin
		}

	case tea.WindowSizeMsg:
		m.ctx.MaxWidth = msg.Width - 2
		m.ctx.MaxHeight = msg.Height
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.

	return m, cmd
}

func (m Model) View() string {
	s := strings.Builder{}

	s.WriteString(m.renderCWD())

	return lipgloss.NewStyle().Border(lipgloss.NormalBorder()).Render(
		lipgloss.JoinVertical(
			0,
			m.asciiwindow.View(),
			lipgloss.NewStyle().Height(m.ctx.MaxHeight-12).Width(m.ctx.MaxWidth-2).BorderTop(true).BorderBottom(true).BorderStyle(lipgloss.NormalBorder()).Render(
				lipgloss.JoinHorizontal(
					lipgloss.Left,
					m.filemanager.View(),
					m.fileviewer.View(),
				),
			),
			s.String(),
		),
	)
}

func (m *Model) renderCWD() string {
	return lipgloss.NewStyle().Align(lipgloss.Center).Foreground(lipgloss.Color("#85DEAD")).Render(m.ctx.CurrentDir)
}
