package analysis

import (
	"bytes"
)

const BytesPerPixel int = 3

// Given a slice of bytes and a number of bytes per pixel, determines the percentage value of a single pixel match
func DetermineSimilarityIncrement(bytes []byte) float64 {
	var numberOfPixels float64 = float64(len(bytes) / BytesPerPixel)
	var incrementPerPixel float64 = 1.00 / numberOfPixels

	return incrementPerPixel
}

// Compares a byte to another byte and returns a boolean representing presence of a match
func CompareBytes(byteOne byte, byteTwo byte) bool {
	return byteOne == byteTwo
}

// Checks images are of same size and iterates over 'pixels'. Increments similarity if a set of bytes denoting a pixel match the comparison set of bytes
func CompareImages(comparisonImage, imageToAnalyse []byte) (similarity float64) {
	if len(comparisonImage) != len(imageToAnalyse) {
		return 0
	}

	similarityIncrement := DetermineSimilarityIncrement(comparisonImage)
	iterations := len(comparisonImage) / BytesPerPixel

	for i := 0; i < iterations; i++ {
		startIndex := i * BytesPerPixel
		bytesMatch := true

		for j := 0; j < BytesPerPixel; j++ {
			currentIndex := startIndex + j
			if comparisonImage[currentIndex] != imageToAnalyse[currentIndex] {
				bytesMatch = false
				break
			}
		}

		if bytesMatch {
			similarity += similarityIncrement
		}
	}

	// fmt.Printf("\nImage had a similaity of %+v", similarity)
	return similarity
}

// Checks images are of same size, iterates over bytes and decrements a similarity score if a set of bytes denoting a pixel don't match a comparison image
func CompareImagesInverted(comparisonImage, imageToAnalyse []byte) float64 {
	if len(comparisonImage) != len(imageToAnalyse) {
		return 0
	}

	var similarity float64 = 1
	similarityDecrement := DetermineSimilarityIncrement(comparisonImage)
	iterations := len(comparisonImage) / BytesPerPixel

	for i := 0; i < iterations; i++ {
		startIndex := i * BytesPerPixel
		bytesMatch := true

		for j := 0; j < BytesPerPixel; j++ {
			currentIndex := startIndex + j
			if comparisonImage[currentIndex] != imageToAnalyse[currentIndex] {
				bytesMatch = false
				break
			}
		}

		if !bytesMatch {
			similarity -= similarityDecrement
		}

	}

	return similarity
}

// Iterates over a set of bytes, converting a set of given length into an array. Compares arrays and increments similarity if they match
func CompareArrays(comparisonImage, imageToAnalyse []byte) float64 {
	if len(comparisonImage) != len(imageToAnalyse) {
		return 0
	}

	similarityIncrement := DetermineSimilarityIncrement(comparisonImage)
	iterations := len(comparisonImage) / BytesPerPixel

	var similarity float64 = 0
	for i := 0; i < iterations; i++ {
		startIndex := i * BytesPerPixel
		endIndex := startIndex + BytesPerPixel

		arrayA := *(*[BytesPerPixel]byte)(comparisonImage[startIndex:endIndex])
		arrayB := *(*[BytesPerPixel]byte)(imageToAnalyse[startIndex:endIndex])

		if arrayA == arrayB {
			similarity += similarityIncrement
		}
	}

	return similarity
}

// Iterates over a set of bytes, converting a set of given length into a slice. Compares slices and increments similarity if they match
func CompareSlices(comparisonImage, imageToAnalyse []byte) float64 {
	if len(comparisonImage) != len(imageToAnalyse) {
		return 0
	}

	similarityIncrement := DetermineSimilarityIncrement(comparisonImage)
	iterations := len(comparisonImage) / BytesPerPixel

	var similarity float64 = 0
	for i := 0; i < iterations; i++ {
		startIndex := i * BytesPerPixel
		endIndex := startIndex + BytesPerPixel

		sliceA := comparisonImage[startIndex:endIndex]
		sliceB := imageToAnalyse[startIndex:endIndex]

		if byteSlicesEqual(sliceA, sliceB) {
			similarity += similarityIncrement
		}
	}

	return similarity
}

// Iterates over a set of bytes and compares to a comparison image useing the Equal method from the bytes package. Increments a similarity score if they match
func CompareUsingEqual(comparisonImage, imageToAnalyse []byte) float64 {
	if len(comparisonImage) != len(imageToAnalyse) {
		return 0
	}

	similarityIncrement := DetermineSimilarityIncrement(comparisonImage)
	iterations := len(comparisonImage) / BytesPerPixel

	var similarity float64 = 0
	for i := 0; i < iterations; i++ {
		startIndex := i * BytesPerPixel
		endIndex := startIndex + BytesPerPixel

		sliceA := comparisonImage[startIndex:endIndex]
		sliceB := imageToAnalyse[startIndex:endIndex]

		if bytes.Equal(sliceA, sliceB) {
			similarity += similarityIncrement
		}
	}

	return similarity
}

// Iterates over a slice and compares byte values to those at the same index of second slice
func byteSlicesEqual(a, b []byte) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
