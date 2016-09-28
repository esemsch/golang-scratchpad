package main

import (
	"fmt"
)

func main() {
	vals := []string{"a", "b", "c", "d", "e"}
	other_vals := []string{"x", "y", "z"}

	fmt.Println(vals)
	fmt.Println(other_vals)

	swap(&vals, &other_vals)

	fmt.Println(vals)
	fmt.Println(other_vals)

	var x *[]string = &vals
	var y *[]string = &other_vals

	swap(x, y)
	fmt.Println(*x)
	fmt.Println(*y)

	swap(x, y)
	fmt.Println(*x)
	fmt.Println(*y)

}

func swap(a *[]string, b *[]string) {
	x := *a
	println("x =", x)
	println("*a =", *a)
	*a = *b
	println("*a =", *a)
	*b = x
	println("*b =", *b)
}
