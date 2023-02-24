package main

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Custom struct {
	MyVal string `yaml:"myval"`
}

func main() {
	data := []byte("myval: cavalo")
	var mess Custom
	err := yaml.Unmarshal(data, &mess)
	if err != nil {
		panic(err)
	}
	fmt.Print(mess.MyVal)
}
