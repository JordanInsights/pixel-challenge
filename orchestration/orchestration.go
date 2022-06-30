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

// ImageFilepaths is a struct that contains the filepath for the directory and image being analysed
type ImageFilepaths struct {
	ComparisonImage string
	ImageDirectory  string
}

type SimilarityResult struct {
	ImageName  string
	Similarity float64
}

type jsonResult struct {
	ComparisonImage, DirectoryFilepath string
	Duration                           time.Duration
	Results                            []SimilarityResult
}

type AnalysisOperation struct {
	img               string
	similarityChannel chan float64
	errorChannel      chan error
}

var analyses chan AnalysisOperation = make(chan AnalysisOperation)
var analysesDone chan struct{} = make(chan struct{})

var similarities chan SimilarityResult = make(chan SimilarityResult)
var similaritiesDone chan struct{} = make(chan struct{})

// TopThreeSimilarities is a slice of length 3 that contains the most similar images to the one being analysed
var TopThreeSimilarities []SimilarityResult = make([]SimilarityResult, 3)

// GetFilepathsFromCommandLineArguments gets filepath strings from the command line
func GetFilepathsFromCommandLineArguments() ImageFilepaths {
	fp := ImageFilepaths{os.Args[1], os.Args[2]}
	return fp
}

// RunAnalysis gets the comparison images, and initialises handle and add functions
func RunAnalysis(filepaths ImageFilepaths) {
	comparisonImageData, _ := os.ReadFile(filepaths.ImageDirectory + "/" + filepaths.ComparisonImage)
	comparisonImage := images.Image{
		Name:  filepaths.ComparisonImage,
		Bytes: comparisonImageData,
	}

	go handleAnalyses(comparisonImage, filepaths.ImageDirectory)
	go handleSimilarities(comparisonImage)

	addAnalyses(filepaths.ImageDirectory)
	StopAnalyses(analyses, similarities)
}

// addAnalyses invokes addAnalysisOperation for each image file in the given directory
func addAnalyses(directoryName string) {
	imagesToBeAnalysed, _ := os.ReadDir(directoryName)

	for _, file := range imagesToBeAnalysed {
		_, err := addAnalysisOperation(file.Name())
		if err != nil {
			fmt.Println(err)
		}
	}
}

// addAnalysisOperation adds an AnalysisOperation to the analyses channel to be executed in handleAnalyses
func addAnalysisOperation(imageName string) (float64, error) {
	similarityChannel := make(chan float64)
	errorChannel := make(chan error)

	op := AnalysisOperation{
		img:               imageName,
		similarityChannel: similarityChannel,
		errorChannel:      errorChannel,
	}

	analyses <- op
	return <-similarityChannel, <-errorChannel
}

// handleAnalyses iterates over the analyses channel, executes CompareImages for each request, and sends the result to the similarities channel
func handleAnalyses(referenceImage images.Image, directoryPath string) {
	for op := range analyses {
		go func(op AnalysisOperation) {
			path := directoryPath + "/" + op.img
			imageData, _ := os.ReadFile(path)
			similarity := analysis.CompareImages(referenceImage.Bytes, imageData)

			similarities <- SimilarityResult{op.img, similarity}
			op.similarityChannel <- similarity
			op.errorChannel <- nil
		}(op)
	}

	defer close(analysesDone)
}

// handleSimilarities iterates over the similarities channel and adds results to the TopThreeSimilaties slice if appropriate
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

// StopAnalyses closes the analyses and similarity channels
func StopAnalyses(analyses chan AnalysisOperation, similarities chan SimilarityResult) {
	close(analyses)
	close(similarities)
}

// OutputSimilarities writes a verbose output to a JSON file and adds a line to the output.txt file in /tmp
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
