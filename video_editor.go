package main

import (
	"fmt"
	"video_editor/editor"
)

func SplitVideoFrom(videoPath, timeJsonPath, dstPath string) {
	videoSpliter := editor.VideoSpliter{}
	err := videoSpliter.Process(videoPath, timeJsonPath, dstPath)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	SplitVideoFrom("/home/musa/Downloads/lifeboy.mp4", "data.json", "final.mp4")

}
