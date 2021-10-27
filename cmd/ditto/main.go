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
	// Prevent infinite go routine spawn
	runtime.GOMAXPROCS(4)

	cwd, _ := os.Getwd()
	fsys := os.DirFS(cwd)

	infoLogger.Println("Reading configuration file")
	config := entity.ReadConfiguration(fsys, infoLogger)

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
		runner := service.NewTestRunner(d)
		wg.Add(1)

		infoLogger.Printf("Executing test %s", d.Name)
		go runner.RunTest(wg, channel)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	for result := range channel {
		var bodies []string

		for _, r := range result.Result {
			bodies = append(bodies, r.Response)
		}
		diff := utils.GetDiff(bodies)
		pass := len(diff) == 0

		formatted := utils.FormatResult(result, pass)

		fmt.Println(formatted)
	}
}
