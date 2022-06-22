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
