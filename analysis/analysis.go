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

		for j := 0; j < 3; j++ {
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
