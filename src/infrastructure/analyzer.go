package infrastructure

import (
	"io/ioutil"

	"github.com/tkc/clean-layer-lint/src/interfaces"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var analyzer AnalyzerInfrastructure

type analyzerInfrastructure struct {
	controller interfaces.Controller
}

type AnalyzerInfrastructure interface {
	Run(jsonByte []byte, pass *analysis.Pass) (interface{}, error)
}

func (a *analyzerInfrastructure) Run(jsonByte []byte, pass *analysis.Pass) (interface{}, error) {
	return a.controller.Analyze(jsonByte, pass)
}

func run(pass *analysis.Pass) (interface{}, error) {
	layerJSON, err := ioutil.ReadFile("./clean-layer.json")
	if err != nil {
		return nil, nil
	}
	return analyzer.Run(layerJSON, pass)
}

func NewAnalyzer(controller interfaces.Controller) *analysis.Analyzer {
	analyzer = &analyzerInfrastructure{controller: controller}
	return &analysis.Analyzer{
		Name: "cleanLayerLint",
		Doc:  "...",
		Run:  run,
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
	}
}
