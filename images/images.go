package images

import (
	"io"
	"io/fs"
)

type Image struct {
	Name  string
	Bytes []byte
}

// Returns a slice of images from the local fs when passed valid relative filepath to a directory
func GetImagesFromFs(fileSystem fs.FS) ([]Image, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var images []Image
	for _, f := range dir {
		image, err := getImage(fileSystem, f.Name())
		if err != nil {
			return nil, err // may be incorrect
		}
		images = append(images, image)
	}
	return images, nil
}

// Returns a single image from the local fs when passed valid relative filepath
func GetSingleImage(images []Image, singleImage string) (Image, error) {
	return Image{}, nil
}

// Reads a single image and returns invoked newImage resulting in Image{} struct
func getImage(fileSystem fs.FS, imageName string) (Image, error) {
	imageFile, err := fileSystem.Open(imageName)
	if err != nil {
		return Image{}, err
	}
	defer imageFile.Close()
	return newImage(imageFile, imageName)
}

// takes an image file, reads the data and returns an image struct containing image name and bytes
func newImage(imageFile io.Reader, imageName string) (Image, error) {
	imageData, err := io.ReadAll(imageFile)
	if err != nil {
		return Image{}, err
	}

	image := Image{Name: imageName, Bytes: imageData}
	return image, nil
}
