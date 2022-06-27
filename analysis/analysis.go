package analysis

const BytesPerPixel int = 3

func DetermineSimilarityIncrement(bytes []byte) float64 {
	var numberOfPixels float64 = float64(len(bytes) / BytesPerPixel)
	var incrementPerPixel float64 = 1.00 / numberOfPixels

	return incrementPerPixel
}

func CompareBytes(byteOne byte, byteTwo byte) bool {
	return byteOne == byteTwo
}

func CompareImages(comparisonImage, imageToAnalyse []byte) (similarity float64) {
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

func CompareImagesInverted(comparisonImage, imageToAnalyse []byte, iterations int, similarityDecrement float64, minimumSimlarity float64) {
	var similarity float64 = 1
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

		if similarity < minimumSimlarity {
			break
		}
	}
}

func CompareArrays(comparisonImage, imageToAnalyse []byte, iterations int, similarityIncrement float64) float64 {
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

func CompareSlices(comparisonImage, imageToAnalyse []byte, iterations int, similarityIncrement float64) float64 {
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

func byteSlicesEqual(a, b []byte) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
