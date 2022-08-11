package service

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
)

type runnerResultLog struct {
	Name   string                `json:"name"`
	Err    string                `json:"error"`
	Result []*entity.FetchResult `json:"result"`
}

// WriteTestLog writes test result in case of test fails
func WriteTestLog(
	result *entity.RunnerResult,
	fsys fs.FS,
	config *entity.Configuration,
) error {
	name := fmt.Sprintf("%s.json", result.Name)
	file := fmt.Sprintf("%s/%s", config.LogDirectory, name)

	errMsg := ""

	if result.Error != nil {
		errMsg = result.Error.Error()
	}

	if config.Parse {
		for idx := range result.Responses {
			var temp interface{}

			err := json.Unmarshal([]byte(result.Responses[idx].Response.(string)), &temp)
			if err == nil {
				result.Responses[idx].Response = temp
			}
		}
	}

	runnerLog := runnerResultLog{
		Name:   result.Name,
		Err:    errMsg,
		Result: result.Responses,
	}

	contents, _ := json.MarshalIndent(runnerLog, "", "\t")
	err := os.WriteFile(file, []byte(contents), constant.FilePerms)

	return err
}
