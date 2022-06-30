package orchestration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"pixel-challenge/analysis"
	"pixel-challenge/images"
	"sort"
	"strings"
	"time"
)

type ImageFilepaths struct {
	ComparisonImage string
	ImageDirectory  string
}

type similarityResult struct {
	ImageName  string
	Similarity float64
}

type jsonResult struct {
	ComparisonImage, DirectoryFilepath string
	Duration                           time.Duration
	Results                            []similarityResult
}

type analysisOperation struct {
	img               string
	similarityChannel chan float64
	errorChannel      chan error
}

var analyses chan analysisOperation = make(chan analysisOperation)
var analysesDone chan struct{} = make(chan struct{})

var similarities chan similarityResult = make(chan similarityResult)
var similaritiesDone chan struct{} = make(chan struct{})

var TopThreeSimilarities []similarityResult = make([]similarityResult, 3)

func GetFilepathsFromCommandLineArguments() ImageFilepaths {
	fp := ImageFilepaths{os.Args[1], os.Args[2]}
	return fp
}

func RunAnalysis(filepaths ImageFilepaths) {
	fsImagesToBeAnalysed, _ := images.GetImagesFromFs(os.DirFS(filepaths.ImageDirectory))
	fsComparisonImage, _ := images.GetSingleImage(fsImagesToBeAnalysed, filepaths.ComparisonImage)

	go monitorAnalyses(fsComparisonImage, filepaths.ImageDirectory)
	go handleSimilarities(fsComparisonImage)
	addAnalyses(fsImagesToBeAnalysed)
}

func addAnalysisOperation(img images.Image) (float64, error) {
	similarityChannel := make(chan float64)
	errorChannel := make(chan error)

	op := analysisOperation{
		img:               img.Name,
		similarityChannel: similarityChannel,
		errorChannel:      errorChannel,
	}

	analyses <- op
	return <-similarityChannel, <-errorChannel
}

func monitorAnalyses(referenceImage images.Image, directoryPath string) {
	for op := range analyses {
		go func(op analysisOperation) {
			path := directoryPath + "/" + op.img
			imageData, _ := os.ReadFile(path)
			similarity := analysis.CompareImages(referenceImage.Bytes, imageData)

			similarities <- similarityResult{op.img, similarity}
			op.similarityChannel <- similarity
			op.errorChannel <- nil
		}(op)
	}

	defer close(analysesDone)
}

func handleSimilarities(referenceImage images.Image) {
	for s := range similarities {
		if referenceImage.Name != s.ImageName && s.Similarity >= TopThreeSimilarities[2].Similarity {
			TopThreeSimilarities[2] = s
			sort.SliceStable(TopThreeSimilarities, func(i, j int) bool {
				return TopThreeSimilarities[i].Similarity > TopThreeSimilarities[j].Similarity
			})
		}
	}

	defer close(similaritiesDone)
}

func addAnalyses(imagesToBeAnalysed []images.Image) {
	for _, img := range imagesToBeAnalysed {
		_, err := addAnalysisOperation(img)
		if err != nil {
			fmt.Println(err)
		}
	}

	stopAnalyses()
}

func stopAnalyses() {
	close(analyses)
	close(similarities)
}

func OutputSimilarities(elapsed time.Duration, filepaths ImageFilepaths) {
	data := jsonResult{
		ComparisonImage:   filepaths.ComparisonImage,
		DirectoryFilepath: filepaths.ImageDirectory,
		Duration:          elapsed,
		Results:           TopThreeSimilarities,
	}

	file, _ := json.MarshalIndent(data, "", " ")
	filename := "./tmp/" + strings.TrimSuffix(filepaths.ComparisonImage, ".raw") + ".json"
	_ = ioutil.WriteFile(filename, file, 0644)

	f, err := os.OpenFile("./tmp/output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	str := "Run: " + filepaths.ComparisonImage + " - " + filepaths.ImageDirectory + " - " + " - Duration:  " + fmt.Sprint(elapsed) + "\n"

	if _, err := f.WriteString(str); err != nil {
		fmt.Println(err)
	}
}
