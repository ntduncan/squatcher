package utils

import (
	"fmt"
	"os"
)

func GetCWDItems() ([]DirItem, error) {
	var fileItems []DirItem
	var dirItems []DirItem

	cwd, err := os.Getwd()
	if err != nil {
		//@TODO: Add a logger eventually
		return nil, fmt.Errorf("Error getting CWD, %s", err)
	}

	enteries, err := os.ReadDir(cwd)
	if err != nil {
		return nil, fmt.Errorf("Error getting CWD Items, %s", err)
	}

	//Add Option to go up a Dir
	backDirItem := DirItem{
		Value: "../",
		IsDir: true,
	}
	dirItems = append(dirItems, backDirItem)

	for _, entry := range enteries {
		i := DirItem{
			Value: entry.Name(),
			IsDir: entry.IsDir(),
		}

		if entry.IsDir() {
			i.Value = i.Value + "/"
			dirItems = append(dirItems, i)
		} else {
			fileItems = append(fileItems, i)
		}
	}

	return append(dirItems, fileItems...), nil
}
