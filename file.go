package main

import (
	"fmt"
	"io/ioutil"
)

func writeFile(message string) {
	bytes := []byte(message)
	ioutil.WriteFile("/home/ibro/go/testgo.txt", bytes, 0644)
	fmt.Println("created a file")
}

func readFile() {
	data, _ := ioutil.ReadFile("/home/ibro/go/testgo.txt")
	fmt.Println("file content:")
	fmt.Println(string(data))
}

func main() {

	//write data into a file
	fmt.Println("Writing data into a file")
	writeFile("Welcome to golang bitties")
	readFile()

	//read data from a file
	fmt.Println("reading data from a file")
	readFile()
}
