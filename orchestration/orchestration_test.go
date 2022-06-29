package orchestration_test

import (
	"pixel-challenge/orchestration"
	"testing"
)

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

func BenchmarkRunBronzeAnalysis(b *testing.B) {
	filepaths := orchestration.ImageFilepaths{
		"1d25ea94-4562-4e19-848e-b60f1b58deee.raw",
		"./test-images/Bronze",
	}

	for i := 0; i < b.N; i++ {
		orchestration.RunAnalysis(filepaths)
	}

}
