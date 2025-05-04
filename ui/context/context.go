package context

import (
	"fmt"
	"os"

	"ntduncan.com/squatcher/ui/utils"
)

type ProgramContext struct {
	ActiveFile      string
	ActiveWindow    int
	EditModeOn      bool
	CurrentDir      string
	CurrentDirItems []utils.DirItem
	MaxWidth        int
	MaxHeight       int
}

func NewProgramContext() *ProgramContext {
	currentDirectory, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("Error getting CWD: %s", err))
	}

	dirList, err := utils.GetCWDItems()
	if err != nil {
		panic(err)
	}

	return &ProgramContext{
		ActiveFile:      "",
		ActiveWindow:    1,
		EditModeOn:      false,
		CurrentDir:      currentDirectory + "/",
		CurrentDirItems: dirList,
		MaxWidth:        0,
		MaxHeight:       0,
	}
}
