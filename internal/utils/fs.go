package utils

import (
	"io/fs"
)

// IsDirExist checks if a directory exists or not
func IsDirExist(fsys fs.FS, name string) bool {
	_, err := fs.Stat(fsys, name)

	return err == nil
}
