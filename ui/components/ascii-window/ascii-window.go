package asciiwindow

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"ntduncan.com/squatcher/ui/context"
	"ntduncan.com/squatcher/ui/utils"
)

type Model struct {
	window   utils.Window
	ctx      *context.ProgramContext
	viewport viewport.Model
}

func NewModel(ctx *context.ProgramContext) Model {
	return Model{
		window: utils.AciiWin,
		ctx:    ctx,
		viewport: viewport.Model{
			Width:  ctx.MaxWidth,
			Height: 0,
		},
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	s := strings.Builder{}
	purple := lipgloss.Color("99")
	s.WriteString(" _____  _____ _   _  ___ _____ _____  _   _  ___________ \n")
	s.WriteString("/  ___||  _  | | | |/ _ \\_   _/  __ \\| | | ||  ___| ___ \\\n")
	s.WriteString("\\ `--. | | | | | | / /_\\ \\| | | /  \\/| |_| || |__ | |_/ /\n")
	s.WriteString(" `--. \\| | | | | | |  _  || | | |    |  _  ||  __||    / \n")
	s.WriteString("/\\__/ /\\ \\/' / |_| | | | || | | \\__/\\| | | || |___| |\\ \\ \n")
	s.WriteString("\\____/  \\_/\\_\\___/\\\\_| |_/\\_/  \\____/\\_| |_/\\____/\\_| \\_|\n")

	return lipgloss.NewStyle().Align(lipgloss.Center).Width(m.ctx.MaxWidth).Foreground(purple).Render(s.String())
}
