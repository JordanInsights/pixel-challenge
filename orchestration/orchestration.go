package orchestration

import (
	"fmt"
	"os"
	"pixel-challenge/analysis"
	"pixel-challenge/images"
	"pixel-challenge/input"
)

func InitialiseTool() {
	imageFilepath, directoryFilepath := input.GetImageAndDirectoryFilepaths()

	fsImagesToBeAnalysed, _ := images.GetImagesFromFs(os.DirFS(directoryFilepath))
	fsReferenceImage, _ := images.GetSingleImage(fsImagesToBeAnalysed, imageFilepath)

	runAnalysis(fsImagesToBeAnalysed, fsReferenceImage)
}

func runAnalysis(imagesToBeAnalysed []images.Image, referenceImage images.Image) {
	for _, img := range imagesToBeAnalysed {
		similarity := analysis.CompareImages(referenceImage.Bytes, img.Bytes)
		str := "\n" + img.Name + ": "
		fmt.Printf("%q %+v", str, similarity)
	}
}
