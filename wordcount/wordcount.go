package main

import (
	"fmt"
	//"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "You have to provide the name of the file")
		os.Exit(1)
	}

	fileName := args[1]

	file,err := os.Open(fileName)
	if err!=nil {
		fmt.Fprintf(os.Stderr, "Couldn't open file %s", fileName)
		os.Exit(1)
	}

	wholeFile := []byte{}
	buffer := make(byte[],4096)

	for {

	}

}
