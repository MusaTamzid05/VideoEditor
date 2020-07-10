package main

import (
	"video_editor/data"
	"video_editor/editor"
)

func main() {
	editor := editor.NewEditor("/home/musa/Downloads/lifeboy.mp4")
	editor.ExtractClip("test.mp4", *data.NewTime(0, 0, 5, 10))
}
