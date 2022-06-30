package images

import "errors"

// ImageErrors contains a list of errors that can be encountered by functionaltiy in the images package
var ImageErrors = map[string]error{
	"400": errors.New("400: Bad Request"),
}
