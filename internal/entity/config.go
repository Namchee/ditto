package entity

import (
	"encoding/json"
	"io/fs"
	"log"

	"github.com/Namchee/ditto/internal/constant"
)

// Configuration alters config
type Configuration struct {
	TestDirectory string `json:"test_directory"`
	LogDirectory  string `json:"log_directory"`
	Strict        bool   `json:"strict"`
}

// ReadConfiguration searchs and parses ditto configuration file in the current working directory
func ReadConfiguration(fsys fs.FS, logger *log.Logger) *Configuration {
	config := Configuration{
		TestDirectory: constant.DefaultTestDir,
		LogDirectory:  constant.DefaultLogDir,
	}

	if _, err := fs.Stat(fsys, constant.ConfigurationFilename); err == nil {
		data, err := fsys.Open(constant.ConfigurationFilename)

		if err != nil {
			logger.Println(constant.ErrReadConfig)
			return &config
		}

		err = json.NewDecoder(data).Decode(&config)

		if err != nil {
			logger.Println(constant.ErrDecodeConfig)
		}

		return &config
	}

	logger.Println(constant.ErrNoConfig)
	return &config
}
