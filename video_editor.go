package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"video_editor/editor"
)

func SplitVideoFrom(videoPath, timeJsonPath, dstPath string) {
	videoSpliter := editor.VideoSpliter{}
	err := videoSpliter.Process(videoPath, timeJsonPath, dstPath)

	if err != nil {
		fmt.Println(err)
	}
}

func Concatenate(videoPaths, dstPath string) {

	paths := strings.Split(videoPaths, ",")

	if len(paths) == 0 {
		log.Println("Video path does only have one video path,maybe you forget to give a comma")
	}

	editor_ := editor.Editor{}
	err := editor_.Concatenate(dstPath, paths)

	if err != nil {
		log.Fatalln(err)
	}
}

func Usage() {
	log.Fatalln("Usage: -s -v path -j json_path -d dst.mp4 | -c -v path -a audio_file_paths -d dst.mp4")
}

func main() {

	splitPtr := flag.Bool("s", false, "Set it to true if what to convert")
	concatPtr := flag.Bool("c", false, "Set it to true if want to concate")
	videoSrcPath := flag.String("v", "", "Video src path")
	jsonPathPtr := flag.String("j", "", "Json path")
	dstPathPtr := flag.String("d", "", "Dst path")
	audioFilePaths := flag.String("a", "", "audiopath")

	flag.Parse()

	if *splitPtr == true {

		if *videoSrcPath == "" {
			Usage()
		}

		if *jsonPathPtr == "" || *dstPathPtr == "" {
			Usage()
		}
		SplitVideoFrom(*videoSrcPath, *jsonPathPtr, *dstPathPtr)

	} else if *concatPtr == true {

		if *audioFilePaths == "" || *dstPathPtr == "" {
			Usage()
		}

		fmt.Println("Concateneting")
		Concatenate(*audioFilePaths, *dstPathPtr)
	} else {
		Usage()
	}

}
