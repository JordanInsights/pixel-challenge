package orchestration

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// 	"pixel-challenge/analysis"
// 	"pixel-challenge/images"
// 	"sort"
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

// var similarities chan similarityResult = make(chan similarityResult)
// var TopThreeSimilarities []similarityResult = make([]similarityResult, 3)

// func GetFilepathsFromCommandLineArguments() ImageFilepaths {
// 	fp := ImageFilepaths{os.Args[1], os.Args[2]}
// 	return fp
// }

// func RunAnalysis(filepaths ImageFilepaths) {
// 	fsImagesToBeAnalysed, _ := images.GetImagesFromFs(os.DirFS(filepaths.ImageDirectory))
// 	fsComparisonImage, _ := images.GetSingleImage(fsImagesToBeAnalysed, filepaths.ComparisonImage)

// 	addAnalyses(fsImagesToBeAnalysed, fsComparisonImage)
// 	handleSimilarities(fsComparisonImage)
// }

// func handleSimilarities(referenceImage images.Image) {
// 	for s := range similarities {
// 		if referenceImage.Name != s.ImageName && s.Similarity >= TopThreeSimilarities[2].Similarity {
// 			TopThreeSimilarities[2] = s
// 			sort.SliceStable(TopThreeSimilarities, func(i, j int) bool {
// 				return TopThreeSimilarities[i].Similarity > TopThreeSimilarities[j].Similarity
// 			})
// 		}
// 	}
// }

// func addAnalyses(imagesToBeAnalysed []images.Image, comparisonImage images.Image) {
// 	var addition = make(chan int)
// 	var total = make(chan int)
// 	length := len(imagesToBeAnalysed)

// 	go func() {
// 		var count int = 0
// 		for {
// 			select {
// 			case increment := <-addition:
// 				count += increment
// 			case total <- count:
// 			}
// 		}
// 	}()

// 	for _, img := range imagesToBeAnalysed {
// 		go func(img images.Image) {
// 			similarity := analysis.CompareImages(img.Bytes, comparisonImage.Bytes)
// 			similarities <- similarityResult{img.Name, similarity}
// 			addition <- 1
// 			if length == <-total {
// 				close(similarities)
// 			}
// 		}(img)
// 	}
// }
