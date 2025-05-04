package fileviewer

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"ntduncan.com/squatcher/ui/context"
	"ntduncan.com/squatcher/ui/utils"
)

type Model struct {
	window   utils.Window
	viewport viewport.Model
	ctx      *context.ProgramContext
}

func NewModel(ctx *context.ProgramContext) Model {
	return Model{
		window: utils.FileViewer,
		ctx:    ctx,
		viewport: viewport.Model{
			Width:  ctx.MaxWidth - 20,
			Height: ctx.MaxHeight,
		},
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.viewport.PageDown()
		case "k", "up":
			m.viewport.PageUp()
		case "esc", "escape":
			m.ctx.ActiveWindow = utils.FileManager
		}

	case tea.WindowSizeMsg:
		m.viewport.Width = msg.Width - 20
	}

	return m, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	return lipgloss.NewStyle().AlignVertical(lipgloss.Center).AlignHorizontal(lipgloss.Center).Render(m.ctx.ActiveFile)
}
