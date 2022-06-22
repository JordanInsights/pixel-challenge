// package input greets the user and gathers filepaths for image processing
package input

import (
	"fmt"
	"io"
)

func PrintMessage(writer io.Writer, message string) {
	fmt.Fprintf(writer, "%s", message)
}
