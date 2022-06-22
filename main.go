package main

import (
	"os"
	"pixel-challenge/input"
)

func main() {
	input.PrintMessage(os.Stdout, "Hello, welcome to pixel comparison!")
	input.PrintMessage(os.Stdout, "\nPlease enter the filepath of the image you wish to process: ")
	imageFilepath := input.GetFilepath(nil)
	input.PrintMessage(os.Stdout, imageFilepath)
}
