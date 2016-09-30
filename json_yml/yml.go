package main

import (
	"fmt"
	"gopkg.in/yaml.v1"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile(`simple.yml`)
	if err != nil {
		panic(err)
	}

	var x struct {
		Very string `yaml:"very"`
	}

	err = yaml.Unmarshal(content, &x)

	if err != nil {
		panic(err)
	}

	fmt.Println(x)

	mrsh, err := yaml.Marshal(x)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(mrsh))

	var mp = map[string]interface{}{
		"something": 10,
		"else": map[string][]string{
			"cool": []string{"stuff", "stuff", "stuff"},
		},
	}
	fmt.Println(mp)

	mrsh, err = yaml.Marshal(mp)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(mrsh))
}
