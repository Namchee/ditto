package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/Namchee/ditto/internal/entity"
	"github.com/Namchee/ditto/internal/service"
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
	infoLogger.Println("Initializing ditto")

	cwd, _ := os.Getwd()
	fsys := os.DirFS(cwd)

	infoLogger.Println("Reading configuration file")
	config := entity.ReadConfiguration(fsys, infoLogger)

	// Prevent infinite go routine spawn
	if config.Worker > 0 {
		runtime.GOMAXPROCS(config.Worker)
	}

	infoLogger.Println("Reading test files")
	files, err := service.GetDefs(fsys, config, infoLogger)

	if err != nil {
		errLogger.Fatalln(err)
	}

	if len(files) == 0 {
		infoLogger.Println("No test to run.")
		os.Exit(0)
	}

	infoLogger.Println("Parsing test files")
	data, err := service.ParseData(files, config, infoLogger)

	if err != nil {
		errLogger.Fatalln(err)
	}

	channel := make(chan *entity.RunnerResult, len(data))
	wg := &sync.WaitGroup{}

	infoLogger.Println("Running tests")
	for _, d := range data {
		runner := service.NewTestRunner(d, config)
		wg.Add(1)

		infoLogger.Printf("Executing test %s", d.Name)
		go runner.RunTest(wg, channel)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	var logs []*entity.TestLog

	for result := range channel {
		hasDiff, diff := utils.HasDiff(result.Responses, config)

		isFailed := result.Error != nil || hasDiff

		formatted := utils.FormatResult(result, !isFailed)
		fmt.Println(formatted)

		if isFailed {
			log := &entity.TestLog{
				RunnerResult: *result,
				Diff:         diff,
			}

			logs = append(logs, log)
		}
	}

	if len(logs) > 0 {
		infoLogger.Println("Writing test logs")

		if !utils.IsFileExist(fsys, config.LogDirectory) {
			err := utils.Mkdir(fsys, config.LogDirectory)

			if err != nil {
				errLogger.Fatalln(err)
			}
		}

		for _, fail := range logs {
			err := service.WriteTestLog(fail, fsys, config)

			if err != nil {
				infoLogger.Printf(
					"[Log] Failed to write test logs for %s. Skipping...",
					fail.Name,
				)
			}
		}
	}

	if len(logs) > 0 {
		fmt.Println("❌ Tests failed")
	} else {
		fmt.Println("✅ All tests passed")
	}
}
