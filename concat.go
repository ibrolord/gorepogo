package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str1 := "hello"
	str2 := "world"
	str3 := str1 + str2
	fmt.Println(str3)

	str4 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str4)

	fmt.Println("--- demo String To Numeric ---")
	str_val1 := "5"
	str_val2 := "2.8769"

	var err error
	var int_val1 int64
	int_val1, err = strconv.ParseInt(str_val1, 10, 32) //base10, 64 size
	if err == nil {
		fmt.Println(int_val1)
	} else {
		fmt.Println(err)
	}

	var float_val2 float64
	float_val2, err = strconv.ParseFloat(str_val2, 64) //64 size
	if err == nil {
		fmt.Println(float_val2)
	} else {
		fmt.Println(err)
	}

	fmt.Println("--- Demo Numeric to String ---")
	num1 := 8
	num2 := 5.872

	var str_num1 string
	str_num1 = fmt.Sprintf("%d", num1)
	fmt.Println(str_num1)

	var str_num2 string
	str_num2 = fmt.Sprintf("%f", num2)
	fmt.Println(str_num2)

	fmt.Println("--- Demo String Parser ---")
	data := "Berlin;Amsterdam;London;Tokyo"
	fmt.Println(data)
	cities := strings.Split(data, ";")
	for _, city := range cities {
		fmt.Println(city)
	}

	fmt.Println("--- Demo String Length ---")
	str_data := "Welcome to go"
	len_data := len(str_data)
	fmt.Printf("len=%d \n", len_data)

	fmt.Println("--- Demo Copy Data ---")
	sample := "Hello World, Go!"
	fmt.Println(sample)
	fmt.Println(sample[0:len(sample)])               //copy all
	fmt.Println(sample[:5])                          //copy 5 characters
	fmt.Println(sample[3:8])                         //copy characters from index 3 until 8
	fmt.Println(sample[len(sample)-5 : len(sample)]) //5 copy characters

	fmt.Println("--- Demo Upper/Lower characters ---")
	message := "Hello World, GO!"
	fmt.Println(message)
	fmt.Println(strings.ToUpper(message)) //Upper Chars
	fmt.Println(strings.ToLower(message)) //Upper chars
}
