package images_test

import (
	"pixel-challenge/images"
	"reflect"
	"testing"
	"testing/fstest"
)

func assertImage(t *testing.T, got images.Image, want images.Image) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestGetImagesFromFs(t *testing.T) {
	fs := fstest.MapFS{
		"test-1.raw": {Data: []byte("one")},
		"test-2.raw": {Data: []byte("two")},
	}

	imagesFromFs, err := images.GetImagesFromFs(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(imagesFromFs) != len(fs) {
		t.Errorf("got %d images, wanted %d images", len(imagesFromFs), len(fs))
	}

	assertImage(t, imagesFromFs[0], images.Image{Name: "test-1.raw", Bytes: []byte("one")})
}

func TestGetSingleImage(t *testing.T) {
	initImageSlice := func() []images.Image {
		imageOne := images.Image{Name: "Image-1.raw", Bytes: []byte("one")}
		imageTwo := images.Image{Name: "Image-2.raw", Bytes: []byte("two")}
		imageSlice := make([]images.Image, 2)

		imageSlice[0] = imageOne
		imageSlice[1] = imageTwo
		return imageSlice
	}

	t.Run("Returns Image{} on happy path", func(t *testing.T) {
		imageSlice := initImageSlice()
		got, err := images.GetSingleImage(imageSlice, "Image-1.raw")

		if err != nil {
			t.Fatal(err)
		}

		want := images.Image{Name: "Image-1.raw", Bytes: []byte("one")}

		assertImage(t, got, want)
	})

	t.Run("Returns an error when no matching image name", func(t *testing.T) {
		imageSlice := initImageSlice()
		_, err := images.GetSingleImage(imageSlice, "will-fail")

		if err != images.ImageErrors["400"] {
			t.Errorf("got %q, wanted %q", err, images.ImageErrors["400"])
		}
	})

}
