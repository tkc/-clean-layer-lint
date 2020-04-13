package domain

import (
	"encoding/json"
	"strings"
)

type Config struct {
	Path   string   `json:"path"`
	Order  []string `json:"order"`
	Ignore []string `json:"ignore"`
}

type Layer struct {
	Path   string
	Order  map[string]int
	Ignore []string
}

type Mesasge struct {
	Current string
	Targer  string
}

func (l *Layer) UnmarshalJSON(data []byte) error {
	var config = Config{}
	err := json.Unmarshal(data, &config)
	if err != nil {
		return err
	}
	m := make(map[string]int, len(config.Order))
	for i, layer := range config.Order {
		m[layer] = i + 1
	}
	l.Path = config.Path
	l.Order = m
	l.Ignore = config.Ignore
	return nil
}

func (l *Layer) IsCorrectImport(currentPackege string, targetPackeges string) (*Mesasge, bool, error) {
	current := l.Package2Layer(currentPackege)
	target := l.Package2Layer(targetPackeges)

	currentOrder := l.Order[current]
	targetOrder := l.Order[target]

	// if currentOrder == 0 {
	// 	return nil, false, errors.New(current + "not　defined")
	// }

	// if targetOrder == 0 {
	// return nil, false, errors.New(target + "not　defined")
	// }

	if currentOrder == targetOrder+1 {
		return nil, true, nil
	}
	return &Mesasge{Current: current, Targer: target}, false, nil
}

func (l *Layer) IsIgnorePackege(packegeName string) bool {
	for _, p := range l.Ignore {
		if p == packegeName {
			return true
		}
	}
	return false
}

func (l *Layer) IsModulePackege(packegeName string) bool {
	if len(l.Path) == 0 {
		return true
	}
	if strings.Contains(packegeName, l.Path) {
		return true
	}
	return false
}

func (l *Layer) Package2Layer(packegeName string) string {
	replaced := strings.Replace(packegeName, l.Path, "", 1)
	sliced := strings.Split(replaced, "/")
	return sliced[1]
}
