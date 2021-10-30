package constant

import "errors"

// Test file errors
var (
	ErrNoDir   = errors.New("[Reader] Tests directory does not exist")
	ErrListDir = errors.New("[Reader] Failed to list file directory")
)

// Fetcher errors
var (
	ErrCreateRequest = errors.New("[Fetcher] Failed to create new request")
	ErrReadResponse  = errors.New("[Fetcher] Failed to read response")
)

// Config errors
var (
	ErrNoConfig      = errors.New("[Configuration] Missing config file. Using default configuration.")
	ErrReadConfig    = errors.New("[Configuration] Failed to read config file. Ignoring configuration file.")
	ErrDecodeConfig  = errors.New("[Configuration] Failed to decode config file. Ignoring configuration file.")
	ErrInvalidConfig = errors.New("[Configuration] Invalid configuration file. Ignoring configuration file.")
)

// Utils errors
var (
	ErrDirExist = errors.New("[MkDir] Failed to create directory. Directory already exists.")
)

// Dynamic error templates
const (
	ErrFileOpen    = "[Reader] Failed to open file %s"
	ErrFileParse   = "[Reader] Failed to parse file %s"
	ErrFileInvalid = "[Reader] Invalid file format for file %s"
	ErrDirFailed   = "[MkDir] Failed to create directory: %s"
)
