package main

import "fmt"

func main() {
	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "and b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "and c is", c)
	b, c = 80, 90 // assign new values to the already declared variables b and c
	fmt.Println("New b is", b, "and c is", c)
}
