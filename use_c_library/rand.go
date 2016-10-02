// Copied from: https://blog.golang.org/c-go-cgo
// Requires gcc

package main

//#include<stdlib.h>
////It is absolutely necessary that the line reads only import "C".  Importing multiple packages doesn't work!!
import "C"
import "fmt"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	Seed(1)
	fmt.Println(Random())
}
