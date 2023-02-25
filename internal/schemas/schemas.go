package schemas

import (
	"github.com/goccy/go-yaml"
)

type Schema struct {
	Steps []struct {
		Bp2 struct {
			Source     string `yaml:"source"`
			Parameters struct {
				Foo string `yaml:"foo"`
				Bar string `yaml:"bar"`
			} `yaml:"parameters"`
		} `yaml:"bp2"`
	} `yaml:"steps"`
}

func Validator(data []byte) (bool, error) {
	var schema Schema

	if err := yaml.Unmarshal(data, &schema); err != nil {
		return false, err
	}
	return true, nil
}

//Source: https://stackoverflow.com/questions/40737122
