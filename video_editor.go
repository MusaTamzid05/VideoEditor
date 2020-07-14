package main

import (
	"fmt"
	"video_editor/data"
	"video_editor/editor"
)

func main() {

	times := []data.Time{}
	times = append(times, *data.NewTime(0, 0, 5, 10))
	times = append(times, *data.NewTime(0, 0, 15, 5))

	editor := editor.NewEditor("/home/musa/Downloads/lifeboy.mp4")
	err := editor.Concatenate("result.mp4", []string{"test-0.mp4", "test-1.mp4"})

	if err != nil {
		fmt.Println("Error :", err)
	}

}
