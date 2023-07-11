package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"

	"github.com/santhosh-tekuri/jsonschema/v5"
	yaml "gopkg.in/yaml.v3"
)

// A good, old-fashioned hack to allow for automatic resolution of YAML files
// by the JSON Schema validator.
// TODO: Fork jsonschema and allow Loader instances to return things other than
// io.ReadCloser, since internally all the library does is unmarshal the ReadCloser
// anyway.
func localHttpsSchemaLoader(s string) (io.ReadCloser, error) {
	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	f := "./" + u.Path + ".yaml"

	data, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	m := map[string]any{}
	if err := yaml.Unmarshal(data, m); err != nil {
		return nil, err
	}

	d, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	return io.NopCloser(bytes.NewReader(d)), nil
}

func main() {
	jsonschema.Loaders["https"] = localHttpsSchemaLoader

	compiler := jsonschema.NewCompiler()
	kingdomSchema, err := compiler.Compile("https://demesne.qwwqe.xyz/schemas/v1/components/kingdom")
	if err != nil {
		panic(err)
	}

	// fmt.Println(schema)

	// componentTypeSchemaData, err := os.ReadFile("./schemas/components/type.yaml")
	// if err != nil {
	// 	panic(err)
	// }

	// if err =

	filenames, err := filepath.Glob("./sets/**/kingdom/*.yaml")
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
		m := map[string]any{}
		if err := yaml.Unmarshal(rawCard, m); err != nil {
			panic(err)
		}

		fmt.Printf("%s:\n", m["name"])

		if err := kingdomSchema.Validate(m); err != nil {
			fmt.Printf("\tinvalid (%s)\n", err)
		} else {
			fmt.Println("\tvalid")
		}

		// 	fmt.Println("-----RAW-----")

		// 	for k, v := range m {
		// 		fmt.Printf("%v: %v\n", k, v)
		// 	}

		// 	fmt.Println("-----PARSED-----")

		// 	c, err := NewCardFromMap(m)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	// if err := yaml.Unmarshal(rawCard, &c); err != nil {
		// 	// 	panic(err)
		// 	// }

		// 	fmt.Println(c)

		// 	fmt.Println("---")
	}
}
