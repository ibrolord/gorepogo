package main

import "fmt"

func main() {
	//a = "This is a random Var" string
	//b = "This is another Var Sting" string
	var a string = "This is a random Var"
	var b string = "This is another Var Sting"
	var c string
	c := a + b
	fmt.Printf("%a -- %a -- %a", a, b, c)
	fmt.Println("\nJust talking smak")
}
