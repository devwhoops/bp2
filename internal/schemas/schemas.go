package schemas

import (
	"fmt"

	"github.com/goccy/go-yaml"
)

func Validator(data []byte) (bool, error) {
	yml := []byte("foo: 1\nbar: carloslindao")

	var schema struct {
		Foo int    `yaml:"foo"`
		Bar string `yaml:"bar"`
	}

	if err := yaml.Unmarshal(yml, &schema); err != nil {
		return false, err
	}
	fmt.Println(schema.Bar, "hello")
	return true, nil
}

//Source: https://stackoverflow.com/questions/40737122
