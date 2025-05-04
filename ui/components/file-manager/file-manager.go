package filemanager

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"ntduncan.com/squatcher/ui/context"
	"ntduncan.com/squatcher/ui/utils"
)

type DirMovement int

const (
	Up = iota
	Down
)

type Model struct {
	window   utils.Window
	ctx      *context.ProgramContext
	viewport viewport.Model
	cursor   int
}

func NewModel(ctx *context.ProgramContext) Model {
	return Model{
		window: utils.FileManager,
		ctx:    ctx,
		viewport: viewport.Model{
			Width:  20,
			Height: ctx.MaxHeight,
		},
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.cursor == len(m.ctx.CurrentDirItems)-1 {
				m.cursor = 0

			} else if m.cursor < len(m.ctx.CurrentDirItems)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor == 0 {
				m.cursor = len(m.ctx.CurrentDirItems) - 1

			} else if m.cursor > 0 {
				m.cursor--
			}
		case "enter", " ":
			cursorItem := m.ctx.CurrentDirItems[m.cursor]
			switch true {
			//Selected Dir
			case cursorItem.Value == "../":
				err := m.StepDirectory(Up)
				if err != nil {
					panic(fmt.Errorf("Error reading new dir: %s", err))
				}
			case cursorItem.IsDir:
				err := m.StepDirectory(Down)
				if err != nil {
					panic(fmt.Errorf("Error reading new dir: %s", err))
				}
			//Deselecting
			case m.ctx.ActiveFile == cursorItem.Value:
				m.ctx.ActiveFile = ""
				m.ctx.ActiveWindow = utils.FileViewer
			//Select File
			default:
				m.ctx.ActiveFile = m.ctx.CurrentDirItems[m.cursor].Value
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.viewport.Height = msg.Height - 11
	}

	return m, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) View() string {
	s := strings.Builder{}

	for i, item := range m.ctx.CurrentDirItems {
		if i == m.cursor {
			s.WriteString(renderCursorListItem(item.Value))
			continue
		}

		if item.Value == m.ctx.ActiveFile {
			s.WriteString(renderActiveListItemStyles(item.Value))
			continue
		}

		s.WriteString(renderListItemStyles(item.Value))
	}

	return lipgloss.NewStyle().Height(m.ctx.MaxHeight - 12).Width(20).BorderRight(true).BorderStyle(lipgloss.NormalBorder()).Foreground(lipgloss.Color("#FFF")).Render(s.String())
}

func (m Model) StepDirectory(direction DirMovement) error {
	newDirItem := m.ctx.CurrentDirItems[m.cursor]
	if !newDirItem.IsDir {
		return fmt.Errorf("Oops %s isn't a dir! isDir value: %t\n", newDirItem.Value, newDirItem.IsDir)
	}

	var err error
	if direction == Up {
		err = os.Chdir("../")
	} else {
		err = os.Chdir(m.ctx.CurrentDir + newDirItem.Value)
	}

	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("Error getting working dir name: %s\n", err)
	}

	m.ctx.CurrentDirItems, err = utils.GetCWDItems()
	if err != nil {
		return err
	}

	m.ctx.CurrentDir = cwd + "/"
	m.ctx.ActiveFile = ""
	m.cursor = 1

	return nil
}

func renderCursorListItem(s string) string {
	lightGray := lipgloss.Color("#FFF")
	purple := lipgloss.Color("99")
	return lipgloss.NewStyle().Align(lipgloss.Center).MarginTop(1).Background(purple).Foreground(lightGray).Render(s)
}

func renderListItemStyles(s string) string {
	gray := lipgloss.Color("245")
	return lipgloss.NewStyle().Align(lipgloss.Center).MarginTop(1).Foreground(gray).Render(s)
}

func renderActiveListItemStyles(s string) string {
	purple := lipgloss.Color("99")
	return lipgloss.NewStyle().Align(lipgloss.Center).MarginTop(1).Foreground(purple).Render(s)
}
