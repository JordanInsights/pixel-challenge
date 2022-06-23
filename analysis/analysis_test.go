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
