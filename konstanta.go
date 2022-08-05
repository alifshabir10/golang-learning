package main

import "fmt"

func main(){
	// contoh pertama
	const firstName string = "john"
	fmt.Print("hallo", firstName, "!\n")

	// contoh ke dua

	const lastName = "wick"
	fmt.Print("nice to meet you", lastName, "!\n")

	// fungsi fmt.println()
	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john ", " ", "wick\n")
}