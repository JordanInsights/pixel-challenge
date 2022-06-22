package images

import (
	"fmt"
	"io"
	"io/fs"
)

type Image struct {
	Name  string
	Bytes string
}

func GetImagesFromFs(fileSystem fs.FS) ([]Image, error) {
	dir, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	var images []Image
	for _, f := range dir {
		fmt.Println(f)
		image, err := getImage(fileSystem, f.Name())
		if err != nil {
			return nil, err // may be incorrect
		}
		images = append(images, image)
	}
	return images, nil
}

func getImage(fileSystem fs.FS, imageName string) (Image, error) {
	imageFile, err := fileSystem.Open(imageName)
	if err != nil {
		return Image{}, err
	}
	defer imageFile.Close()
	return newImage(imageFile, imageName)
}

func newImage(imageFile io.Reader, imageName string) (Image, error) {
	imageData, err := io.ReadAll(imageFile)
	if err != nil {
		return Image{}, err
	}

	image := Image{Name: imageName, Bytes: string(imageData)}
	return image, nil
}
