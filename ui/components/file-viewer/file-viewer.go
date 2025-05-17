package fileviewer

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"

	//"github.com/charmbracelet/lipgloss"
	"ntduncan.com/squatcher/ui/context"
	"ntduncan.com/squatcher/ui/utils"
)

type Model struct {
	window   utils.Window
	Viewport viewport.Model
	ctx      *context.ProgramContext
	content  string
}

func NewModel(ctx *context.ProgramContext) Model {

	m := Model{
		window: utils.FileViewer,
		ctx:    ctx,
		Viewport: viewport.Model{
			Width:  ctx.MaxWidth - 18,
			Height: ctx.MaxHeight - 22,
		},
	}

	return m
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Viewport.ScrollDown(5)
			m.Viewport, cmd = m.Viewport.Update(msg)
			return m, cmd

		case "k", "up":
			m.Viewport.ScrollUp(5)
			m.Viewport, cmd = m.Viewport.Update(msg)
			return m, cmd
		case "esc", "escape", "h", "left":
			m.ctx.ActiveWindow = utils.FileManager
			//m.ctx.ActiveFile = ""
		}

	case tea.WindowSizeMsg:
		m.Viewport = viewport.New(msg.Width-18, msg.Height-22)
		m.Viewport.YPosition = 22

	}

	content, err := m.getActiveFileContent()
	if err != nil {
		panic("An Error happened")
	}

	m.Viewport.SetContent(content)
	m.Viewport, cmd = m.Viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
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

		r, _ := glamour.NewTermRenderer(
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(m.ctx.MaxWidth-22),
		)

		out, err := r.Render(content)
		if err != nil {
			panic(fmt.Errorf("Error render file: %s", err))
		}

		borderColor := lipgloss.Color("99")
		if m.ctx.ActiveWindow != utils.FileViewer {
			borderColor = lipgloss.Color("#FFF")
		}

		return lipgloss.NewStyle().
			Width(m.Viewport.Width).
			Height(m.ctx.MaxHeight - 10).
			MaxHeight(m.ctx.MaxHeight - 10).
			BorderStyle(lipgloss.DoubleBorder()).
			BorderForeground(borderColor).
			Render(out)
	}

	return ""
}

func (m Model) getActiveFileContent() (string, error) {
	filedata, err := os.ReadFile(m.ctx.CurrentDir + m.ctx.ActiveFile)
	if err != nil {
		return "", fmt.Errorf("Could not read file: %s", err)
	}

	return string(filedata), nil

}
