package utils

type DirItem struct {
	Value string
	IsDir bool
}

type Window int

const (
	AciiWin = iota
	FileManager
	FileViewer
)
