package orchestration

import (
	"fmt"
	"os"
	"pixel-challenge/analysis"
	"pixel-challenge/images"
)

type imageFilepaths struct {
	comparisonImage string
	imageDirectory  string
}

func getFilepathsFromCommandLineArguments() imageFilepaths {
	fp := imageFilepaths{os.Args[1], os.Args[2]}
	fmt.Println(fp)
	return fp
}

func InitialiseTool() {
	// imageFilepath, directoryFilepath := input.GetImageAndDirectoryFilepaths()
	// fsImagesToBeAnalysed, _ := images.GetImagesFromFs(os.DirFS(directoryFilepath))
	// fsComparisonImage, _ := images.GetSingleImage(fsImagesToBeAnalysed, imageFilepath)

	filepaths := getFilepathsFromCommandLineArguments()
	fsImagesToBeAnalysed, _ := images.GetImagesFromFs(os.DirFS(filepaths.imageDirectory))
	fsComparisonImage, _ := images.GetSingleImage(fsImagesToBeAnalysed, filepaths.comparisonImage)

	go monitorAnalyses(fsComparisonImage)
	addAnalyses(fsImagesToBeAnalysed)
}

type analysisOperation struct {
	img               images.Image
	similarityChannel chan float64
	errorChannel      chan error
}

var analyses chan analysisOperation = make(chan analysisOperation)
var done chan struct{} = make(chan struct{})

func addAnalysisOperation(img images.Image) (float64, error) {
	similarityChannel := make(chan float64)
	errorChannel := make(chan error)

	op := analysisOperation{
		img:               img,
		similarityChannel: similarityChannel,
		errorChannel:      errorChannel,
	}

	analyses <- op
	return <-similarityChannel, <-errorChannel
}

func monitorAnalyses(referenceImage images.Image) {
	for op := range analyses {
		go func(op analysisOperation) {
			op.similarityChannel <- analysis.CompareImages(referenceImage.Bytes, op.img.Bytes)
			op.errorChannel <- nil
		}(op)
	}

	close(done)
}

func addAnalyses(imagesToBeAnalysed []images.Image) {
	fmt.Println("\nAnalysis: ")

	for _, img := range imagesToBeAnalysed {
		addAnalysisOperation(img)
	}

	stopAnalyses()
}

func stopAnalyses() {
	close(analyses)
	<-done
}
