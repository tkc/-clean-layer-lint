package main

import (
	"log"
	"github.com/tkc/clean-layer-lint/src/infrastructure"
	"github.com/tkc/clean-layer-lint/src/interfaces"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func init() {
	log.Print("clean-layer-lint start...")
}

func main() {
	var (
		controller = interfaces.NewController()
		analyzer   = infrastructure.NewAnalyzer(controller)
	)
	singlechecker.Main(analyzer)
}
