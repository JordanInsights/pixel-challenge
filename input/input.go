// package input greets the user and gathers filepaths for image processing
package input

import (
	"fmt"
	"io"
	"os"
)

// PrintMessage prints a passed message to the terminal
func PrintMessage(writer io.Writer, message string) {
	fmt.Fprintf(writer, "%s", message)
}

// Reads user input from the command line and returns a string
func GetFilepath(in *os.File) string {
	if in == nil {
		in = os.Stdin
	}

	var filepath string
	_, err := fmt.Fscan(in, &filepath)
	if err != nil {
		panic(err)
	}

	return filepath
}

// Not used any more
// Returns imageFilepath and directoryFilepath strings
// func GetImageAndDirectoryFilepaths() (imageFilepath string, directoryFilepath string) {
// 	PrintMessage(os.Stdout, "\nEnter the filepath of the image you wish to find the three closest matches for: ")
// 	imageFilepath = GetFilepath(nil)

// 	PrintMessage(os.Stdout, "\nEnter the filepath for the directory containing the other images to compare against: ")
// 	directoryFilepath = GetFilepath(nil)

// 	return imageFilepath, directoryFilepath
// }
