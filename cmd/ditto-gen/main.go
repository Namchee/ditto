package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
)

var (
	infoLogger *log.Logger
	errLogger  *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Lmsgprefix)
	errLogger = log.New(os.Stderr, "[ERROR] ", log.Lmsgprefix)
}

func main() {
	cwd, _ := os.Getwd()
	fsys := os.DirFS(cwd)

	config := entity.ReadConfiguration(fsys, infoLogger)

	filename := fmt.Sprintf("%d.json", time.Now().Unix())
	testname := constant.DefaultTestName

	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if len(os.Args) > 2 {
		testname = os.Args[2]
	}

	if !strings.HasSuffix(testname, ".json") {
		testname += ".json"
	}

	infoLogger.Println("Checking test directory availabilty")

	if _, err := fs.Stat(fsys, config.Directory); os.IsNotExist(err) {
		infoLogger.Println("Creating test directory")

		testDir := fmt.Sprintf("%s/%s", cwd, config.Directory)
		err = os.MkdirAll(testDir, os.ModePerm)

		if err != nil {
			errLogger.Fatalln("❌ Failed to create test directory")
		}
	}

	infoLogger.Println("Creating new sample test file")
	filePath := fmt.Sprintf("%s/%s", config.Directory, filename)

	if _, err := fs.Stat(fsys, filePath); os.IsNotExist(err) {
		testDef := fmt.Sprintf(constant.TestTemplate, testname)
		err := os.WriteFile(filePath, []byte(testDef), 0755)

		if err != nil {
			errLogger.Fatalln("❌ Failed to create sample test file.")
		}

		infoLogger.Printf("✔️ Successfully created new sample test file %s\n", filename)
	} else {
		errLogger.Fatalln("❌ Failed to create sample test file. File already exist.")
	}
}
