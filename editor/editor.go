package editor

import "fmt"

type Editor struct {
	path string
}

func (e *Editor) Process() {
	fmt.Println("Basic editor struct")
}

func NewEditor(path string) *Editor {
	return &Editor{path: path}
}
