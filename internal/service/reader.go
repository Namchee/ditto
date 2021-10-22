package service

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
)

// ReadTestData read and parse structured test data from current working directory
func ReadTestData(sys fs.FS) ([]*entity.TestData, error) {
	var data entity.TestData
	var testData []*entity.TestData

	dir, err := fs.ReadDir(sys, constant.TEST_DIR)

	if err != nil {
		return nil, err
	}

	for _, entry := range dir {
		file, err := sys.Open(
			fmt.Sprintf("%s/%s", constant.TEST_DIR, entry.Name()),
		)

		if err != nil {
			return nil, err
		}
		defer file.Close()

		byteContent, err := ioutil.ReadAll(file)

		if err != nil {
			return nil, err
		}

		json.Unmarshal(byteContent, &data)
		testData = append(testData, &data)
	}

	return testData, nil
}
