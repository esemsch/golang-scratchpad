package main

import "fmt"

// dispells my confusion about interface{}.  interface{} is actually a supertype of both interface{} and *interface{}
func main() {
	x := 1
	y := new(interface{})

	f(x)
	f(y)

	//g(x) - this errors because int is not a subtype of *interface{}
	g(y)
}

func f(i interface{}) {
	fmt.Println(i)
}

func g(i *interface{}) {
	fmt.Println(i)
}
