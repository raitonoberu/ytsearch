package ytsearch

import (
	"errors"
)

var (
	// PageDoesntExistError is returned if the page does not exist.
	PageDoesntExistError = errors.New("the page does not exist")
)
