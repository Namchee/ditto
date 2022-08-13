package service

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
	"github.com/r3labs/diff/v3"
)

type runnerResultLog struct {
	Name       string                `json:"name"`
	Error      string                `json:"error"`
	Result     []*entity.FetchResult `json:"result"`
	Difference []diff.Changelog      `json:"difference"`
}

// WriteTestLog writes test result in case of test fails
func WriteTestLog(
	result *entity.TestLog,
	fsys fs.FS,
	config *entity.Configuration,
) error {
	name := fmt.Sprintf("%s.json", result.Name)
	file := fmt.Sprintf("%s/%s", config.LogDirectory, name)

	errMsg := ""
	if result.Error != nil {
		errMsg = result.Error.Error()
	}

	runnerLog := runnerResultLog{
		Name:       result.Name,
		Error:      errMsg,
		Result:     result.Responses,
		Difference: result.Diff,
	}

	contents, _ := json.MarshalIndent(runnerLog, "", "\t")
	err := os.WriteFile(file, []byte(contents), constant.FilePerms)

	return err
}
