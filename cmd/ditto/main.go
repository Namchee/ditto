package main

import (
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
	fsys := os.DirFS(cwd)

	config := entity.ReadConfiguration(fsys, infoLogger)

	dir, err := service.GetDefs(fsys, config, infoLogger)

	if err != nil {
		errLogger.Fatalln(err)
	}

	files, err := service.ParseData(dir, config, infoLogger)

	if err != nil {
		errLogger.Fatalln(err)
	}

	for _, file := range files {

	}
}
