package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	x := unmarsh("{\"x\":1,\"y\":2}", new(interface{}))

	fmt.Println(x)

	fmt.Println(marsh(x))

	var y struct {
		A string `json:"a"`
		B int    "json:\"b\""
	}

	y.A, y.B = "Something", 10

	fmt.Println(marsh(y))

	z := unmarsh(`{"a":"something","b":100}`, &y)

	fmt.Println(z)

}

func unmarsh(in string, out interface{}) interface{} {
	err := json.Unmarshal([]byte(in), &out)
	if err != nil {
		panic(err)
	} else {
		return out
	}
}

func marsh(in interface{}) string {
	out, err := json.Marshal(in)
	if err != nil {
		panic(err)
	} else {
		return string(out)
	}
}
