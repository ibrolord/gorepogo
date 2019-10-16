package main

import "fmt"

func main() {

	fmt.Print("Enter 0 - 3: ")
	var selected int
	fmt.Scanf("%d",&selected)
	fmt.Printf("You entered %d ",selected) 
	//selected := 3
	fmt.Print("\n")

	switch selected {
		case 0: 
			fmt.Println("Selected = 0")
		case 1:
			fmt.Println("Selected = 1")
		case 2:
			fmt.Println("Selected = 2")
		case 3: 
			fmt.Println("Seleccted = 3")
		default:
			fmt.Println("Other...")
		}
	}
		
