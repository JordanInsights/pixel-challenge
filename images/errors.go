package images

import "errors"

var ImageErrors = map[string]error{
	"400": errors.New("400: Bad Request"),
}
