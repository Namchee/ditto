package utils

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/Namchee/ditto/internal/constant"
)

// IsFileExist checks if a directory exists or not
func IsFileExist(fsys fs.FS, name string) bool {
	_, err := fs.Stat(fsys, name)

	return err == nil
}

// Mkdir creates a new directory
func Mkdir(fsys fs.FS, name string) error {
	if _, err := fs.Stat(fsys, name); os.IsNotExist(err) {
		err = os.MkdirAll(name, os.ModePerm)

		if err != nil {
			return fmt.Errorf(constant.ErrDirFailed, err.Error())
		}

		return nil
	} else {
		return constant.ErrDirExist
	}
}
