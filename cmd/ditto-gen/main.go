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
	testname := constant.DefaultTestName

	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if len(os.Args) > 2 {
		testname = os.Args[2]
	}

	s := spinner.New(spinner.CharSets[14], time.Second)
	s.Start()

	s.Suffix = "Checking test directory availabilty"

	cwd, _ := os.Getwd()
	fsys := os.DirFS(cwd)

	if _, err := fs.Stat(fsys, constant.TEST_DIR); os.IsNotExist(err) {
		s.Suffix = "Creating test directory"

		testDir := fmt.Sprintf("%s/%s", cwd, constant.TEST_DIR)
		err = os.MkdirAll(testDir, os.ModePerm)

		if err != nil {
			s.FinalMSG = "❌ Failed to create test directory"
			s.Stop()
			return
		}
	}

	s.Suffix = "Creating new sample test file"
	filePath := fmt.Sprintf("%s/%s", constant.TEST_DIR, filename)

	if _, err := fs.Stat(fsys, filePath); os.IsNotExist(err) {
		testDef := fmt.Sprintf(constant.TestTemplate, testname)
		err := os.WriteFile(filePath, []byte(testDef), 0755)

		if err != nil {
			s.FinalMSG = "❌ Failed to create sample test file."
			s.Stop()
			return
		}

		s.FinalMSG = fmt.Sprintf("✔️ Successfully created new sample test file %s", filename)
	} else {
		s.FinalMSG = "❌ Failed to create sample test file. File already exist."
	}

	s.Stop()
}
