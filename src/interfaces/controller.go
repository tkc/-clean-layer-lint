package interfaces

import (
	"github.com/tkc/clean-layer-lint/src/usecase"
	"golang.org/x/tools/go/analysis"
)

type controller struct {
	report usecase.Report
}

type Controller interface {
	Analyze(jsonByte []byte, pass *analysis.Pass) (interface{}, error)
}

func NewController() Controller {
	return &controller{}
}

func (c *controller) Analyze(jsonByte []byte, pass *analysis.Pass) (interface{}, error) {
	report := usecase.NewReport()
	layer, err := report.ReadConfig(jsonByte)
	if err != nil {
		return nil, err
	}
	return report.Validate(layer, pass)
}
