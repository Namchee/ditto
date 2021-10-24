package main

import (
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/briandowns/spinner"
)

func main() {
	filename := fmt.Sprintf("%d.json", time.Now().Unix())

	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Second)
	s.Start()

	s.Suffix = "Checking test directory availabilty"

	cwd, _ := os.Getwd()
	fsys := os.DirFS(cwd)

	if _, err := fs.Stat(fsys, constant.TEST_DIR); os.IsNotExist(err) {
		s.Suffix = "Creating test directory"

		err = os.MkdirAll(cwd+"/"+constant.TEST_DIR, os.ModePerm)

		if err != nil {
			s.Suffix = "Failed to create test directory"
			s.Stop()
			return
		}
	}

	s.Suffix = "Creating new sample test file"

	filepath := cwd + "\\" + constant.TEST_DIR + "\\" + filename

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		err := os.WriteFile(filepath, []byte(constant.TestTemplate), 0755)

		if err != nil {
			s.Suffix = "Failed to create sample test file."
			s.Stop()
			return
		}

		s.Suffix = fmt.Sprintf("Successfully created new sample test file %s", filename)
		s.Stop()

	} else {
		s.Suffix = "Failed to create sample test file. File already exist"
		s.Stop()
		return
	}
}
