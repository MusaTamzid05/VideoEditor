package main

import (
	"flag"
	"fmt"
	"log"
	"video_editor/editor"
)

func SplitVideoFrom(videoPath, timeJsonPath, dstPath string) {
	videoSpliter := editor.VideoSpliter{}
	err := videoSpliter.Process(videoPath, timeJsonPath, dstPath)

	if err != nil {
		fmt.Println(err)
	}
}

func Usage() {
	log.Fatalln("Usage: -s -v path -j json_path -d dst.mp4")
}

func main() {

	splitPtr := flag.Bool("s", false, "Set it to false")
	videoSrcPath := flag.String("v", "", "Video src path")
	jsonPathPtr := flag.String("j", "", "Json path")
	dstPathPtr := flag.String("d", "", "Dst path")

	flag.Parse()

	if *videoSrcPath == "" {
		Usage()
	}

	if *splitPtr == true {
		if *jsonPathPtr == "" || *dstPathPtr == "" {
			Usage()
		}
		SplitVideoFrom(*videoSrcPath, *jsonPathPtr, *dstPathPtr)
	}

	//SplitVideoFrom("/home/musa/Downloads/lifeboy.mp4", "data.json", "final.mp4")

}
