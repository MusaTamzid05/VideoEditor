package editor

import "fmt"

type Editor struct {
}

func (e *Editor) Process() {
	fmt.Println("Basic editor struct")
}

func NewEditor() *Editor {
	return &Editor{}
}
