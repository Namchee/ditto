package constant

import "errors"

// Test file errors
var (
	ErrNoDir   = errors.New("[Reader] Tests directory does not exist")
	ErrListDir = errors.New("[Reader] Failed to list file directory")
)

var (
	ErrCreateRequest = errors.New("[Fetcher] Failed to create new request")
	ErrFetchResponse = errors.New("[Fetcher] Failed to fetch response")
	ErrReadResponse  = errors.New("[Fetcher] Failed to read response")
)

// Dynamic error templates
const (
	ErrFileOpen    = "[Reader] Failed to open file %s"
	ErrFileParse   = "[Reader] Failed to parse file %s"
	ErrFileInvalid = "[Reader] Invalid file format for file %s"
)