package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//define slice
	fmt.Println("defince slices")
	var numbers[] int
	numbers = make([]int,5)
	matrix := make([][]int,3*3)

	//insert data
	fmt.Println(">>>>>>insert slice data")
	for i:=0;i<5;i++ {
		numbers[i] = rand.Intn(100) //random number
