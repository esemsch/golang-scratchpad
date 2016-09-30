package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type X struct {
	aa string
	bb int
}

func main() {
	var x interface{}
	err := json.Unmarshal([]byte("{\"x\":1,\"y\":2}"), &x)
	if err != nil {
		fmt.Println(os.Stderr, "Error unmarshalling:", err)
	} else {
		fmt.Println(x)
	}

	marsh(x)

	/*	var y struct {
			a string
			b int
		}

		y.a, y.b = "Something", 10*/
	y := X{"Ahoj", 10}

	marsh(y)

}

func marsh(x interface{}) {
	fmt.Println(x)
	jsn, err := json.Marshal(x)
	if err != nil {
		fmt.Println(os.Stderr, "Error marshalling:", err)
	} else {
		fmt.Println(string(jsn))
	}
}
