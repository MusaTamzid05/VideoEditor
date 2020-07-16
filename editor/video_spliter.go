package editor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"video_editor/data"
)

type TimeJson struct {
	StartTime string
	EndTime   string
}

type VideoSpliter struct {
}

func (v *VideoSpliter) LoadJson(filepath string) ([]TimeJson, error) {

	timeJson := []TimeJson{}
	data, err := ioutil.ReadFile(filepath)

	if err != nil {
		return timeJson, err
	}

	err = json.Unmarshal(data, &timeJson)

	if err != nil {
		return timeJson, err
	}

	return timeJson, nil
}

func (v *VideoSpliter) convertToTime(timeJson TimeJson) (data.Time, error) {

	var err error
	var temp int64
	data := data.Time{}

	timeData := strings.Split(timeJson.StartTime, ":")

	for i := 0; i < 3; i++ {

		temp, err = strconv.ParseInt(timeData[i], 10, 64)

		if err != nil {
			return data, err
		}

		if i == 0 {
			data.Hour = int(temp)
		} else if i == 1 {

			data.Minute = int(temp)
		} else {
			data.Second = int(temp)
		}
	}

	return data, err

}

func (v *VideoSpliter) loadTime(timeJsonPath string) ([]data.Time, error) {

	timeJson, err := v.LoadJson(timeJsonPath)
	splitTimes := []data.Time{}

	if err != nil {
		return splitTimes, err
	}

	for _, data := range timeJson {

		timeData, err := v.convertToTime(data)

		if err != nil {
			return splitTimes, err
		}

		currentDiff, err := v.GetDiffirence(data)

		if err != nil {
			log.Println(err)
		}

		timeData.Duration = int(currentDiff)
		splitTimes = append(splitTimes, timeData)

	}

	return splitTimes, nil
}

func (v *VideoSpliter) Process(videoPath, timeJsonPath, dstPath string) error {

	splitTimes, err := v.loadTime(timeJsonPath)

	if err != nil {
		return err
	}

	outputPath := "temp"
	videoEditor := NewEditor(videoPath)

	fmt.Println("Extracting clips")
	splitVideoPaths, err := videoEditor.ExtractClips(outputPath, splitTimes)

	if err != nil {
		fmt.Println("Error extracing clips")
		return err
	}

	fmt.Println(splitVideoPaths)
	err = videoEditor.Concatenate(dstPath, splitVideoPaths)

	if err != nil {
		return err
	}
	return nil
}

func (v *VideoSpliter) GetDiffirence(data TimeJson) (int64, error) {

	startTime := strings.Split(data.StartTime, ":")
	endTime := strings.Split(data.EndTime, ":")

	var diff int64

	multipier := []int64{3600, 60, 1}
	currentIndex := 2

	var err error
	var startTimeInt int64
	var endTimeInt int64

	for currentIndex >= 0 {

		startTimeInt, err = strconv.ParseInt(startTime[currentIndex], 10, 64)

		if err != nil {
			return diff, nil
		}

		endTimeInt, err = strconv.ParseInt(endTime[currentIndex], 10, 64)

		if err != nil {
			return diff, nil
		}

		diff = diff + (endTimeInt-startTimeInt)*multipier[currentIndex]

		currentIndex = currentIndex - 1
	}

	return diff, nil

}
