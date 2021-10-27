package service

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
	"github.com/Namchee/ditto/internal/utils"
)

// WriteTestLog writes test result in case of test fails
func WriteTestLog(result *entity.RunnerResult, fsys fs.FS, config *entity.Configuration) error {
	name := fmt.Sprintf("%s.json", result.Name)
	file := fmt.Sprintf("%s/%s", config.LogDirectory, name)

	if utils.IsFileExist(fsys, file) {
		return fmt.Errorf(constant.ErrLogExist, name)
	}

	contents, _ := json.MarshalIndent(result.Result, "", "\t")
	err := os.WriteFile(file, []byte(contents), 0755)

	return err
}
