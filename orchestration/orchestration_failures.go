package orchestration

// package orchestration

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

// type analysisOperation struct {
// 	img images.Image
// }

// // type analysisOperation struct {
// // 	img               images.Image
// // 	similarityChannel chan float64
// // 	errorChannel      chan error
// // }

// // var analyses chan analysisOperation = make(chan analysisOperation)
// var analyses []analysisOperation

// // var analysesDone chan struct{} = make(chan struct{})

// var similarities chan similarityResult = make(chan similarityResult)
// var similaritiesDone chan struct{} = make(chan struct{})
// var done chan struct{} = make(chan struct{})

// var TopThreeSimilarities []similarityResult = make([]similarityResult, 3)

// func GetFilepathsFromCommandLineArguments() ImageFilepaths {
// 	fp := ImageFilepaths{os.Args[1], os.Args[2]}
// 	return fp
// }

// func RunAnalysis(filepaths ImageFilepaths) {
// 	fsImagesToBeAnalysed, _ := images.GetImagesFromFs(os.DirFS(filepaths.ImageDirectory))
// 	fsComparisonImage, _ := images.GetSingleImage(fsImagesToBeAnalysed, filepaths.ComparisonImage)

// 	addAnalyses(fsImagesToBeAnalysed, fsComparisonImage)
// 	close(similarities)
// 	<-done
// 	// monitorAnalyses(fsComparisonImage)
// 	// go handleSimilarities(fsComparisonImage)

// }

// func addAnalysisOperation(img images.Image) {
// 	op := analysisOperation{
// 		img: img,
// 	}

// 	analyses = append(analyses, op)
// }

// // func addAnalysisOperation(img images.Image) (float64, error) {
// // 	similarityChannel := make(chan float64)
// // 	errorChannel := make(chan error)

// // 	op := analysisOperation{
// // 		img:               img,
// // 		similarityChannel: similarityChannel,
// // 		errorChannel:      errorChannel,
// // 	}

// // 	analyses <- op
// // 	return <-similarityChannel, <-errorChannel
// // }

// func monitorAnalyses(referenceImage images.Image) {
// 	for _, op := range analyses {
// 		go func(op analysisOperation) {
// 			similarity := analysis.CompareImages(referenceImage.Bytes, op.img.Bytes)
// 			fmt.Println(similarity)
// 			similarities <- similarityResult{op.img.Name, similarity}
// 		}(op)
// 	}
// 	close(similaritiesDone)
// 	// close(analysesDone)
// 	// stopAnalyses()
// }

// // func monitorAnalyses(referenceImage images.Image) {
// // 	for op := range analyses {
// // 		go func(op analysisOperation) {
// // 			similarity := analysis.CompareImages(referenceImage.Bytes, op.img.Bytes)

// // 			similarities <- similarityResult{op.img.Name, similarity}
// // 			op.similarityChannel <- similarity
// // 			op.errorChannel <- nil
// // 		}(op)
// // 	}

// // 	defer close(analysesDone)
// // }

// func handleSimilarities(referenceImage images.Image) {
// 	for s := range similarities {
// 		func(s similarityResult) {
// 			if referenceImage.Name != s.ImageName && s.Similarity >= TopThreeSimilarities[2].Similarity {
// 				TopThreeSimilarities[2] = s
// 				sort.SliceStable(TopThreeSimilarities, func(i, j int) bool {
// 					return TopThreeSimilarities[i].Similarity > TopThreeSimilarities[j].Similarity
// 				})
// 			}
// 		}(s)
// 	}

// 	close(similaritiesDone)

// }

// func addAnalyses(imagesToBeAnalysed []images.Image, referenceImage images.Image) {
// 	for _, imageToBeAnalysed := range imagesToBeAnalysed {
// 		// addAnalysisOperation(img)
// 		go func(imageToBeAnalysed images.Image, referenceImage images.Image) {
// 			similarity := analysis.CompareImages(referenceImage.Bytes, imageToBeAnalysed.Bytes)
// 			fmt.Println(similarity)
// 			similarities <- similarityResult{imageToBeAnalysed.Name, similarity}
// 		}(imageToBeAnalysed, referenceImage)
// 	}

// 	// <-analysesDone
// 	// <-similaritiesDone
// 	// stopAnalyses()
// 	close(done)
// }

// // func addAnalyses(imagesToBeAnalysed []images.Image) {
// // 	for _, img := range imagesToBeAnalysed {
// // 		_, err := addAnalysisOperation(img)
// // 		if err != nil {
// // 			fmt.Println(err)
// // 		}
// // 	}

// // 	stopAnalyses()
// // }

// func stopAnalyses() {
// 	// close(analyses)
// 	close(similarities)
// }

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
// func RunAnalysis(filepaths ImageFilepaths) {
// 	fsImagesToBeAnalysed, _ := images.GetImagesFromFs(os.DirFS(filepaths.ImageDirectory))
// 	fsComparisonImage, _ := images.GetSingleImage(fsImagesToBeAnalysed, filepaths.ComparisonImage)

// 	done := make(chan struct{}, 3)
// 	c := gen(fsImagesToBeAnalysed)
// 	out1 := comp(c, fsComparisonImage, done)
// 	out2 := comp(c, fsComparisonImage, done)
// 	out3 := comp(c, fsComparisonImage, done)
// 	out4 := comp(c, fsComparisonImage, done)

// 	outMerged := merge(done, out1, out2, out3, out4)

// 	generateSimilarityOutput(outMerged, fsComparisonImage)

// 	// go monitorSimilarities(fsComparisonImage.Name)
// 	// addAnalyses(fsImagesToBeAnalysed, fsComparisonImage)
// 	// close(similarities)
// 	// close(similarities)
// 	// <-done
// 	// close(similarities)

// }

// func merge(done chan struct{}, cs ...<-chan similarityResult) <-chan similarityResult {
// 	out := make(chan similarityResult)
// 	defer close(out)
// 	output := func(c <-chan similarityResult) {
// 		for n := range c {
// 			out <- n
// 		}
// 	}

// 	for _, c := range cs {
// 		go output(c)
// 	}

// 	<-done
// 	<-done
// 	<-done
// 	<-done
// 	close(done)
// 	return out
// }

// func gen(imagesToBeAnalysed []images.Image) <-chan images.Image {
// 	out := make(chan images.Image)
// 	go func() {
// 		defer close(out)
// 		for _, imageToBeAnalysed := range imagesToBeAnalysed {
// 			out <- imageToBeAnalysed
// 		}
// 	}()
// 	return out
// }

// func comp(in <-chan images.Image, referenceImage images.Image, done chan struct{}) <-chan similarityResult {
// 	out := make(chan similarityResult)
// 	defer func() {
// 		done <- struct{}{}
// 	}()
// 	go func() {
// 		defer close(out)
// 		for image := range in {
// 			similarity := analysis.CompareImages(referenceImage.Bytes, image.Bytes)
// 			out <- similarityResult{
// 				ImageName:  image.Name,
// 				Similarity: similarity,
// 			}
// 		}
// 	}()
// 	return out
// }

// func generateSimilarityOutput(in <-chan similarityResult, referenceImage images.Image) {
// 	for result := range in {
// 		if referenceImage.Name != result.ImageName && result.Similarity >= TopThreeSimilarities[2].Similarity {
// 			TopThreeSimilarities[2] = result
// 			sort.SliceStable(TopThreeSimilarities, func(i, j int) bool {
// 				return TopThreeSimilarities[i].Similarity > TopThreeSimilarities[j].Similarity
// 			})
// 		}
// 	}

// 	fmt.Println(TopThreeSimilarities)
// }
