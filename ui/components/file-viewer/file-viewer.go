package fileviewer

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	//"github.com/charmbracelet/lipgloss"
	"ntduncan.com/squatcher/ui/context"
	"ntduncan.com/squatcher/ui/utils"
)

type Model struct {
	window   utils.Window
	viewport viewport.Model
	ctx      *context.ProgramContext
	content  string
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
	if m.ctx.ActiveFile != "" {

		content, err := m.getActiveFileContent()
		if err != nil {
			panic(fmt.Errorf("Error read file: %s", err))
		}

		out, err := glamour.Render(content, "dark")
		if err != nil {
			panic(fmt.Errorf("Error render file: %s", err))
		}

		return out
	} else {
		return ""
	}
}

func (m Model) getActiveFileContent() (string, error) {
	filedata, err := os.ReadFile(m.ctx.CurrentDir + m.ctx.ActiveFile)
	if err != nil {
		return "", fmt.Errorf("Could not read file: %s", err)
	}

	return string(filedata), nil

}
