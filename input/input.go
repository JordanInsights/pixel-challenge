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

// Reads user input from the command line and returns string
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
