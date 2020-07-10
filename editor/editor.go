package editor

import (
	"fmt"
	"video_editor/data"
	"video_editor/util"
)

type Editor struct {
	path string
}

func (e *Editor) ExtractClip(output string, startTime data.Time) error {

	commandStr := fmt.Sprintf("ffmpeg -i %s -ss %s -codec copy -t %d %s", e.path, startTime.String(), startTime.Duration, output)
	fmt.Println(commandStr)
	_, err := util.ExecuteCommand(commandStr)
	return err
}

func (e *Editor) ExtractClips(output string, times []data.Time) ([]string, error) {

	var err error
	dstPaths := []string{}

	for index, timeData := range times {
		dstPath := fmt.Sprintf("%s-%d.mp4", output, index)
		err = e.ExtractClip(dstPath, timeData)

		if err != nil {
			return dstPaths, err
		}
	}

	return dstPaths, nil
}

func (e *Editor) ConvertToMp4(dstPath string) error {

	commandStr := fmt.Sprintf("ffmpeg -i %s -c:v libx24 %s", e.path, dstPath)
	_, err := util.ExecuteCommand(commandStr)
	return err
}

func NewEditor(path string) *Editor {
	return &Editor{path: path}
}
