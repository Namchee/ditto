package main

import (
	"log"
	"os"

	"github.com/Namchee/ditto/internal/service"
)

func main() {
	cwd, err := os.Getwd()

	if err != nil {
		log.Fatalln("wut")
	}
	fs := os.DirFS(cwd)

	service.ReadTestData(fs)
}
