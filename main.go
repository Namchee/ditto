package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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
	cwd, _ := os.Getwd()
	fs := os.DirFS(cwd)

	strict := flag.Bool("strict", false, "Panics if one of the test file is invalid")
	flag.Parse()

	cfg := &entity.Configuration{Strict: *strict}

	files, err := service.ReadTestData(fs, cfg, infoLogger)

	if err != nil {
		errLogger.Fatalln(err)
	}

	fmt.Println(files)
}
