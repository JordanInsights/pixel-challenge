package main

import (
	"pixel-challenge/orchestration"
	"time"
)

func main() {
	start := time.Now()

	filepaths := orchestration.GetFilepathsFromCommandLineArguments()
	orchestration.RunAnalysis(filepaths)

	elapsed := time.Since(start)
	orchestration.OutputSimilarities(elapsed, filepaths)
}
