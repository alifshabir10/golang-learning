package main

import "fmt"

func operator(){

	// operator aritmatika
	var value = (((2 + 6) % 3)* 4-2 )/3

	fmt.Print(value)

	// operator perbandingan
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n",value, isEqual)

	// operator logika
	var left = false
	var right = true

	var leftAndRight  = left && right
	fmt.Printf("left && right \t(%t) \n", leftAndRight)
	
	var leftOrRight  = left || right
	fmt.Printf("left && right \t(%t) \n", leftOrRight)

	var leftReverse  = !left
	fmt.Printf("left && right \t(%t) \n", leftReverse)

}