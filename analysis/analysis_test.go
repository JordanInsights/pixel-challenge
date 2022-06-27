package analysis_test

import (
	"os"
	"pixel-challenge/analysis"
	"testing"
)

func TestDetermineSimilarityIncrement(t *testing.T) {
	t.Run("Increments by 0.01 when 300 bytes", func(t *testing.T) {
		bytes := [300]byte{}
		got := analysis.DetermineSimilarityIncrement(bytes[:])
		want := 0.01

		if got != want {
			t.Errorf("got %+v similarity increment, wanted %+v similarity increment", got, want)
		}
	})
}

func TestCompareBytes(t *testing.T) {
	t.Run("Returns true when bytes match", func(t *testing.T) {
		byteOne := byte(1)
		byteTwo := byte(1)

		got := analysis.CompareBytes(byteOne, byteTwo)
		want := true

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})

	t.Run("Returns false when bytes don't match", func(t *testing.T) {
		byteOne := byte(1)
		byteTwo := byte(2)

		got := analysis.CompareBytes(byteOne, byteTwo)
		want := false

		if got != want {
			t.Errorf("got %t, want %t", got, want)
		}
	})
}

func TestCompareImages(t *testing.T) {
	t.Run("Same image passed twice results in similarity of 1.00", func(t *testing.T) {
		comparisonImageData, err := os.ReadFile("../test-images/Bronze/1d25ea94-4562-4e19-848e-b60f1b58deee.raw")
		if err != nil {
			t.Fatal(err)
		}

		imageToAnalyseData, err := os.ReadFile("../test-images/Bronze/1d25ea94-4562-4e19-848e-b60f1b58deee.raw")
		if err != nil {
			t.Fatal(err)
		}

		got := analysis.CompareImages(comparisonImageData, imageToAnalyseData)
		var want float64 = 1

		if got != want {
			t.Errorf("got %+v similarity, wanted %+v similarity", got, want)
		}
	})
}

func BenchmarkCompareImages(b *testing.B) {
	comparisonImageData, err := os.ReadFile("../test-images/Gold/0a0f8f44-3b78-4bff-adee-14bc708e4ba7.raw")
	if err != nil {
		b.Fatal(err)
	}

	imageToAnalyseData, err := os.ReadFile("../test-images/Gold/0a1de745-e548-4676-a954-e445bf7d0182.raw")
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		analysis.CompareImages(comparisonImageData, imageToAnalyseData)
	}
}

func BenchmarkCompareImagesInverted(b *testing.B) {
	comparisonImageData, err := os.ReadFile("../test-images/Gold/0a0f8f44-3b78-4bff-adee-14bc708e4ba7.raw")
	if err != nil {
		b.Fatal(err)
	}

	imageToAnalyseData, err := os.ReadFile("../test-images/Gold/0a1de745-e548-4676-a954-e445bf7d0182.raw")
	if err != nil {
		b.Fatal(err)
	}

	iterations := len(comparisonImageData) / analysis.BytesPerPixel
	similarityDecrement := analysis.DetermineSimilarityIncrement(comparisonImageData)
	minimumSimlarity := 0.0629739761352539

	for i := 0; i < b.N; i++ {
		analysis.CompareImagesInverted(comparisonImageData, imageToAnalyseData, iterations, similarityDecrement, minimumSimlarity)
	}
}

func BenchmarkCompareArrays(b *testing.B) {
	comparisonImageData, err := os.ReadFile("../test-images/Gold/0a0f8f44-3b78-4bff-adee-14bc708e4ba7.raw")
	if err != nil {
		b.Fatal(err)
	}

	imageToAnalyseData, err := os.ReadFile("../test-images/Gold/0a1de745-e548-4676-a954-e445bf7d0182.raw")
	if err != nil {
		b.Fatal(err)
	}

	iterations := len(comparisonImageData) / analysis.BytesPerPixel
	similarityIncrement := analysis.DetermineSimilarityIncrement(comparisonImageData)

	for i := 0; i < b.N; i++ {
		analysis.CompareArrays(comparisonImageData, imageToAnalyseData, iterations, similarityIncrement)
	}
}
