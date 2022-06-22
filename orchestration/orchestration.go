package orchestration

import (
	"fmt"
	"pixel-challenge/input"
)

func InitialiseTool() {
	imageFilepath, directoryFilepath := input.GetImageAndDirectoryFilepaths()
	fmt.Println(imageFilepath, directoryFilepath)
}
