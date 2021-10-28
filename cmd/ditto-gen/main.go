package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Namchee/ditto/internal/constant"
	"github.com/Namchee/ditto/internal/entity"
	"github.com/Namchee/ditto/internal/utils"
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
	infoLogger.Println("Initializing ditto-gen")
	cwd, _ := os.Getwd()
	fsys := os.DirFS(cwd)

	infoLogger.Println("Reading configuration files")
	config := entity.ReadConfiguration(fsys, infoLogger)

	filename := fmt.Sprintf("%d.json", time.Now().Unix())
	testname := constant.DefaultTestName

	infoLogger.Println("Reading shell inputs")
	if len(os.Args) > 1 {
		filename = os.Args[1]
	}

	if len(os.Args) > 2 {
		testname = os.Args[2]
	}

	if !strings.HasSuffix(filename, ".json") {
		filename += ".json"
	}

	infoLogger.Println("Checking test directory availabilty")

	if !utils.IsFileExist(fsys, config.TestDirectory) {
		err := utils.Mkdir(fsys, config.TestDirectory)

		if err != nil {
			errLogger.Fatalln(err)
		}
	}

	infoLogger.Println("Creating new sample test file")
	filePath := fmt.Sprintf("%s/%s", config.TestDirectory, filename)

	if !utils.IsFileExist(fsys, filePath) {
		testDef := fmt.Sprintf(constant.TestTemplate, testname)
		err := os.WriteFile(filePath, []byte(testDef), constant.FilePerms)

		if err != nil {
			errLogger.Fatalln("❌ Failed to create sample test file.")
		}

		infoLogger.Printf("✅ Successfully created new sample test file %s\n", filename)
	} else {
		errLogger.Fatalln("❌ Failed to create sample test file. File already exist.")
	}
}
