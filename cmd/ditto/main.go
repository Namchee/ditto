package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/Namchee/ditto/internal/entity"
	"github.com/Namchee/ditto/internal/service"
)

var (
	infoLogger *log.Logger
	errLogger  *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "[INFO] ", log.Lmsgprefix)
	errLogger = log.New(os.Stderr, "[ERROR] ", log.Lmsgprefix)
}

func executeRunner(
	runner *service.TestRunner,
	channel chan<- *entity.TestResult,
) {

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
	dir, err := service.GetDefs(fsys, config, infoLogger)

	if err != nil {
		errLogger.Fatalln(err)
	}

	infoLogger.Println("Parsing test files")
	files, err := service.ParseData(dir, config, infoLogger)

	if err != nil {
		errLogger.Fatalln(err)
	}

	channel := make(chan *entity.TestResult, len(files))

	infoLogger.Println("Running tests")
	for _, file := range files {
		runner := service.NewTestRunner(file)

		infoLogger.Printf("Executing test %s", file.Name)
		go executeRunner(runner, channel)
	}

	for i := 0; i < len(files); i++ {
		result := <-channel
		formatted := service.FormatResult(result)

		fmt.Println(formatted)
	}
}
