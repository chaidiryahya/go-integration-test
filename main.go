package main

import (
	// golang package
	"flag"
	"log"
	"time"

	// internal package
	"github.com/chaidiryahya/integration-test/app"
)

type flags struct {
	Module string
}

// func (fl *flags) Parse(fs *flag.FlagSet, args []string) error {
// 	fs.StringVar(&fl.Module, "module", "ALL", "app module:API,GQL")
// 	return fs.Parse(args)
// }

// var flg flags
// var module string

var (
	emptyString = ""
	allModule   = "ALL"
)

func main() {
	startTime := time.Now()
	log.Println("[Integration Test] Starting")

	module := flag.String("module", "ALL", "module that want to test")
	flag.Parse()

	err := app.IntegrationTest(*module)
	if err != nil {
		log.Fatalln(err)
	}

	totalProcessTimeInSeconds := time.Since(startTime).Seconds()
	log.Println("[Integration Test] End")
	log.Printf("[Integration Test] Total Process Time : %v seconds", totalProcessTimeInSeconds)

}
