package service

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/Namchee/ditto/internal/constant"
)

// GetTestFiles searchs desired test directory for test files
func GetTestFiles(fsys fs.FS, logger *log.Logger) ([]fs.File, error) {
	if _, err := fs.Stat(fsys, constant.TEST_DIR); os.IsNotExist(err) {
		return nil, constant.ErrNoDir
	}

	dir, err := fs.ReadDir(fsys, constant.TEST_DIR)

	if err != nil {
		return nil, constant.ErrListDir
	}

	var testFiles []fs.File

	for i := range dir {
		name := dir[i].Name()

		if filepath.Ext(name) == "json" {
			file, err := fsys.Open(
				fmt.Sprintf("%s/%s", constant.TEST_DIR, name),
			)

			if err != nil {
				logger.Printf("[Directory] Unable to read file %s, skipping", name)
				continue
			}

			testFiles = append(testFiles, file)
		}
	}

	return testFiles, nil
}
