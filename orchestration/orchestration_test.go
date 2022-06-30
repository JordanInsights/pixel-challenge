package orchestration_test

import (
	"os"
	"pixel-challenge/orchestration"
	"reflect"
	"testing"
)

func TestGetFilepathsFromCommandLineArguments(t *testing.T) {
	os.Args[1] = "testing"
	os.Args[2] = "testing2"

	got := orchestration.GetFilepathsFromCommandLineArguments()
	want := orchestration.ImageFilepaths{"testing", "testing2"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %+v, want %+v", got, want)
	}
}

func TestAddAnalyses(t *testing.T) {

}

func TestStopAnalyses(t *testing.T) {
	var analyses chan orchestration.AnalysisOperation = make(chan orchestration.AnalysisOperation)
	var similarities chan orchestration.SimilarityResult = make(chan orchestration.SimilarityResult)

	orchestration.StopAnalyses(analyses, similarities)

	select {
	case <-analyses:
	case <-similarities:
	default:
		t.Error("Channel is not closed")
	}
}

// type SpyAnalysisOperationAdder struct {
// 	Calls int
// }

// func (s *SpyAnalysisOperationAdder) addAnalysisOperation() {
// 	s.Calls++
// }

// func TestAddAnalyses(t *testing.T) {
// 	imagesToBeAnalysed := make([]images.Image, 5)
// 	spyAddAnalysisOperation := &SpyAnalysisOperationAdder{}

// 	orchestration.AddAnalyses(imagesToBeAnalysed, spyAddAnalysisOperation)

// 	if spyAddAnalysisOperation.Calls != 5 {
// 		t.Errorf("Not enough calls to sleeper, want 5 got %d", spyAddAnalysisOperation.Calls)
// 	}

// }

// issues with closing channels
// stopping the channel from closing manually gets it to run
// but the results aren't truly representative. Need to come back to

// func BenchmarkRunAnalysis(b *testing.B) {

// 	b.Run("Run the Bronze analysis", func(b *testing.B) {
// 		filepaths := orchestration.ImageFilepaths{
// 			"1d25ea94-4562-4e19-848e-b60f1b58deee.raw",
// 			"./test-images/Bronze",
// 		}

// 		for i := 0; i < b.N; i++ {
// 			orchestration.RunAnalysis(filepaths)
// 		}
// 	})

// 	b.Run("Run the Silver analysis", func(b *testing.B) {
// 		filepaths := orchestration.ImageFilepaths{
// 			"00c0c724-3eb1-472c-b0db-3d6fce3237f7.raw",
// 			"./test-images/Silver",
// 		}

// 		for i := 0; i < b.N; i++ {
// 			orchestration.RunAnalysis(filepaths)
// 		}
// 	})

// 	b.Run("Run the Gold analysis", func(b *testing.B) {
// 		filepaths := orchestration.ImageFilepaths{
// 			"0a0f8f44-3b78-4bff-adee-14bc708e4ba7.raw",
// 			"./test-images/Gold",
// 		}

// 		for i := 0; i < b.N; i++ {
// 			orchestration.RunAnalysis(filepaths)
// 		}
// 	})
// }

// func BenchmarkRunBronzeAnalysis(b *testing.B) {
// 	filepaths := orchestration.ImageFilepaths{
// 		"1d25ea94-4562-4e19-848e-b60f1b58deee.raw",
// 		"./test-images/Bronze",
// 	}

// 	for i := 0; i < b.N; i++ {
// 		orchestration.RunAnalysis(filepaths, true)
// 	}

// }
