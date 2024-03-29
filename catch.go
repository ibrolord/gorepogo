package main

import (
	"fmt"
	"github.com/manucorporat/try"
)

func main() {
	try.This(func() {
		a := 10
		b := 0
		c := 0

		c = a / b
		fmt.Printf("result = %.2f \n", c)

	}).Finally(func() {
		fmt.Println("Done")

	}).Catch(func(_ try.E) {
		fmt.Println("exception caught") //print
		try.Throw()
	})
}
