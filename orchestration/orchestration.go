package orchestration

import (
	"fmt"
	"os"
	"pixel-challenge/images"
	"pixel-challenge/input"
)

func InitialiseTool() {
	imageFilepath, directoryFilepath := input.GetImageAndDirectoryFilepaths()
	fmt.Println(imageFilepath, directoryFilepath)

	fsImages, _ := images.GetImagesFromFs(os.DirFS("./test-images/Bronze"))
	fmt.Println(fsImages[0].Name, len(fsImages[0].Bytes))
}
