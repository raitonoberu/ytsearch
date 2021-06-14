package youtube_search

import "fmt"

// PageDoesntExistError is returned if the page (2+) does not exist.
type PageDoesntExistError struct{}

func (e *PageDoesntExistError) Error() string {
	return "The next page does not exist"
}

// ParsingError is returned if the page could not be parsed.
type ParsingError struct {
	path string
}

func (e *ParsingError) Error() string {
	return fmt.Sprintf("A parsing error occurred while trying to parse %s", e.path)
}
