package orchestration

import (
	"os"
	"pixel-challenge/analysis"
	"pixel-challenge/images"
	"sort"
)

type ImageFilepaths struct {
	ComparisonImage string
	ImageDirectory  string
}

type similarityResult struct {
	imageName  string
	similarity float64
}

func GetFilepathsFromCommandLineArguments() ImageFilepaths {
	fp := ImageFilepaths{os.Args[1], os.Args[2]}
	return fp
}

func RunAnalysis(filepaths ImageFilepaths) {
	// imageFilepath, directoryFilepath := input.GetImageAndDirectoryFilepaths()
	// fsImagesToBeAnalysed, _ := images.GetImagesFromFs(os.DirFS(directoryFilepath))
	// fsComparisonImage, _ := images.GetSingleImage(fsImagesToBeAnalysed, imageFilepath)

	fsImagesToBeAnalysed, _ := images.GetImagesFromFs(os.DirFS(filepaths.ImageDirectory))
	fsComparisonImage, _ := images.GetSingleImage(fsImagesToBeAnalysed, filepaths.ComparisonImage)

	go monitorAnalyses(fsComparisonImage)
	go handleSimilarities(fsComparisonImage)
	addAnalyses(fsImagesToBeAnalysed)
}

type analysisOperation struct {
	img               images.Image
	similarityChannel chan float64
	errorChannel      chan error
}

var analyses chan analysisOperation = make(chan analysisOperation)
var analysesDone chan struct{} = make(chan struct{})

var similarities chan similarityResult = make(chan similarityResult)
var similaritiesDone chan struct{} = make(chan struct{})

var TopThreeSimilarities []similarityResult = make([]similarityResult, 3)

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
			similarity := analysis.CompareImages(referenceImage.Bytes, op.img.Bytes)

			similarities <- similarityResult{op.img.Name, similarity}
			op.similarityChannel <- similarity
			op.errorChannel <- nil
		}(op)
	}

	close(analysesDone)
}

func handleSimilarities(referenceImage images.Image) {
	for s := range similarities {
		if referenceImage.Name != s.imageName && s.similarity >= TopThreeSimilarities[2].similarity {
			TopThreeSimilarities[2] = s
			sort.SliceStable(TopThreeSimilarities, func(i, j int) bool {
				return TopThreeSimilarities[i].similarity > TopThreeSimilarities[j].similarity
			})
		}
	}

	close(similaritiesDone)
}

func addAnalyses(imagesToBeAnalysed []images.Image) {
	for _, img := range imagesToBeAnalysed {
		addAnalysisOperation(img)
	}

	stopAnalyses()
}

func stopAnalyses() {
	close(analyses)
	close(similarities)
}
