//import "gopkg.in/yaml.v3"

package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

var file = "./test.yaml"

func main() {

	var manifest yaml.Node
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &manifest)
	if err != nil {
		fmt.Printf("cannot unmarshal file content: %w", err)
	}
	readValues(manifest)
}

func readValues(m yaml.Node) {
	for _, data := range m.Content {

		if data.Value != "" {
			fmt.Printf("line: %d, value: %s\n", data.Line, data.Value)
		}
		readValues(*data)
	}
}
