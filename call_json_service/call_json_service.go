package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := `https://openlibrary.org/api/get?text=true&prettyprint=true&stats=true&key=%2F`
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got %d from %s\n", resp.StatusCode, url)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var jsn interface{}
	err = json.Unmarshal(body, &jsn)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsn.(map[string]interface{})[`result`].(map[string]interface{})[`body`].(map[string]interface{})[`value`])
}
