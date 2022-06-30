package orchestration

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"pixel-challenge/analysis"
// 	"pixel-challenge/images"
// 	"strings"
// 	"time"
// )

// type ImageFilepaths struct {
// 	ComparisonImage string
// 	ImageDirectory  string
// }

// type similarityResult struct {
// 	ImageName  string
// 	Similarity float64
// }

// type jsonResult struct {
// 	ComparisonImage, DirectoryFilepath string
// 	Duration                           time.Duration
// 	Results                            []similarityResult
// }

// type analysisOperation struct {
// 	similarityResultChannel chan similarityResult
// 	imageToAnalyse          images.Image
// }

// var analysisOperationsChannel chan analysisOperation = make(chan analysisOperation)
// var similarityResults chan similarityResult = make(chan similarityResult)
// var done chan struct{} = make(chan struct{})

// // var similarities chan similarityResult = make(chan similarityResult)

// // var queue chan func(imageToBeAnalysed, referenceImage images.Image) = make(chan func(imageToBeAnalysed, referenceImage images.Image))
// // var done chan struct{} = make(chan struct{})

// var TopThreeSimilarities []similarityResult = make([]similarityResult, 3)

// func GetFilepathsFromCommandLineArguments() ImageFilepaths {
// 	fp := ImageFilepaths{os.Args[1], os.Args[2]}
// 	return fp
// }

// func RunAnalysis(filepaths ImageFilepaths) {
// 	fsImagesToBeAnalysed, _ := images.GetImagesFromFs(os.DirFS(filepaths.ImageDirectory))
// 	fsComparisonImage, _ := images.GetSingleImage(fsImagesToBeAnalysed, filepaths.ComparisonImage)

// 	handleSimilarities(fsComparisonImage)
// 	for _, imageToAnalyse := range fsImagesToBeAnalysed {

// 		result := analyseImage(imageToAnalyse, fsComparisonImage)
// 		fmt.Println(result)
// 		similarityResults <- result
// 	}
// }

// func analyseImage(imageToAnalyse images.Image, comparisonImage images.Image) similarityResult {
// 	similarityResultChannel := make(chan similarityResult)
// 	a := analysisOperation{
// 		similarityResultChannel: similarityResultChannel,
// 		imageToAnalyse:          imageToAnalyse,
// 	}
// 	analysisOperationsChannel <- a
// 	return <-similarityResultChannel
// }

// func handleSimilarities(comparisonImage images.Image) {
// 	for a := range analysisOperationsChannel {
// 		go func(a analysisOperation) {
// 			similarity := analysis.CompareImages(comparisonImage.Bytes, a.imageToAnalyse.Bytes)
// 			similarityResult := similarityResult{
// 				a.imageToAnalyse.Name,
// 				similarity,
// 			}
// 			a.similarityResultChannel <- similarityResult
// 		}(a)
// 	}
// }

// // func generateSimilarityOutput(imageToAnalyse images.Image, comparisonImage images.Image) float64 {

// // }

// func OutputSimilarities(elapsed time.Duration, filepaths ImageFilepaths) {
// 	data := jsonResult{
// 		ComparisonImage:   filepaths.ComparisonImage,
// 		DirectoryFilepath: filepaths.ImageDirectory,
// 		Duration:          elapsed,
// 		Results:           TopThreeSimilarities,
// 	}

// 	file, _ := json.MarshalIndent(data, "", " ")
// 	filename := "./tmp/" + strings.TrimSuffix(filepaths.ComparisonImage, ".raw") + ".json"
// 	_ = ioutil.WriteFile(filename, file, 0644)

// 	f, err := os.OpenFile("./tmp/output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	defer f.Close()

// 	str := "Run: " + filepaths.ComparisonImage + " - " + filepaths.ImageDirectory + " - " + " - Duration:  " + fmt.Sprint(elapsed) + "\n"

// 	if _, err := f.WriteString(str); err != nil {
// 		fmt.Println(err)
// 	}
// }
