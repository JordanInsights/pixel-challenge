package main

import (
	"fmt"
	"pixel-challenge/orchestration"
)

func main() {
	filepaths := orchestration.GetFilepathsFromCommandLineArguments()
	orchestration.RunAnalysis(filepaths)
	fmt.Println(orchestration.TopThreeSimilarities)
}
