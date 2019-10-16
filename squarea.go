package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Square Area Calculation")
	fmt.Print("Enter a length: ")
	var length float64
	fmt.Scanf("%f", &length)

	area := math.Pi * math.Pow(length,4)
	fmt.Printf("Square area with length %.2f = %.2f \n",length, area)
}
