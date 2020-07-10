package main

import (
	"video_editor/data"
	"video_editor/editor"
)

func main() {

	times := []data.Time{}
	times = append(times, *data.NewTime(0, 0, 5, 10))
	times = append(times, *data.NewTime(0, 0, 15, 5))

	editor := editor.NewEditor("/home/musa/Downloads/lifeboy.mp4")
	editor.ExtractClips("test", times)

}
