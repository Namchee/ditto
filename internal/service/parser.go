package service

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
	"github.com/go-playground/validator/v10"
)

// ParseData read and parse structured test data from current working directory
func ParseData(
	files []fs.File,
	cfg *entity.Configuration,
	logger *log.Logger,
) ([]*entity.TestData, error) {
	var testData []*entity.TestData
	var err error

	structValidator := validator.New()

	for _, file := range files {
		f, _ := file.Stat()
		var data entity.TestData

		if err = json.NewDecoder(file).Decode(&data); err != nil {
			if cfg.Strict {
				return nil, fmt.Errorf(constant.ErrFileParse, f.Name())
			}

			logger.Printf("Failed to parse file %s, skipping file", f.Name())
			continue
		}

		for i := range data.Endpoints {
			if data.Endpoints[i].Method == "" {
				data.Endpoints[i].Method = "GET"
			}
		}

		if err = structValidator.Struct(data); err != nil {
			if cfg.Strict {
				return nil, fmt.Errorf(constant.ErrFileInvalid, f.Name())
			}

			logger.Printf("Invalid file format for file %s, skipping", f.Name())
			continue
		}

		testData = append(testData, &data)
	}

	return testData, nil
}
