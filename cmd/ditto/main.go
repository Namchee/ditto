package main

import (
	"flag"
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
	errLogger = log.New(os.Stderr, "[ERROR]", log.Ldate|log.Ltime|log.Lmsgprefix)
}

func main() {
	// Prevent infinite go routine spawn
	runtime.GOMAXPROCS(4)

	cwd, _ := os.Getwd()
	fs := os.DirFS(cwd)

	strict := flag.Bool("strict", false, "Panics if one of the test file is invalid")
	flag.Parse()

	cfg := &entity.Configuration{Strict: *strict}

	dir, err := service.GetDefs(fs, infoLogger)

	if err != nil {
		errLogger.Fatalln(err)
	}

	files, err := service.ParseData(dir, cfg, infoLogger)

	if err != nil {
		errLogger.Fatalln(err)
	}

	for _, file := range files {

	}
}
