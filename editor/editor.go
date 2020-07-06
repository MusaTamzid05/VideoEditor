package editor

import (
	"fmt"
	"video_editor/data"
	"video_editor/util"
)

type Editor struct {
	path string
}

func (e *Editor) ExtractClip(output string, startTime data.Time, duration int) error {

	commandStr := fmt.Sprintf("ffmpeg -i %s -ss %s -codec copy -t %d %s", e.path, startTime.String(), duration, output)
	fmt.Println(commandStr)
	_, err := util.ExecuteCommand(commandStr)
	return err
}

func NewEditor(path string) *Editor {
	return &Editor{path: path}
}
