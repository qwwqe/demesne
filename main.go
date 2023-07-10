package main

import (
	"fmt"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v3"
	"github.com/santhosh-tekuri/jsonschema/v5"
)

func main() {
	compiler := jsonschema.NewCompiler()
	componentTypeSchemaData, err := os.ReadFile("./schemas/components/type.yaml")
	if err != nil {
		panic(err)
	}

	if err =

	filenames, err := filepath.Glob("./sets/dominion/cards/*.yaml")
	if err != nil {
		panic(err)
	}

	rawCards := make([][]byte, 0, len(filenames))
	for _, filename := range filenames {
		data, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		rawCards = append(rawCards, data)
	}

	for _, rawCard := range rawCards {
		m := map[any]any{}
		if err := yaml.Unmarshal(rawCard, m); err != nil {
			panic(err)
		}

		fmt.Println("-----RAW-----")

		for k, v := range m {
			fmt.Printf("%v: %v\n", k, v)
		}

		fmt.Println("-----PARSED-----")

		c, err := NewCardFromMap(m)
		if err != nil {
			panic(err)
		}
		// if err := yaml.Unmarshal(rawCard, &c); err != nil {
		// 	panic(err)
		// }

		fmt.Println(c)

		fmt.Println("---")
	}
}
