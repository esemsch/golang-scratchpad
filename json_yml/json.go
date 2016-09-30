package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var x interface{}
	err := json.Unmarshal([]byte("{\"x\":1,\"y\":2}"), &x)
	if err != nil {
		fmt.Println(os.Stderr, "Error unmarshalling:", err)
	} else {
		fmt.Println(x)
	}

	marsh(x)

	var y struct {
		A string
		B int
	}

	y.A, y.B = "Something", 10

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
