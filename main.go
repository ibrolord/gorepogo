package main

import "fmt"

func main() {
	//declare variables
	var str string
	var n, m int
	var mn float32

	//assign values
	str = "Hello World"
	n = 10
	m = 2.45

	fmt.Println("Value of str= ", str)
	fmt.Println("value of n= ", n)
	fmt.Println("value of m= ", m)
	fmt.Println("value of mn= ", mn)

	//declare and assign values to variables
	var city string = "London"
	var x int = 100

	fmt.Println("Value of city= ", city)
	fmt.Println("Value of x= ", x)

	//declare variable with defining its type
	country := "DE"
	val := 15
	fmt.Println("Value of country= ", country)
	fmt.Println("Value of val= ", val)

	//declare multiple variables
	var (
		name  string
		email string
		age   int
	)
	name = "John"
	email = "john@email.com"
	age = 27

	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(age)
}
