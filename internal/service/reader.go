package service

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
	"github.com/go-playground/validator/v10"
)

// ReadTestData read and parse structured test data from current working directory
func ReadTestData(
	fsys fs.FS,
	cfg *entity.Configuration,
	logger *log.Logger,
) ([]*entity.TestData, error) {
	var data entity.TestData
	var testData []*entity.TestData
	var err error

	structValidator := validator.New()

	if _, err = fs.Stat(fsys, constant.TEST_DIR); os.IsNotExist(err) {
		return nil, constant.ErrNoDir
	}

	dir, err := fs.ReadDir(fsys, constant.TEST_DIR)

	if err != nil {
		return nil, constant.ErrListDir
	}

	for _, entry := range dir {
		name := entry.Name()

		file, err := fsys.Open(
			fmt.Sprintf("%s/%s", constant.TEST_DIR, name),
		)

		if err != nil {
			return nil, fmt.Errorf(constant.ErrFileOpen, name)
		}
		defer file.Close()

		if err = json.NewDecoder(file).Decode(&data); err != nil {
			if cfg.Strict {
				return nil, fmt.Errorf(constant.ErrFileParse, name)
			}

			logger.Printf("Failed to parse file %s, skipping file", name)
			continue
		}

		for i := range data.Endpoints {
			if data.Endpoints[i].Method == "" {
				data.Endpoints[i].Method = "GET"
			}
		}

		if err = structValidator.Struct(data); err != nil {
			if cfg.Strict {
				return nil, fmt.Errorf(constant.ErrFileInvalid, name)
			}

			logger.Printf("Invalid file format for file %s, skipping", name)
			continue
		}

		testData = append(testData, &data)
	}

	return testData, nil
}
