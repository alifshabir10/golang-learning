package main

import "fmt"

func main(){

	// tipe data non-Desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -123242322422

	fmt.Println("bilangan positif: %d\n", positiveNumber)
	fmt.Println("bilangan positif: %d\n", negativeNumber)

	//  tipe data Desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	// tipe data Boolean
	var exist bool = false
	fmt.Printf("exist? %t \n", exist)

	// tipe data String
	var message string = "hallo"
	
	fmt.Printf("message: %s \n", message)
	
	var message2 string = `Nama saya "JOHN WICK". 
	salam kenal.
	mari kita belajar "Golang"
	`
	
	fmt.Printf("message: %s \n", message2)
}