package usecase

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/tkc/clean-layer-lint/src/domain"
	"golang.org/x/tools/go/analysis"
)

type report struct{}

type Report interface {
	Validate(Layer *domain.Layer, pass *analysis.Pass) (interface{}, error)
	ReadConfig(jsonByte []byte) (*domain.Layer, error)
	IsTargetFile(fileName string) bool
}

func NewReport() Report {
	return &report{}
}

func (r *report) ReadConfig(jsonByte []byte) (*domain.Layer, error) {
	l := &domain.Layer{}
	if err := json.Unmarshal(jsonByte, l); err != nil {
		return nil, err
	}
	return l, nil
}

func (r *report) IsTargetFile(fileName string) bool {
	if !strings.HasSuffix(fileName, ".go") {
		return false
	}
	if strings.HasSuffix(fileName, "_test.go") {
		return false
	}
	return true
}

func (r *report) Validate(Layer *domain.Layer, pass *analysis.Pass) (interface{}, error) {
	currentPackage := pass.Pkg.Path()
	for _, f := range pass.Files {
		if !r.IsTargetFile(pass.Fset.File(f.Pos()).Name()) {
			continue
		}
		for _, i := range f.Imports {
			path, err := strconv.Unquote(i.Path.Value)
			if err != nil {
				return nil, err
			}
			if !Layer.IsIgnorePackege(path) {
				if Layer.IsModulePackege(path) {
					message, isSafe, err := Layer.IsCorrectImport(currentPackage, path)
					if err != nil {
						pass.Reportf(i.Pos(), "%s", err.Error())
					}
					if !isSafe {
						pass.Reportf(i.Pos(), "%s must not import %s", message.Current, message.Targer)
					}
				}
			}
		}
	}
	return nil, nil
}
