package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "You have to provide the name of the file")
		os.Exit(1)
	}

	fileName := args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't open file %s", fileName)
		os.Exit(1)
	}

	wholeFile := []byte{}
	buffer := make([]byte, 4096)

	for {
		n, err := file.Read(buffer)
		if n == 0 && err == io.EOF {
			err := file.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Couldn't close file %s", fileName)
				os.Exit(1)
			}
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from file %s. Error: ", fileName, err)
			os.Exit(1)
		} else {
			wholeFile = append(wholeFile, buffer[0:n]...)
		}

	}

	fmt.Println(string(wholeFile))

}
