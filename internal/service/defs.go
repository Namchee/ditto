package service

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
	"github.com/Namchee/ditto/internal/utils"
)

// GetDefs searchs desired test directory for test definition
func GetDefs(fsys fs.FS, config *entity.Configuration, logger *log.Logger) ([]fs.File, error) {
	if !utils.IsDirExist(fsys, config.TestDirectory) {
		return nil, constant.ErrNoDir
	}

	dir, err := fs.ReadDir(fsys, config.TestDirectory)
	if err != nil {
		return nil, constant.ErrListDir
	}

	var testFiles []fs.File

	for i := range dir {
		name := dir[i].Name()

		if filepath.Ext(name) == ".json" {
			file, err := fsys.Open(
				fmt.Sprintf("%s/%s", config.TestDirectory, name),
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
