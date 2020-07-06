package main

import (
	"fmt"
	"video_editor/data"
)

func main() {
	t1 := data.NewTime(6, 5, 7)
	fmt.Println(t1.String())

	t2 := data.NewTime(11, 12, 13)
	fmt.Println(t2.String())
}
